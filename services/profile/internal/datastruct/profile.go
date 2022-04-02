package datastruct

import (
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Email is required if password does exist
// Username is required if Provider does not exist
type Profile struct {
	Id        primitive.ObjectID `json:"id"                 bson:"_id"                 validate:"required"`
	Username  string             `json:"username"           bson:"username,omitempty"  validate:"required"`
	Firstname string             `json:"firstname"          bson:"firstname,omitempty" validate:"required"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Avatar    string             `json:"avatar,omitempty"   bson:"avatar,omitempty"`
}

func (p Profile) ToPublicProfile() *pg.PublicProfile {
	return &pg.PublicProfile{
		Id:        p.Id.Hex(),
		Username:  p.Username,
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Avatar:    p.Avatar,
	}
}
