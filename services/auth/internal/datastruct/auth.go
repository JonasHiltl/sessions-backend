package datastruct

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Provider is required if password does not exist
// Email is required if password exists
// Password is required if Provider does not exists
type AuthUser struct {
	Id            primitive.ObjectID `json:"id"                      bson:"_id"`
	Provider      string             `json:"provider"                bson:"provider,omitempty"      validate:"required_without=PasswordHash"`
	Email         string             `json:"email,omitempty"         bson:"email,omitempty"         validate:"email,required_with=PasswordHash"`
	EmailVerified bool               `json:"email_verified"          bson:"email_verified"`
	EmailCode     string             `json:"email_code"              bson:"email_code,omitempty"`
	PasswordHash  string             `json:"password_hash,omitempty" bson:"password_hash,omitempty" validate:"required_without=Provider"`
	Role          string             `json:"role,omitempty"          bson:"role,omitempty"          validate:"required"`
}
