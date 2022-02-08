package datastruct

import (
	"github.com/jonashiltl/sessions-backend/services/party/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PartyCollectionName = "parties"

type Party struct {
	ID        primitive.ObjectID `json:"id,omitempty"        bson:"_id,omitempty"`
	Title     string             `json:"title"               bson:"title"                validate:"required"`
	CreatorId string             `json:"creatorId"           bson:"creatorId"`
	Stories   []string           `json:",omitempty"          bson:"stories,omitempty"`
	Location  utils.GeoJson      `json:"location,omitempty"  bson:"location,omitempty"   validate:"required"`
	IsGlobal  bool               `json:"isGlobal"            bson:"isGlobal"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

type RequestPary struct {
	CreatorId string   `json:"creatorId" validate:"required,uuid"`
	Title     string   `json:"title"     validate:"required"`
	Location  Location `json:"location"  validate:"required"`
	IsGlobal  bool     `json:"isGlobal"`
}

type Location struct {
	Long float64 `json:"long"      validate:"required,longitude"`
	Lat  float64 `json:"lat"       validate:"required,latitude"`
}
