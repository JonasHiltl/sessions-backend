package handler

import (
	"log"

	"github.com/jonashiltl/sessions-backend/packages/events"
)

func (*server) PartyCreated(p *events.PartyCreated) {
	log.Printf("%v", p.Party)
}
