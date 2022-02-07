package handlers

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/party/internal/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *httpApp) UpdateParty(c echo.Context) error {
	var reqBody datastruct.RequestPary

	pId := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(pId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	p := datastruct.Party{
		ID:        _id,
		Title:     reqBody.Title,
		CreatorId: reqBody.CreatorId,
		Location:  utils.NewPoint(reqBody.Long, reqBody.Lat),
		IsGlobal:  reqBody.IsGlobal,
	}

	p, err = a.partyService.Update(c.Request().Context(), p)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}
