package datastruct

import "github.com/jonashiltl/sessions-backend/services/user/ent"

type PublicUser struct {
	ID          string `json:"id"`
	Username    string `json:"username"              validate:"required"`
	FirstName   string `json:"firstame"             validate:"required"`
	LastName    string `json:"lastname,omitempty"`
	Picture     string `json:"picture,omitempty"`
	Role        string `json:"role"`
	FriendCount int    `json:"friendCount"`
}

type RequestUser struct {
	Username  string `json:"username"              validate:"required"`
	FirstName string `json:"firstame"              validate:"required"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email"                 validate:"required,email"`
	Password  string `json:"password,omitempty"    validate:"required,gte=8"`
	Picture   string `json:"picture,omitempty"`
}

func AddCount(u *ent.User, count int) PublicUser {
	return PublicUser{
		ID:          u.ID.String(),
		Username:    u.Username,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Picture:     u.Picture,
		Role:        string(u.Role),
		FriendCount: count,
	}
}
