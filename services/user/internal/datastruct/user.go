package datastruct

import (
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id            primitive.ObjectID `json:"id"                      bson:"_id"                     validate:"required"`
	Provider      string             `json:"provider"                bson:"provider,omitempty"      validate:"required_without=PasswordHash"`
	Email         string             `json:"email,omitempty"         bson:"email,omitempty"         validate:"email,required_with=PasswordHash"`
	Username      string             `json:"username"                bson:"username,omitempty"      validate:"required"`
	Firstname     string             `json:"firstname"               bson:"firstname,omitempty"     validate:"required"`
	Lastname      string             `json:"lastname,omitempty"      bson:"lastname,omitempty"`
	Avatar        string             `json:"avatar,omitempty"        bson:"avatar,omitempty"`
	FriendCount   uint64             `json:"friend_count"            bson:"friend_count"`
	EmailVerified bool               `json:"email_verified"          bson:"email_verified"`
	EmailCode     string             `json:"email_code"              bson:"email_code,omitempty"`
	PasswordHash  string             `json:"password_hash,omitempty" bson:"password_hash,omitempty" validate:"required_without=Provider"`
	Role          string             `json:"role,omitempty"          bson:"role,omitempty"          validate:"required"`
}

func (p User) ToGRPCUser() *ug.User {
	return &ug.User{
		Id:          p.Id.Hex(),
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
		Id:          p.Id.Hex(),
		Username:    p.Username,
		Firstname:   p.Firstname,
		Lastname:    p.Lastname,
		Avatar:      p.Avatar,
		FriendCount: p.FriendCount,
	}
}
