package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/labstack/echo/v4"
)

func (a *httpApp) UpdateParty(c echo.Context) error {
	type body struct {
		Title    string  `json:"title"     validate:"required"`
		Lat      float64 `json:"lat"       validate:"required,latitude"`
		Long     float64 `json:"long"      validate:"required,longitude"`
		IsGlobal bool    `json:"isGlobal"`
	}
	var reqBody body
	if err := c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't find request body")
	}

	cId := c.Param("cId")
	pId := c.Param("pId")
	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if cId != me.Sub {
		return echo.NewHTTPError(http.StatusBadRequest, "You can only delete your own parties")
	}

	p, err := a.partyService.Update(c.Request().Context(), cId, pId, reqBody.Title)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}
