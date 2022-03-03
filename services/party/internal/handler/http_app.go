package handler

import (
	"github.com/jonashiltl/sessions-backend/services/party/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
)

type HttpApp interface {
	CreateParty(c echo.Context) error
	DeleteParty(c echo.Context) error
	UpdateParty(c echo.Context) error
	GetParty(c echo.Context) error
	GetByUser(c echo.Context) error
	GeoSearch(c echo.Context) error
}

type httpApp struct {
	partyService service.PartyService
	nc           *nats.EncodedConn
}

func NewHttpApp(partyService service.PartyService, nc *nats.EncodedConn) HttpApp {
	return &httpApp{
		partyService: partyService,
		nc:           nc,
	}
}
