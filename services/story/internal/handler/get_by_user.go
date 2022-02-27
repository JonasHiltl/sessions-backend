package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/services/story/internal/datastruct"
	"github.com/labstack/echo/v4"
)

func (a *httpApp) GetByUser(c echo.Context) error {
	uid := c.Param("uId")
	if uid == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid User Id")
	}

	stories, err := a.sService.GetByUser(c.Request().Context(), uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var result []datastruct.PublicStory
	for _, s := range stories {
		result = append(result, s.ToPublicStory())
	}

	return c.JSON(http.StatusOK, result)
}
