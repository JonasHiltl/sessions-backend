package handler

import (
	"github.com/jonashiltl/sessions-backend/packages/events"
	"github.com/jonashiltl/sessions-backend/services/notification/internal/config"
	mail "github.com/xhit/go-simple-mail/v2"
)

type server struct {
	mail   *mail.SMTPClient
	config config.Config
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
