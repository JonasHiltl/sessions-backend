package handlers

import (
	"github.com/jonashiltl/sessions-backend/services/party/internal/service"
	"github.com/labstack/echo/v4"
)

type HttpApp interface {
	CreateParty(c echo.Context) error
	DeleteParty(c echo.Context) error
	UpdateParty(c echo.Context) error
	GetParty(c echo.Context) error
}

type httpApp struct {
	partyService service.PartyService
}

func NewHttpApp(partyService service.PartyService) HttpApp {
	return &httpApp{
		partyService: partyService,
	}
}
