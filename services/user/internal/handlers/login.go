package handlers

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"github.com/labstack/echo/v4"
)

// @Summary Login
// @Description Login with credentials
// @Tags auth
// @Accept json
// @Produce json
// @Param Body body datastruct.LoginBody true "The required credentials"
// @Success 200 {object} datastruct.AuthRes
// @Failure 400 {object} echo.HTTPError
// @Router /auth/login [post]
func (a *httpApp) Login(c echo.Context) error {
	var reqBody datastruct.LoginBody
	if err := c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't find request body")
	}

	if err := c.Validate(reqBody); err != nil {
		return err
	}

	token, err := a.authService.Login(c.Request().Context(), reqBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := datastruct.AuthRes{
		Token:   token,
		Message: "Successfully Logged in",
	}

	return c.JSON(http.StatusOK, res)
}
