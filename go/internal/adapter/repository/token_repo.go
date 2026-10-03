package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
)

type tokenRepository struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) port.TokenRepository {
	return &tokenRepository{
		db: db,
	}
}

func (r *tokenRepository) Create(ctx context.Context, token *domain.Token) error {
	query := `
		INSERT INTO tokens (user_id, token, refresh_token, expires_at, created_at, revoked)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		token.UserID, token.Token, token.RefreshToken,
		token.ExpiresAt, token.CreatedAt, token.Revoked)

	return err
}

func (r *tokenRepository) GetByToken(ctx context.Context, token string) (*domain.Token, error) {
	query := `SELECT id, user_id, token, refresh_token, expires_at, created_at, revoked FROM tokens WHERE token = $1`

	var t domain.Token
	err := r.db.GetContext(ctx, &t, query, token)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *tokenRepository) GetByRefreshToken(ctx context.Context, refreshToken string) (*domain.Token, error) {
	query := `SELECT id, user_id, token, refresh_token, expires_at, created_at, revoked FROM tokens WHERE refresh_token = $1`

	var t domain.Token
	err := r.db.GetContext(ctx, &t, query, refreshToken)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *tokenRepository) Revoke(ctx context.Context, token string) error {
	query := `UPDATE tokens SET revoked = true WHERE token = $1`
	_, err := r.db.ExecContext(ctx, query, token)
	return err
}

func (r *tokenRepository) RevokeAllUserTokens(ctx context.Context, userID string) error {
	query := `UPDATE tokens SET revoked = true WHERE user_id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}

func (r *tokenRepository) DeleteExpired(ctx context.Context) error {
	query := `DELETE FROM tokens WHERE expires_at < NOW() OR revoked = true`
	_, err := r.db.ExecContext(ctx, query)
	return err
}
