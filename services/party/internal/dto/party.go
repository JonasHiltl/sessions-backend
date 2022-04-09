package dto

import (
	"time"
)

type Party struct {
	UserId        string
	Id            string
	Title         string
	IsPublic      bool
	Lat           float32
	Long          float32
	StreetAddress string
	PostalCode    string
	State         string
	Country       string
	StartDate     time.Time
	Ttl           time.Time
}
