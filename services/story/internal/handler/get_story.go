package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *httpApp) GetStory(c echo.Context) error {
	sId := c.Param("sId")

	s, err := a.sService.Get(c.Request().Context(), sId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, s.ToPublicStory())
}
