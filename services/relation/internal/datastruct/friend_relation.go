package datastruct

import (
	"time"

	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
)

type FriendRelation struct {
	UserId     string    `json:"user_id"    db:"user_id"    validate:"required"`
	FriendId   string    `json:"friend_id"  db:"friend_id"  validate:"required"`
	Accepted   bool      `json:"accepted"   db:"accepted"`
	CreatedAt  time.Time `json:"created_at" db:"created_at" validate:"required"`
	AcceptedAt time.Time `json:"accepted_at" db:"accepted_at" validate:"required"`
}

func (fr FriendRelation) ToGRPCProfile() *rg.FriendRelation {
	return &rg.FriendRelation{
		UserId:     fr.UserId,
		FriendId:   fr.FriendId,
		Accepted:   fr.Accepted,
		CreatedAt:  fr.CreatedAt.UTC().Format(time.RFC3339),
		AcceptedAt: fr.AcceptedAt.UTC().Format(time.RFC3339),
	}
}
