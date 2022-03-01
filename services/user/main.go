package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/jonashiltl/sessions-backend/packages/comutils"
	_ "github.com/jonashiltl/sessions-backend/services/user/docs"
	_ "github.com/jonashiltl/sessions-backend/services/user/ent/runtime"
	"github.com/jonashiltl/sessions-backend/services/user/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/user/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/user/internal/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title User Microservice
// @description This Microservice manages user entities
// @version 1.0
// @host localhost:8081
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := repository.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	tokenManager := service.NewTokenManager()

	userService := service.NewUserService(client)
	authService := service.NewAuthService(client, tokenManager)
	friendService := service.NewFriendService(client)
	uploadService := service.NewUploadService()

	e := echo.New()
	e.Validator = &comutils.CustomValidator{Validator: validator.New()}

	e.Use(comutils.NewLogger(e))

	httpApp := handler.NewHttpApp(userService, authService, friendService, uploadService)

	e.GET("/docs/*", echoSwagger.WrapHandler)

	e.POST("/", httpApp.CreateUser)
	e.DELETE("/:id", httpApp.DeleteUser)
	e.GET("/:id", httpApp.GetUser)
	e.PATCH("/:id", httpApp.UpdateUser)

	e.GET("/me", httpApp.GetMe)
	e.GET("/username-exists/:username", httpApp.UsernameExists)

	e.POST("/auth/login", httpApp.Login)
	e.POST("/auth/register", httpApp.Register)

	e.GET("/friend/:id", httpApp.GetFriends)
	e.PUT("/friend/:id", httpApp.FriendRequest)
	e.GET("/friend/search/:id", httpApp.FriendSearch)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
