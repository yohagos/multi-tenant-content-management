package jwt

import (
	"time"

	gjwt "github.com/golang-jwt/jwt/v5"
)

type Config struct {
	Secret string
	Expiry time.Duration
}

type Claims struct {
	UserID string `json:"user_id"`
	gjwt.RegisteredClaims
}

func GenerateToken(userID, secret string, expiresAt time.Time) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: gjwt.RegisteredClaims{
			ExpiresAt: gjwt.NewNumericDate(expiresAt),
			IssuedAt:  gjwt.NewNumericDate(time.Now()),
		},
	}

	token := gjwt.NewWithClaims(gjwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenStr, secret string) (*Claims, error) {
	token, err := gjwt.ParseWithClaims(tokenStr, &Claims{}, func(token *gjwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, gjwt.ErrSignatureInvalid
}
