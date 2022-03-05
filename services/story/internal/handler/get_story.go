package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Get a Story
// @Description Get a Story by it's id
// @Tags CRUD
// @Accept json
// @Produce json
// @Param sId path string true "Story Id"
// @Success 200 {object} datastruct.PublicStory
// @Failure 400 {object} echo.HTTPError
// @Router /{sId} [get]
func (a *httpApp) GetStory(c echo.Context) error {
	sId := c.Param("sId")

	s, err := a.sService.Get(c.Request().Context(), sId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, s.ToPublicStory())
}
