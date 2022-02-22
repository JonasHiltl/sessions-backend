package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Get user
// @Description Gets the user information by id user
// @Tags CRUD
// @Accept json
// @Produce json
// @Param id path string true "User Id"
// @Success 200 {object} datastruct.PublicUser
// @Failure 400 {object} echo.HTTPError
// @Router /{id} [get]
func (a *httpApp) GetUser(c echo.Context) error {
	uId := c.Param("id")

	u, err := a.userService.GetById(c.Request().Context(), uId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, u.ToPublicProfile())
}
