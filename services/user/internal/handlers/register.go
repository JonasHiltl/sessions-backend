package handlers

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"github.com/labstack/echo/v4"
)

// @Summary Register a new User
// @Description Saves a new User in the Database and returns auth token
// @Tags auth
// @Accept json
// @Produce json
// @Param Body body datastruct.RequestUser true "The body to create a user"
// @Success 201 {object} datastruct.AuthRes
// @Failure 400 {object} echo.HTTPError
// @Router /auth/register [post]
func (a *httpApp) Register(c echo.Context) error {
	var reqBody datastruct.RequestUser
	if err := c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't find request body")
	}

	if err := c.Validate(reqBody); err != nil {
		return err
	}

	token, err := a.authService.Register(c.Request().Context(), reqBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := datastruct.AuthRes{
		Token:   token,
		Message: "Successfully Registered",
	}

	return c.JSON(http.StatusCreated, res)
}
