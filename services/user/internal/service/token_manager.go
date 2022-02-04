package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/jonashiltl/sessions-backend/services/user/ent"
	"github.com/spf13/viper"
)

type TokenManager interface {
	NewJWT(u ent.User) (string, error)
}

type tokenManager struct {
	secret string
}

func NewTokenManager() TokenManager {
	secret := viper.GetString("jwt.secret")
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
