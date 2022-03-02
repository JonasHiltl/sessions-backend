package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type presignRes struct {
	url string
}

func (a *httpApp) PresignURL(c echo.Context) error {
	key := c.Param("key")

	url, err := a.us.PresignURL(c.Request().Context(), key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, presignRes{url: url})
}
