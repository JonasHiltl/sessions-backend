package datastruct

import (
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Provider is required if password does not exist
// Email is required if password does exist
// Password is required if password does exist
// Username is required if Provider does not exist
type User struct {
	Id        primitive.ObjectID `json:"id"                 bson:"_id"`
	Provider  Provider           `json:"provider"           bson:"provider,omitempty"  validate:"required_without=Password"`
	Username  string             `json:"username"           bson:"username,omitempty"  validate:"required_without=Provider"`
	Email     string             `json:"email,omitempty"    bson:"email,omitempty"     validate:"email,required_with=Password"`
	Firstname string             `json:"firstname"          bson:"firstname,omitempty" validate:"required"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"  validate:"min=8,required_with=Email"`
	Avatar    string             `json:"avatar,omitempty"   bson:"avatar,omitempty"`
	Role      string             `json:"role,omitempty"     bson:"role,omitempty" validate:"required"`
}

func (u User) ToPublicUser() *ug.PublicUser {
	return &ug.PublicUser{
		Id:        u.Id.Hex(),
		Username:  u.Username,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Avatar:    u.Avatar,
		Role:      u.Role,
	}
}
