package handler

import (
	"log"

	"github.com/jonashiltl/sessions-backend/packages/events"
)

func (*server) FriendRequested(fr *events.FriendRequested) {
	log.Printf("%v got a friend request from %v", fr.FriendId, fr.UserId)
}
