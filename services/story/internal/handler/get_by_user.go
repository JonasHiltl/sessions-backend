package handler

import (
	"encoding/base64"
	"net/http"

	"github.com/jonashiltl/sessions-backend/services/story/internal/datastruct"
	"github.com/labstack/echo/v4"
)

// @Summary Get stories of a user
// @Description Returns a list of stories of a user
// @Produce json
// @Param uId path string true "User id"
// @Success 200 {object} datastruct.PagedStories
// @Failure 400 {object} echo.HTTPError
// @Router /user/{uId} [get]
func (a *httpApp) GetByUser(c echo.Context) error {
	uid := c.Param("uId")
	if uid == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid User Id")
	}

	pageQuery := c.QueryParam("nextPage")
	p, err := base64.URLEncoding.DecodeString(pageQuery)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Next Page Param")
	}

	stories, p, err := a.sService.GetByUser(c.Request().Context(), uid, p)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	nextPage := base64.URLEncoding.EncodeToString(p)

	var result []datastruct.PublicStory
	for _, s := range stories {
		result = append(result, s.ToPublicStory())
	}

	return c.JSON(http.StatusOK, datastruct.PagedStories{Stories: result, NextPage: nextPage})
}
