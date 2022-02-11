package handlers

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/labstack/echo/v4"
)

// @Summary Delete a user
// @Description Deletes a user from our DB
// @Tags CRUD
// @Accept json
// @Produce json
// @Param id path string true "User Id"
// @Success 200 {object} datastruct.PublicUser
// @Failure 400 {object} echo.HTTPError
// @Router /{id} [delete]
func (a *httpApp) DeleteUser(c echo.Context) error {
	uId := c.Param("id")
	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if uId != me.Sub {
		return echo.NewHTTPError(http.StatusBadRequest, "You can only delete your own account")
	}

	err = a.userService.Delete(c.Request().Context(), me.Sub)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, comtypes.MessageRes{Message: "User deleted"})
}
