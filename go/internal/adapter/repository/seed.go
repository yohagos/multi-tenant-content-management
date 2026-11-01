package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yohagos/multi-content-management/internal/core/domain"
	"golang.org/x/crypto/bcrypt"
)

func SeedAdminUser(db *sqlx.DB, email, password string) error {
	ctx := context.Background()

	var exists bool
	query := `
		SELECT EXISTS(
			SELECT 1 
			FROM users 
			WHERE email = $1 AND role = 'admin' AND deleted_at IS NULL
		)
	`
	if err := db.GetContext(ctx, &exists, query, email); err != nil {
		return err
	}

	if exists {
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	user := domain.User{
		Email: email,
		PasswordHash: string(hashedPassword),
		FirstName: "Admin",
		LastName: "User",
		Role: "admin",
		Active: true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	insertQuery := `
		INSERT INTO users (email, password_hash, first_name, last_name, role, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err = db.ExecContext(ctx, insertQuery,
		user.Email, user.PasswordHash, user.FirstName, user.LastName,
		user.Role, user.Active, user.CreatedAt, user.UpdatedAt,
	)

	return err
}