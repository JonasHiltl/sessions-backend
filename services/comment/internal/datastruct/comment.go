package datastruct

import "time"

type Comment struct {
	Id         string    `json:"id"        db:"comment_id" validate:"required"`
	PId        string    `json:"party_id"  db:"party_id"   validate:"required"`
	AId        string    `json:"author_id" db:"author_id"  validate:"required"`
	Body       string    `json:"body"      db:"body"       validate:"required"`
	Created_at time.Time `json:"createdAt" db:"created_at" validate:"required"`
}
