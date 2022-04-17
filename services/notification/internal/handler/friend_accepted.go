package handler

import (
	"log"

	"github.com/jonashiltl/sessions-backend/packages/events"
)

func (s *server) FriendAccepted(fa *events.FriendAccepted) {
	log.Printf("%v accepted request from %v", fa.FriendId, fa.UserId)
}
