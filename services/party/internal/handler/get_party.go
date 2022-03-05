package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Get a party
// @Description Get a Party by it's id
// @Tags CRUD
// @Accept json
// @Produce json
// @Param pId path string true "Party Id"
// @Success 200 {object} datastruct.PublicParty
// @Failure 400 {object} echo.HTTPError
// @Router /{pId} [get]
func (a *httpApp) GetParty(c echo.Context) error {
	pId := c.Param("pId")

	p, err := a.partyService.Get(c.Request().Context(), pId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, p.ToPublicParty())
}
