package handlers

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// @Summary Search friends of user
// @Description Search friends by username, first name or last name
// @Tags friend
// @Accept json
// @Produce json
// @Param id path string true "Id of user which friends to search"
// @Param accepted query bool false "If only accepted friends to search"
// @Param query query string true "Query string"
// @Success 200 {array} datastruct.PublicUser
// @Failure 400 {object} echo.HTTPError
// @Router /friend/search/{id} [get]
func (a *httpApp) FriendSearch(c echo.Context) error {
	query := c.QueryParam("query")
	accepted := c.QueryParam("accepted")
	acceptedBool, err := strconv.ParseBool(accepted)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid accepted parameter")
	}

	userId := c.Param("id")

	uuid, err := uuid.Parse(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	friends, err := a.friendService.Search(c.Request().Context(), uuid, query, acceptedBool)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, friends)
}
