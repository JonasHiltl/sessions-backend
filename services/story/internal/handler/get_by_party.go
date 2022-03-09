package handler

import (
	"encoding/base64"
	"net/http"

	"github.com/jonashiltl/sessions-backend/services/story/internal/datastruct"
	"github.com/labstack/echo/v4"
)

// @Summary Get stories of a party
// @Description Returns a list of stories of a party
// @Produce json
// @Param pId path string true "Party id"
// @Success 200 {object} datastruct.PagedStories
// @Failure 400 {object} echo.HTTPError
// @Router /party/{pId} [get]
func (a *httpApp) GetByParty(c echo.Context) error {
	pId := c.Param("pId")
	if pId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Party Id")
	}

	pageQuery := c.QueryParam("nextPage")
	p, err := base64.URLEncoding.DecodeString(pageQuery)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Next Page Param")
	}

	stories, p, err := a.sService.GetByParty(c.Request().Context(), pId, p)
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
