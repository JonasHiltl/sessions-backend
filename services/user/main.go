package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/jonashiltl/sessions-backend/services/user/docs"
	_ "github.com/jonashiltl/sessions-backend/services/user/ent/runtime"
	"github.com/jonashiltl/sessions-backend/services/user/internal/handlers"
	"github.com/jonashiltl/sessions-backend/services/user/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/user/internal/service"
	"github.com/jonashiltl/sessions-backend/services/user/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title User Microservice
// @description This Microservice manages user entities
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	client, err := repository.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	tokenManager := service.NewTokenManager()

	userService := service.NewUserService(client)
	authService := service.NewAuthService(client, tokenManager)
	friendService := service.NewFriendService(client)

	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "[${time_custom}] ${status} ${method} ${path} ${latency_human} ${error}\n",
		CustomTimeFormat: "02.01.2006 15:04:05",
		Output:           e.Logger.Output(),
	}))

	httpApp := handlers.NewHttpApp(userService, authService, friendService)

	e.GET("/docs/*", echoSwagger.WrapHandler)

	e.POST("/", httpApp.CreateUser)
	e.DELETE("/:id", httpApp.DeleteUser)
	e.GET("/:id", httpApp.GetUser)
	e.PATCH("/:id", httpApp.UpdateUser)

	e.GET("/me", httpApp.GetMe)
	e.GET("/username-exists/:username", httpApp.UsernameExists)

	e.POST("/auth/login", httpApp.Login)
	e.POST("/auth/register", httpApp.Register)

	e.PUT("/friend/:id", httpApp.FriendRequest)
	e.GET("/friend/search/:id", httpApp.FriendSearch)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
