package ent

import (
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

func (u *User) ToPublicUser() *ug.PublicUser {
	return &ug.PublicUser{
		Id:          u.ID,
		Username:    u.Username,
		Firstname:   u.FirstName,
		Lastname:    u.LastName,
		Avatar:      u.Avatar,
		Role:        string(u.Role),
		FriendCount: int32(u.FriendCount),
	}
}
