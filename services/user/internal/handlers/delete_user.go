package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/jonashiltl/sessions-backend/services/user/internal/handlers/middleware"
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

	uUUID, err := uuid.Parse(uId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if uUUID != me.Sub {
		return echo.NewHTTPError(http.StatusBadRequest, "You can only delete your own account")
	}

	err = a.userService.Delete(c.Request().Context(), uUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, comtypes.MessageRes{Message: "User deleted"})
}
