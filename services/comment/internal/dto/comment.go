package dto

type Comment struct {
	PartyId  string `json:"party_id"  db:"party_id"   validate:"required"`
	AuthorId string `json:"author_id" db:"author_id"  validate:"required"`
	Body     string `json:"body"      db:"body"       validate:"required"`
}
