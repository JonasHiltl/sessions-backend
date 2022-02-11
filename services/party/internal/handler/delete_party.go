package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/labstack/echo/v4"
)

func (a *httpApp) DeleteParty(c echo.Context) error {
	cId := c.Param("cId")
	pId := c.Param("pId")

	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if cId != me.Sub {
		return echo.NewHTTPError(http.StatusBadRequest, "You can only delete your own parties")
	}

	err = a.partyService.Delete(c.Request().Context(), me.Sub, pId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, comtypes.MessageRes{Message: "Party removed"})
}
