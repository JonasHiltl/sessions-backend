package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/jonashiltl/sessions-backend/services/user/internal/handlers/middleware"
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
	fUUID, err := uuid.Parse(fId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if fUUID == me.Sub {
		return echo.NewHTTPError(http.StatusBadRequest, "You can't add yourself")
	}

	err = a.friendService.FriendRequest(c.Request().Context(), fUUID, me.Sub)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, comtypes.MessageRes{Message: "Friend request send"})
}
