package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/labstack/echo/v4"
)

// @Summary Delete a Story
// @Description Deletes a Story from our db
// @Tags CRUD
// @Accept json
// @Produce json
// @Param sId path string true "Story Id"
// @Success 200 {object} comtypes.MessageRes
// @Failure 400 {object} echo.HTTPError
// @Router /{sId} [delete]
func (a *httpApp) DeleteStory(c echo.Context) error {
	sId := c.Param("sId")

	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = a.sService.Delete(c.Request().Context(), me.Sub, sId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, comtypes.MessageRes{Message: "Party removed"})
}
