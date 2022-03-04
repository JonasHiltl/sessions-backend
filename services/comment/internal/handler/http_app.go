package handler

import (
	"github.com/jonashiltl/sessions-backend/services/comment/internal/service"
	"github.com/labstack/echo/v4"
)

type HttpApp interface {
	CreateComment(c echo.Context) error
	DeleteComment(c echo.Context) error
	GetCommentByParty(c echo.Context) error
	GetCommentByPartyUser(c echo.Context) error
}

type httpApp struct {
	cs service.CommentService
}

func NewHttpApp(cs service.CommentService) HttpApp {
	return &httpApp{
		cs: cs,
	}
}
