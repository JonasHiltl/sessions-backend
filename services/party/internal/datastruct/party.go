package datastruct

import (
	"github.com/jonashiltl/sessions-backend/services/party/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PartyCollectionName = "parties"

type Party struct {
	ID        primitive.ObjectID  `json:"id"         bson:"_id"`
	Title     string              `json:"title"      bson:"title" validate:"required"`
	CreatorId string              `json:"creatorId"  bson:"creatorId"`
	Stories   []string            `json:"stories"    bson:"title,omitempty"`
	Location  utils.GeoJson       `json:"location"   bson:"location,omitempty" validate:"required"`
	IsGlobal  bool                `json:"isGlobal"   bson:"isGlobal,omitempty" validate:"required"`
	CreatedAt primitive.Timestamp `json:"createdAt"  bson:"createdAt,omitempty"`
}

type RequestPary struct {
	CreatorId string  `json:"creatorId" validate:"required"`
	Title     string  `json:"title"     validate:"required"`
	Lat       float64 `json:"lat"       validate:"required, latitude"`
	Long      float64 `json:"long"      validate:"required, longitude"`
	IsGlobal  bool    `json:"isGlobal"      validate:"required, boolean"`
}
