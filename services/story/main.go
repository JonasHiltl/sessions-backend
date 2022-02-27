package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/jonashiltl/sessions-backend/packages/comutils"
	"github.com/jonashiltl/sessions-backend/services/story/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/story/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/story/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sess, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	dao := repository.NewDAO(&sess)

	storyService := service.NewStoryServie(dao)

	httpApp := handler.NewHttpApp(storyService)

	e := echo.New()

	e.Validator = &comutils.CustomValidator{Validator: validator.New()}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "[${time_custom}] ${status} ${method} ${path} ${latency_human} ${error}\n",
		CustomTimeFormat: "02.01.2006 15:04:05",
		Output:           e.Logger.Output(),
	}))

	e.GET("/:sId", httpApp.GetStory)
	e.POST("/", httpApp.CreateStory)
	e.DELETE("/:sId", httpApp.DeleteStory)
	e.GET("/user/:uId", httpApp.GetByUser)
	e.GET("/party/:pId", httpApp.GetByParty)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
