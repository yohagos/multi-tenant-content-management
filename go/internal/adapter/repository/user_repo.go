package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
)

type userRepository struct {
	db     *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) port.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	query := `
		INSERT INTO users (email, password_hash, first_name, last_name, role, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.ExecContext(
		ctx, query, user.Email, user.PasswordHash, user.FirstName,
		user.LastName, user.Role, user.Active, user.CreatedAt, user.UpdatedAt,
	)

	return err
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	query := `
		SELECT is, tenant_id, email, password_hash, first_name, last_name, role,
		active, last_login_at, created_at, updated_at, deleted_at 
		FROM users
		WHERE id = $1 AND deleted_at IS NULL
	`

	var user domain.User
	err := r.db.GetContext(ctx, &user, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT is, tenant_id, email, password_hash, first_name, last_name, role,
		active, last_login_at, created_at, updated_at, deleted_at 
		FROM users
		WHERE email = $1 AND deleted_at IS NULL
	`

	var user domain.User
	err := r.db.GetContext(ctx, &user, query, email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) List(ctx context.Context, tenantID string, limit, offset int) ([]domain.User, int, error) {
	countQuery := `SELECT COUNT(*) FROM users WHERE tenant_id = $1 AND deleted_at IS NULL`

	var total int
	if err := r.db.GetContext(ctx, &total, countQuery, tenantID); err != nil {
		return nil, 0, err
	}

	if limit == 0 {
		limit = 20
	}

	query := `
		SELECT is, tenant_id, email, password_hash, first_name, last_name, role,
		active, last_login_at, created_at, updated_at, deleted_at
		FROM users
		WHERE tenant_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	var users []domain.User
	err := r.db.SelectContext(ctx, &users, query, tenantID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	user.UpdatedAt = time.Now()

	query := `
		UPDATE users
		SET email = $1, first_name = $2, last_name = $3, role = $4, active = $5, updated_at = $6
		WHERE id = $7 AND deleted_at IS NULL
	`

	_, err := r.db.ExecContext(ctx, query,
		user.Email, user.FirstName, user.LastName, user.Role,
		user.Active, user.UpdatedAt, user.ID,
	)

	return err
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE users SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *userRepository) UpdateLastLogin(ctx context.Context, id string) error {
	now := time.Now()
	query := `UPDATE users SET last_login_at = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, now, id)
	return err
}
