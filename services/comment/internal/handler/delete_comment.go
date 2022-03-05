package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/labstack/echo/v4"
)

// @Summary Delete a Comment
// @Description Delete a Comment by its id with the party id
// @Tags CRUD
// @Accept json
// @Produce json
// @Param cId path string true "Comment Id"
// @Param cId path string true "Party Id"
// @Success 200 {object} comtypes.MessageRes
// @Failure 400 {object} echo.HTTPError
// @Router /{pId}/{cId} [delete]
func (a *httpApp) DeleteComment(c echo.Context) error {
	cId := c.Param("cId")
	pId := c.Param("pId")

	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = a.cs.Delete(c.Request().Context(), me.Sub, pId, cId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, comtypes.MessageRes{Message: "Comment removed"})
}
