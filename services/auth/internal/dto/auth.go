package dto

import (
	"github.com/jonashiltl/sessions-backend/packages/comtypes"
)

type AuthUser struct {
	Id            string
	Provider      comtypes.Provider
	Email         string
	EmailVerified bool
	PasswordHash  string
	Role          comtypes.Role
}
