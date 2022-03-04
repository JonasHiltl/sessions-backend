package dto

type Comment struct {
	PId  string `json:"party_id"  db:"party_id"   validate:"required"`
	AId  string `json:"author_id" db:"author_id"  validate:"required"`
	Body string `json:"body"      db:"body"       validate:"required"`
}
