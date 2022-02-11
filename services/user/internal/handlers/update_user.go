package handlers

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"github.com/labstack/echo/v4"
)

// @Summary Update a user
// @Description Updates user with provided values
// @Tags CRUD
// @Accept json
// @Produce json
// @Param Body body datastruct.RequestUser true "The body to create a thing"
// @Success 200 {object} datastruct.PublicUser
// @Failure 400 {object} echo.HTTPError
// @Router /{id} [patch]
func (a *httpApp) UpdateUser(c echo.Context) error {
	uId := c.Param("id")
	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if uId != me.Sub {
		return echo.NewHTTPError(http.StatusBadRequest, "You can only update your own information")
	}

	var reqBody datastruct.RequestUser
	if err := c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't find request body")
	}

	u, err := a.userService.Update(c.Request().Context(), me.Sub, reqBody)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	friendCount := a.userService.CountFriends(c.Request().Context(), u.ID)

	return c.JSON(http.StatusOK, u.ToPublicProfile().AddCount(friendCount))
}
