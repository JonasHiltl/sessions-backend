package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/user/internal/handlers/middleware"
	"github.com/labstack/echo/v4"
)

// FriendRequest goDoc
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
	friendId := c.Param("id")
	friendUUID, err := uuid.Parse(friendId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if friendUUID == me.Sub {
		return echo.NewHTTPError(http.StatusBadRequest, "You can't add yourself")
	}

	a.friendService.FriendRequest(c.Request().Context(), friendUUID, me.Sub)

	return c.JSON(http.StatusCreated, datastruct.MessageRes{Message: "Friend request send"})
}
