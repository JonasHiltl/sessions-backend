package handlers

import (
	"net/http"
	"time"

	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/party/internal/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		Location:  utils.NewPoint(reqBody.Location.Long, reqBody.Location.Lat),
		IsGlobal:  reqBody.IsGlobal,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	p, err := a.partyService.Create(c.Request().Context(), p)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, p)
}
