package datastruct

import (
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

type Profile struct {
	Id          string `json:"id"                 db:"id"        validate:"required"`
	Email       string `json:"email,omitempty"    db:"email"     validate:"email,required_with=PasswordHash"`
	Username    string `json:"username"           db:"username"  validate:"required"`
	Firstname   string `json:"firstname"          db:"firstname" validate:"required"`
	Lastname    string `json:"lastname,omitempty" db:"lastname"`
	Avatar      string `json:"avatar,omitempty"   db:"avatar"`
	FriendCount uint32 `json:"friend_count"       db:"friend_count"`
}

func (p Profile) ToGRPCProfile() *ug.Profile {
	return &ug.Profile{
		Id:          p.Id,
		Username:    p.Username,
		Firstname:   p.Firstname,
		Lastname:    p.Lastname,
		Avatar:      p.Avatar,
		FriendCount: p.FriendCount,
	}
}
