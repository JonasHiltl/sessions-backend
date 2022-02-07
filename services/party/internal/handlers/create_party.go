package handlers

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/party/internal/utils"
	"github.com/labstack/echo/v4"
)

func (a *httpApp) CreateParty(c echo.Context) error {
	var reqBody datastruct.RequestPary
	if err := c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't find request body")
	}

	if err := c.Validate(reqBody); err != nil {
		return err
	}

	p := datastruct.Party{
		Title:     reqBody.Title,
		CreatorId: reqBody.CreatorId,
		Location:  utils.NewPoint(reqBody.Long, reqBody.Lat),
		IsGlobal:  reqBody.IsGlobal,
	}

	p, err := a.partyService.Create(c.Request().Context(), p)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, p)
}
