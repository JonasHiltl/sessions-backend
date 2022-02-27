package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/labstack/echo/v4"
)

// @Summary Get current user
// @Description Gets the user information of currently logged in user
// @Accept json
// @Produce json
// @Param jwt_payload header string true "Base64 encoded JWT Payload"
// @Success 200 {object} datastruct.PublicUser
// @Failure 400 {object} echo.HTTPError
// @Router /me [get]
func (a *httpApp) GetMe(c echo.Context) error {
	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	u, err := a.userService.GetById(c.Request().Context(), me.Sub)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, u.ToPublicProfile())
}
