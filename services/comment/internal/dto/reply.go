package dto

import "time"

type Reply struct {
	CommentId  string    `json:"comment_id" db:"comment_id" validate:"required"`
	AuthorId   string    `json:"author_id"  db:"author_id"  validate:"required"`
	Body       string    `json:"body"       db:"body"       validate:"required"`
	Created_at time.Time `json:"createdAt"  db:"created_at" validate:"required"`
}
