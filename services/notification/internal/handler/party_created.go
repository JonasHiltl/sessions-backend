package handler

import (
	"log"

	"github.com/jonashiltl/sessions-backend/packages/events"
)

func (s *server) PartyCreated(p *events.PartyCreated) {
	log.Printf("%v", p.Party)
}
