package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *httpApp) GetParty(c echo.Context) error {
	cId := c.Param("cId")
	pId := c.Param("pId")

	p, err := a.partyService.Get(c.Request().Context(), cId, pId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}
