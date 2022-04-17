package handler

import (
	"net/smtp"

	"github.com/jonashiltl/sessions-backend/packages/events"
)

type server struct {
	smtp *smtp.Client
}

type Server interface {
	Registered(m *events.Registered)
	PartyCreated(p *events.PartyCreated)
	FriendRequested(m *events.FriendRequested)
	FriendAccepted(m *events.FriendAccepted)
}

func NewServer(smtp *smtp.Client) Server {
	return &server{smtp: smtp}
}
