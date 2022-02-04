package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/user/internal/handlers/middleware"
	"github.com/labstack/echo/v4"
)

// CreateUser goDoc
// @Summary Update a user
// @Description Updates user with provided values
// @Tags CRUD
// @Accept json
// @Produce json
// @Param Body body datastruct.User true "The body to create a thing"
// @Success 201 {object} datastruct.PublicUser
// @Failure 400 {object} echo.HTTPError
// @Router /{id} [patch]
func (a *httpApp) UpdateUser(c echo.Context) error {
	user, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		fmt.Println(uuid)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if uuid != user.Sub {
		return echo.NewHTTPError(http.StatusBadRequest, "You can only update your own information")
	}

	var reqBody datastruct.RequestUser
	if err := c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't find request body")
	}

	u, err := a.userService.Update(c.Request().Context(), user.Sub, reqBody)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, u)
}
