package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/labstack/echo/v4"
)

func (a *httpApp) DeleteComment(c echo.Context) error {
	cId := c.Param("cId")
	pId := c.Param("pId")

	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = a.cs.Delete(c.Request().Context(), me.Sub, pId, cId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, comtypes.MessageRes{Message: "Comment removed"})
}
