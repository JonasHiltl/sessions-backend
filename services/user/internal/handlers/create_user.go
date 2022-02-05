package handlers

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"github.com/labstack/echo/v4"
)

// @Summary Create a user
// @Description Saves a user in our DB
// @Tags CRUD
// @Accept json
// @Produce json
// @Param Body body datastruct.RequestUser true "The body to create a user"
// @Success 201 {object} datastruct.PublicUser
// @Failure 400 {object} echo.HTTPError
// @Router / [post]
func (a *httpApp) CreateUser(c echo.Context) error {
	var reqBody datastruct.RequestUser
	if err := c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't find request body")
	}

	if err := c.Validate(reqBody); err != nil {
		return err
	}

	u, err := a.userService.Create(c.Request().Context(), reqBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, u.ToPublicProfile())
}
