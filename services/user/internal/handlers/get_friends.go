package handlers

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// @Summary Get all friends
// @Description Returns a list of friends of a user
// @Tags friend
// @Accept json
// @Produce json
// @Param id path string true "Id of the user to get friends of"
// @Param offset query int false "Offset list of friends"
// @Param limit query int false "Limits list of friends"
// @Success 200 {array} datastruct.PublicUser
// @Failure 400 {object} echo.HTTPError
// @Router /friend/{id} [get]
func (a *httpApp) GetFriends(c echo.Context) error {
	userId := c.Param("id")

	uuid, err := uuid.Parse(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	offset := c.QueryParam("offset")
	limit := c.QueryParam("limit")

	offsetNum, _ := strconv.Atoi(offset)
	limitNum, _ := strconv.Atoi(limit)

	friends, err := a.friendService.Get(c.Request().Context(), uuid, offsetNum, limitNum)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, friends)
}
