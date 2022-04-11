package dto

import (
	"github.com/jonashiltl/sessions-backend/packages/types"
)

type AuthUser struct {
	Id            string
	Provider      types.Provider
	Email         string
	EmailVerified bool
	PasswordHash  string
	Role          types.Role
}
