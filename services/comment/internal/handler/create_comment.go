package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/dto"
	"github.com/labstack/echo/v4"
)

type CreateCommentBody struct {
	PId  string `json:"p_id"     validate:"required"`
	Body string `json:"body"      validate:"required"`
}

func (a *httpApp) CreateComment(ctx echo.Context) error {
	var reqBody CreateCommentBody

	if err := ctx.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't find request body")
	}
	if err := ctx.Validate(reqBody); err != nil {
		return err
	}

	me, err := middleware.ParseUser(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	dc := dto.Comment{
		PId:  reqBody.PId,
		AId:  me.Sub,
		Body: reqBody.Body,
	}

	c, err := a.cs.Create(ctx.Request().Context(), dc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusCreated, c)

}
