package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *httpApp) GetCommentByParty(ctx echo.Context) error {
	pId := ctx.Param("pId")

	c, err := a.cs.GetByParty(ctx.Request().Context(), pId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, c)
}
