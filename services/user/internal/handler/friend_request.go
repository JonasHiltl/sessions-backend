package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/labstack/echo/v4"
)

// @Summary Send friend request
// @Description Sends friend request to user with id from params
// @Tags friend
// @Accept json
// @Produce json
// @Param id path string true "Friend Id"
// @Success 201 {object} datastruct.MessageRes
// @Failure 400 {object} echo.HTTPError
// @Router /friend/{id} [put]
func (a *httpApp) FriendRequest(c echo.Context) error {
	fId := c.Param("id")
	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if fId == me.Sub {
		return echo.NewHTTPError(http.StatusBadRequest, "You can't add yourself")
	}

	err = a.friendService.FriendRequest(c.Request().Context(), fId, me.Sub)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, comtypes.MessageRes{Message: "Friend request send"})
}
