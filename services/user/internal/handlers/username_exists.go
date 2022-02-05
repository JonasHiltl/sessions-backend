package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Check Username availability
// @Description Check if username is already taken
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} bool
// @Failure 400 {object} echo.HTTPError
// @Router /username-exists/{username} [get]
func (a *httpApp) UsernameExists(c echo.Context) error {
	username := c.Param("username")

	usernameExists, err := a.userService.UsernameExists(c.Request().Context(), username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, usernameExists)
}
