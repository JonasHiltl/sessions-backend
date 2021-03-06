// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package repository

import (
	"database/sql"

	"github.com/twpayne/go-geom"
)

type Party struct {
	ID            string
	UserID        string
	Title         string
	IsPublic      bool
	Location      geom.Point
	StreetAddress sql.NullString
	PostalCode    sql.NullString
	State         sql.NullString
	Country       sql.NullString
	StartDate     sql.NullTime
	EndDate       sql.NullTime
}
