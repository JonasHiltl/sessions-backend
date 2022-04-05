package dto

import (
	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/datastruct"
)

type AuthUser struct {
	Id            string
	Provider      datastruct.Provider
	Email         string
	EmailVerified bool
	PasswordHash  string
	Role          comtypes.Role
}
