package handler

import (
	"github.com/jonashiltl/sessions-backend/services/user/internal/service"
	"github.com/labstack/echo/v4"
)

type HttpApp interface {
	CreateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	GetUser(c echo.Context) error
	GetMe(c echo.Context) error
	UpdateUser(c echo.Context) error
	UsernameExists(c echo.Context) error
	Login(c echo.Context) error
	Register(c echo.Context) error
	FriendRequest(c echo.Context) error
	FriendSearch(c echo.Context) error
	GetFriends(c echo.Context) error
}

type httpApp struct {
	userService   service.UserService
	authService   service.AuthService
	friendService service.FriendService
	uploadService service.UploadService
}

func NewHttpApp(userService service.UserService, authService service.AuthService, friendService service.FriendService, uploadService service.UploadService) HttpApp {
	return &httpApp{
		authService:   authService,
		userService:   userService,
		friendService: friendService,
		uploadService: uploadService,
	}
}
