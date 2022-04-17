package handler

import (
	"log"

	"github.com/jonashiltl/sessions-backend/packages/events"
)

func (s *server) Registered(m *events.Registered) {
	log.Printf("%v", m)
}
