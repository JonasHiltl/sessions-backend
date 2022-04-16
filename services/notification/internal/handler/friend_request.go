package handler

import (
	"log"

	"github.com/nats-io/nats.go"
)

func (*server) FriendRequest(m *nats.Msg) {
	log.Printf("%v", m.Data)
}
