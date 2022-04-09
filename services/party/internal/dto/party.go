package dto

import (
	"time"
)

type Party struct {
	UserId    string
	Id        string
	IsPublic  bool
	Lat       float64
	Long      float64
	Title     string
	StartDate time.Time
	Ttl       time.Time
}
