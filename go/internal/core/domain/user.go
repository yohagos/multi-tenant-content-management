package domain

import "time"

type User struct {
	ID           string     `db:"id" json:"id"`
	TenantID     string     `db:"tenant_id" json:"tenant_id"`
	Email        string     `db:"email" json:"email"`
	PasswordHash string     `db:"password_hash" json:"password_hash"`
	FirstName    string     `db:"first_name" json:"first_name"`
	LastName     string     `db:"last_name" json:"last_name"`
	Role         string     `db:"role" json:"role"`
	Active       bool       `db:"active" json:"active"`
	LastLoginAt  *time.Time `db:"last_login_at" json:"last_login_at"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at" json:"deleted_at"`
}

type Token struct {
	ID           string    `db:"id" json:"id"`
	UserID       string    `db:"user_id" json:"user_id"`
	Token        string    `db:"token" json:"token"`
	RefreshToken string    `db:"refresh_token" json:"refresh_token"`
	ExpiresAt    time.Time `db:"expires_at" json:"expires_at"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	Revoked      bool      `db:"revoked" json:"revoked"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	TenantID  string `json:"tenant_id"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    string `json:"expires_at"`
	User         User   `json:"user"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
