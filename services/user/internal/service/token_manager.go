package service

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/jonashiltl/sessions-backend/services/user/ent"
)

type TokenManager interface {
	NewJWT(u ent.User) (string, error)
}

type tokenManager struct {
	secret string
}

func NewTokenManager() TokenManager {
	secret := os.Getenv("TOKEN_SECRET")
	fmt.Println(secret)
	return &tokenManager{secret: secret}
}

func (t *tokenManager) NewJWT(u ent.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":  u.ID.String(),
		"iss":  "sessions.com",
		"role": u.Role.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(t.secret))
}
