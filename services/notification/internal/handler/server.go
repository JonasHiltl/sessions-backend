package handler

import (
	"github.com/jonashiltl/sessions-backend/packages/events"
	mail "github.com/xhit/go-simple-mail/v2"
)

type server struct {
	mail *mail.SMTPClient
}

type Server interface {
	Registered(m *events.Registered)
	PartyCreated(p *events.PartyCreated)
	FriendRequested(m *events.FriendRequested)
	FriendAccepted(m *events.FriendAccepted)
}

func NewServer(mail *mail.SMTPClient) Server {
	return &server{mail: mail}
}
