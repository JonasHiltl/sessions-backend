package datastruct

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Provider is required if password does not exist
// Email is required if password does exist
// Password is required if password does exist
// Username is required if Provider does not exist
type AuthUser struct {
	Id        primitive.ObjectID `json:"id"                 bson:"_id"`
	Provider  Provider           `json:"provider"           bson:"provider,omitempty"  validate:"required_without=Password"`
	Email     string             `json:"email,omitempty"    bson:"email,omitempty"     validate:"email,required_with=Password"`
	EmailVerified bool `json:"email_verified" bson:"email_verified"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"  validate:"min=8,required_with=Email"`
	Role      string             `json:"role,omitempty"     bson:"role,omitempty" validate:"required"`
}
