package ent

import "github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"

func (u *User) ToPublicProfile() datastruct.PublicUser {
	return datastruct.PublicUser{
		ID:          u.ID,
		Username:    u.Username,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Picture:     u.Picture,
		Role:        string(u.Role),
		FriendCount: u.FriendCount,
	}
}
