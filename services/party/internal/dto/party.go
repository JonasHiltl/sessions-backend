package dto

import (
	"time"
)

type Party struct {
	UserId    string
	Id        string
	IsPublic  bool
	Lat       float32
	Long      float32
	Title     string
	StartDate time.Time
	Ttl       time.Time
}
