package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/labstack/echo/v4"
)

func (a *httpApp) DeleteStory(c echo.Context) error {
	sId := c.Param("sId")

	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = a.sService.Delete(c.Request().Context(), me.Sub, sId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, comtypes.MessageRes{Message: "Party removed"})
}
