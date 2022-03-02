package handler

import (
	"github.com/jonashiltl/sessions-backend/services/story/internal/service"
	"github.com/labstack/echo/v4"
)

type HttpApp interface {
	CreateStory(c echo.Context) error
	GetStory(c echo.Context) error
	DeleteStory(c echo.Context) error
	GetByUser(c echo.Context) error
	GetByParty(c echo.Context) error
	PresignURL(c echo.Context) error
}

type httpApp struct {
	sService service.StoryService
	us       service.UploadService
}

func NewHttpApp(sService service.StoryService, us service.UploadService) HttpApp {
	return &httpApp{
		sService: sService,
		us:       us,
	}
}
