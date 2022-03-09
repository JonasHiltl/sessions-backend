package handler

import (
	gonats "github.com/nats-io/nats.go"
)

type server struct {
}

type Server interface {
	EmailVerification(m *gonats.Msg)
	PartyCreated(m *gonats.Msg)
	Commented(m *gonats.Msg)
	Replied(m *gonats.Msg)
	FriendRequest(m *gonats.Msg)
	FriendAccepted(m *gonats.Msg)
}

func NewServer() Server {
	return &server{}
}
