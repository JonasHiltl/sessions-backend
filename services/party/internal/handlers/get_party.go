package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *httpApp) GetParty(c echo.Context) error {
	pId := c.Param("id")

	p, err := a.partyService.GetById(c.Request().Context(), pId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}
