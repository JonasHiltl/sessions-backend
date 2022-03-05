package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/services/story/internal/datastruct"
	"github.com/labstack/echo/v4"
)

// @Summary Get stories of a party
// @Description Returns a list of stories of a party
// @Produce json
// @Param pId path string true "Party id"
// @Success 200 {array} datastruct.PublicStory
// @Failure 400 {object} echo.HTTPError
// @Router /party/{pId} [get]
func (a *httpApp) GetByParty(c echo.Context) error {
	pId := c.Param("pId")
	if pId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Party Id")
	}

	stories, err := a.sService.GetByParty(c.Request().Context(), pId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var result []datastruct.PublicStory
	for _, s := range stories {
		result = append(result, s.ToPublicStory())
	}

	return c.JSON(http.StatusOK, result)
}
