package handlers

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
}

type httpApp struct {
	userService   service.UserService
	authService   service.AuthService
	friendService service.FriendService
}

func NewHttpApp(userService service.UserService, authService service.AuthService, friendService service.FriendService) HttpApp {
	return &httpApp{
		authService:   authService,
		userService:   userService,
		friendService: friendService,
	}
}
