package datastruct

import (
	"github.com/google/uuid"
	"github.com/jonashiltl/sessions-backend/services/user/ent/user"
)

type JwtPayload struct {
	Iss  string    `json:"iss"`
	Sub  uuid.UUID `json:"sub"`
	Iat  int       `json:"iat"`
	Role user.Role `json:"role"`
}
