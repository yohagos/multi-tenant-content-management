package port

import (
	"context"

	"github.com/yohagos/multi-content-management/internal/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	List(ctx context.Context, tenantID string, limit, offset int) ([]domain.User, int, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
	UpdateLastLogin(ctx context.Context, id string) error
	UpdateUserMetrics(ctx context.Context) error
}

type TokenRepository interface {
	Create(ctx context.Context, token *domain.Token) error
	GetByToken(ctx context.Context, token string) (*domain.Token, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (*domain.Token, error)
	Revoke(ctx context.Context, token string) error
	RevokeAllUserTokens(ctx context.Context, userID string) error
	DeleteExpired(ctx context.Context) error
}
