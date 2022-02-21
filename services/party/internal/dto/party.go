package dto

import (
	"time"
)

type Party struct {
	UId      string    `json:"userId"`
	Id       string    `json:"id"`
	IsPublic bool      `json:"isPublic"`
	Lat      float64   `json:"lat"`
	Long     float64   `json:"long"`
	Title    string    `json:"title"`
	Ttl      time.Time `json:"ttl"`
}
