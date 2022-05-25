package datastruct

import (
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

type User struct {
	Id            string `json:"id"                      db:"id"             validate:"required"`
	Username      string `json:"username"                db:"username"       validate:"required"`
	Email         string `json:"email,omitempty"         db:"email"          validate:"email,required_with=PasswordHash"`
	Firstname     string `json:"firstname"               db:"firstname"      validate:"required"`
	Lastname      string `json:"lastname,omitempty"      db:"lastname"`
	Avatar        string `json:"avatar,omitempty"        db:"avatar"`
	FriendCount   uint32 `json:"friend_count"            db:"friend_count"`
	Provider      string `json:"provider"                db:"provider"       validate:"required_without=PasswordHash"`
	EmailVerified bool   `json:"email_verified"          db:"email_verified"`
	EmailCode     string `json:"email_code"              db:"email_code"`
	PasswordHash  string `json:"password_hash,omitempty" db:"password_hash"  validate:"required_without=Provider"`
	Role          string `json:"role,omitempty"          db:"role"           validate:"required"`
}

func (p User) ToGRPCUser() *ug.User {
	return &ug.User{
		Id:          p.Id,
		Email:       p.Email,
		Username:    p.Username,
		Firstname:   p.Firstname,
		Lastname:    p.Lastname,
		Avatar:      p.Avatar,
		FriendCount: p.FriendCount,
	}
}

func (p User) ToGRPCProfile() *ug.Profile {
	return &ug.Profile{
		Id:          p.Id,
		Username:    p.Username,
		Firstname:   p.Firstname,
		Lastname:    p.Lastname,
		Avatar:      p.Avatar,
		FriendCount: p.FriendCount,
	}
}
