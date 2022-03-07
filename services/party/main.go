package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/jonashiltl/sessions-backend/packages/comutils"
	"github.com/jonashiltl/sessions-backend/packages/nats"
	_ "github.com/jonashiltl/sessions-backend/services/party/docs"
	"github.com/jonashiltl/sessions-backend/services/party/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/party/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/party/internal/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Party Microservice
// @description This Microservice manages Party entities
// @version 1.0
// @host localhost:8082
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	nc, err := nats.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()

	sess, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	dao := repository.NewDAO(&sess)

	partyService := service.NewPartyServie(dao)

	httpApp := handler.NewHttpApp(partyService, nc)

	e := echo.New()

	e.Validator = &comutils.CustomValidator{Validator: validator.New()}

	e.Use(comutils.NewLogger(e))

	e.GET("/docs/*", echoSwagger.WrapHandler)

	e.GET("/:pId", httpApp.GetParty)
	e.POST("/", httpApp.CreateParty)
	e.DELETE("/:pId", httpApp.DeleteParty)
	e.PATCH("/:pId", httpApp.UpdateParty)

	e.GET("/user/:uId", httpApp.GetByUser)
	e.PATCH("/near", httpApp.GeoSearch)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
