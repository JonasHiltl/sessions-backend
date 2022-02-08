package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *httpApp) SearchParty(c echo.Context) error {
	query := c.QueryParam("query")
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}

	p, err := strconv.Atoi(page)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result, err := a.partyService.Search(c.Request().Context(), query, p)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}
