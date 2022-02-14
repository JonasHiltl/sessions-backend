package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/jonashiltl/sessions-backend/services/party/internal/dto"
	"github.com/labstack/echo/v4"
)

type CreatePartyBody struct {
	Title    string  `json:"title"     validate:"required"`
	Lat      float64 `json:"lat"       validate:"required,latitude"`
	Long     float64 `json:"long"      validate:"required,longitude"`
	IsPublic bool    `json:"isPublic"`
}

// @Summary Create a party
// @Description Create a new Party in DB
// @Tags CRUD
// @Accept json
// @Produce json
// @Param Body body CreatePartyBody true "The body to create a party"
// @Success 201 {object} datastruct.PublicParty
// @Failure 400 {object} echo.HTTPError
// @Router / [post]
func (a *httpApp) CreateParty(c echo.Context) error {
	var reqBody CreatePartyBody

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
		Title:     reqBody.Title,
		CreatorId: me.Sub,
		Lat:       reqBody.Lat,
		Long:      reqBody.Long,
		IsPublic:  reqBody.IsPublic,
	}

	p, err := a.partyService.Create(c.Request().Context(), d)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, p.ToPublicParty())
}
