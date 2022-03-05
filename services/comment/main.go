package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/jonashiltl/sessions-backend/packages/comutils"
	_ "github.com/jonashiltl/sessions-backend/services/comment/docs"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/service"
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

	cs := service.NewCommentServie(dao)

	httpApp := handler.NewHttpApp(cs)

	e := echo.New()

	e.Validator = &comutils.CustomValidator{Validator: validator.New()}

	e.GET("/docs/*", echoSwagger.WrapHandler)

	e.Use(comutils.NewLogger(e))

	e.POST("/", httpApp.CreateComment)
	e.DELETE("/:pId/:cId", httpApp.DeleteComment)

	e.GET("/party/:pId", httpApp.GetCommentByParty)
	// TODO: currently not implementable with schema
	//e.GET("/party/:pId/user/:uId", httpApp.GetCommentByPartyUser)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
