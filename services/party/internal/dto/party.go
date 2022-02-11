package dto

import (
	"time"

	"github.com/segmentio/ksuid"
)

type Party struct {
	CId      string      `json:"creatorId"`
	KSUID    ksuid.KSUID `json:"id,omitempty"`
	IsGlobal bool        `json:"isGlobal"`
	Lat      float64     `json:"lat"`
	Long     float64     `json:"long"`
	Title    string      `json:"title"`
	Ttl      time.Time   `json:"ttl"`
}
