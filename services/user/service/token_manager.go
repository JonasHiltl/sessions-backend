package service

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jonashiltl/sessions-backend/services/user/datastruct"
)

type TokenManager interface {
	NewJWT(u datastruct.User) (string, error)
}

type tokenManager struct {
	secret string
}

func NewTokenManager(secret string) TokenManager {
	return tokenManager{secret: secret}
}

func (t tokenManager) NewJWT(u datastruct.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":           u.Id,
		"iss":           "sessions.com",
		"emailVerified": u.EmailVerified,
		"role":          u.Role,
		"iat":           time.Now().Unix(),
	}

	if u.Provider != "" {
		claims["provider"] = u.Provider
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(t.secret))
}
