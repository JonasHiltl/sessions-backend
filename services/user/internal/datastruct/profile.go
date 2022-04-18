package datastruct

import (
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	Id          primitive.ObjectID `json:"id"                 bson:"_id"                 validate:"required"`
	Email       string             `json:"email,omitempty"    bson:"email,omitempty"     validate:"email,required_with=PasswordHash"`
	Username    string             `json:"username"           bson:"username,omitempty"  validate:"required"`
	Firstname   string             `json:"firstname"          bson:"firstname,omitempty" validate:"required"`
	Lastname    string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Avatar      string             `json:"avatar,omitempty"   bson:"avatar,omitempty"`
	FriendCount int64              `json:"friend_count"       bson:"friend_count"`
}

func (p Profile) ToGRPCProfile() *ug.Profile {
	return &ug.Profile{
		Id:          p.Id.Hex(),
		Username:    p.Username,
		Firstname:   p.Firstname,
		Lastname:    p.Lastname,
		Avatar:      p.Avatar,
		FriendCount: p.FriendCount,
	}
}
