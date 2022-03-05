package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/jonashiltl/sessions-backend/packages/comutils"
	_ "github.com/jonashiltl/sessions-backend/services/story/docs"
	"github.com/jonashiltl/sessions-backend/services/story/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/story/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/story/internal/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
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

	sService := service.NewStoryServie(dao)
	us := service.NewUploadService()

	httpApp := handler.NewHttpApp(sService, us)

	e := echo.New()

	e.Validator = &comutils.CustomValidator{Validator: validator.New()}

	e.GET("/docs/*", echoSwagger.WrapHandler)

	e.Use(comutils.NewLogger(e))

	e.GET("/:sId", httpApp.GetStory)
	e.POST("/", httpApp.CreateStory)
	e.DELETE("/:sId", httpApp.DeleteStory)
	e.GET("/user/:uId", httpApp.GetByUser)
	e.GET("/party/:pId", httpApp.GetByParty)

	e.GET("/presign/:key", httpApp.PresignURL)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
