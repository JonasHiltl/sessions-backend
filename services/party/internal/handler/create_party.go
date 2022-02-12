package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/jonashiltl/sessions-backend/services/party/internal/dto"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

func (a *httpApp) CreateParty(c echo.Context) error {
	type body struct {
		Title    string  `json:"title"     validate:"required"`
		Lat      float64 `json:"lat"       validate:"required,latitude"`
		Long     float64 `json:"long"      validate:"required,longitude"`
		IsPublic bool    `json:"isGlobal"`
	}
	var reqBody body

	if err := c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't find request body")
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}

	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	d := dto.Party{
		Title:    reqBody.Title,
		CId:      me.Sub,
		KSUID:    ksuid.New(),
		Lat:      reqBody.Lat,
		Long:     reqBody.Long,
		IsPublic: reqBody.IsPublic,
	}

	p, err := a.partyService.Create(c.Request().Context(), d)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, p.ToPublicParty())
}
