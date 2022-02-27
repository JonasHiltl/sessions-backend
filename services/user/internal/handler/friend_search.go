package handler

import (
	"net/http"
	"strconv"

	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
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
	acceptedB, err := strconv.ParseBool(accepted)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid accepted parameter")
	}

	uId := c.Param("id")

	friends, err := a.friendService.Search(c.Request().Context(), uId, query, acceptedB)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var publicFriends []datastruct.PublicUser

	for _, user := range friends {
		publicFriends = append(publicFriends, user.ToPublicProfile())
	}

	return c.JSON(http.StatusOK, publicFriends)
}
