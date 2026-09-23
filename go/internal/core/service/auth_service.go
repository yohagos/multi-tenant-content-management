package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
	"github.com/yohagos/multi-content-management/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
)

type AuthService struct {
	userRepo  port.UserRepository
	tokenRepo port.TokenRepository
	jwtConfig *jwt.Config
}

func NewAuthService(
	userRepo port.UserRepository,
	tokenRepo port.TokenRepository,
	jwtConfig *jwt.Config,
) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		jwtConfig: jwtConfig,
	}
}

func (s *AuthService) Register(ctx context.Context, req *domain.RegisterRequest) (*domain.User, error) {
	existing, _ := s.userRepo.GetByEmail(ctx, req.Email)
	if existing != nil {
		return nil, ErrUserExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Role:         "editor",
		Active:       true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return nil, ErrInvalidCredentials
	}

	if !user.Active {
		return nil, errors.New("user account is deactivated")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	token, refreshToken, expiresAt, err := s.createTokens(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	_ = s.userRepo.UpdateLastLogin(ctx, user.ID)

	return &domain.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt.Format(time.RFC3339),
		User:         *user,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*domain.LoginResponse, error) {
	token, err := s.tokenRepo.GetByRefreshToken(ctx, refreshToken)
	if err != nil || token == nil || token.Revoked {
		return nil, ErrInvalidToken
	}

	if time.Now().After(token.ExpiresAt) {
		return nil, ErrTokenExpired
	}

	user, err := s.userRepo.GetByID(ctx, token.UserID)
	if err != nil || user == nil {
		return nil, ErrUserNotFound
	}

	_ = s.tokenRepo.Revoke(ctx, token.Token)

	newToken, newRefreshToken, expiresAt, err := s.createTokens(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		Token:        newToken,
		RefreshToken: newRefreshToken,
		ExpiresAt:    expiresAt.Format(time.RFC3339),
		User:         *user,
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) error {
	return s.tokenRepo.Revoke(ctx, token)
}

func (s *AuthService) ValidateToken(ctx context.Context, tokenStr string) (*domain.User, error) {
	claims, err := jwt.ValidateToken(tokenStr, s.jwtConfig.Secret)
	if err != nil {
		return nil, ErrInvalidToken
	}

	token, err := s.tokenRepo.GetByToken(ctx, tokenStr)
	if err != nil || token == nil || token.Revoked {
		return nil, ErrInvalidToken
	}

	if time.Now().After(token.ExpiresAt) {
		return nil, ErrTokenExpired
	}

	user, err := s.userRepo.GetByID(ctx, claims.UserID)
	if err != nil || user == nil || !user.Active {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *AuthService) createTokens(ctx context.Context, userID string) (string, string, time.Time, error) {
	expiresAt := time.Now().Add(s.jwtConfig.Expiry)

	tokenStr, err := jwt.GenerateToken(userID, s.jwtConfig.Secret, expiresAt)
	if err != nil {
		return "", "", time.Time{}, err
	}

	refreshToken := generateRandomToken(32)

	token := &domain.Token{
		UserID:       userID,
		Token:        tokenStr,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
		CreatedAt:    time.Now(),
		Revoked:      false,
	}

	if err := s.tokenRepo.Create(ctx, token); err != nil {
		return "", "", time.Time{}, err
	}

	return tokenStr, refreshToken, expiresAt, nil
}

func generateRandomToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
