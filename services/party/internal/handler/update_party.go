package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/jonashiltl/sessions-backend/services/party/internal/dto"
	"github.com/labstack/echo/v4"
)

type UpdatePartyBody struct {
	Title    string  `json:"title"     validate:"required"`
	Lat      float64 `json:"lat"       validate:"required,latitude"`
	Long     float64 `json:"long"      validate:"required,longitude"`
	IsPublic bool    `json:"isPublic"`
}

// @Summary Update a Party
// @Description Updates a party with the provided values
// @Tags CRUD
// @Accept json
// @Produce json
// @Param Body body UpdatePartyBody true "The body to create a Party"
// @Success 200 {object} datastruct.PublicParty
// @Failure 400 {object} echo.HTTPError
// @Router /{pId} [patch]
func (a *httpApp) UpdateParty(c echo.Context) error {
	var reqBody UpdatePartyBody
	if err := c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't find request body")
	}

	pId := c.Param("pId")
	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	d := dto.Party{
		CreatorId: me.Sub,
		Title:     reqBody.Title,
		Lat:       reqBody.Lat,
		Long:      reqBody.Long,
		IsPublic:  reqBody.IsPublic,
	}

	p, err := a.partyService.Update(c.Request().Context(), pId, d)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p.ToPublicParty())
}
