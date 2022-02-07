package handlers

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/labstack/echo/v4"
)

func (a *httpApp) DeleteParty(c echo.Context) error {
	pId := c.Param("id")

	err := a.partyService.Delete(c.Request().Context(), pId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, comtypes.MessageRes{Message: "Party removed"})
}
