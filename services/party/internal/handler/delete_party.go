package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/labstack/echo/v4"
)

// @Summary Delete a party
// @Description Deletes a party from our db
// @Tags CRUD
// @Accept json
// @Produce json
// @Param pId path string true "Party Id"
// @Success 200 {object} comtypes.MessageRes
// @Failure 400 {object} echo.HTTPError
// @Router /{pId} [delete]
func (a *httpApp) DeleteParty(c echo.Context) error {
	pId := c.Param("pId")

	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = a.partyService.Delete(c.Request().Context(), me.Sub, pId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, comtypes.MessageRes{Message: "Party removed"})
}
