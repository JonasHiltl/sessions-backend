package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/jonashiltl/sessions-backend/packages/comutils"
	_ "github.com/jonashiltl/sessions-backend/services/party/docs"
	"github.com/jonashiltl/sessions-backend/services/party/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/party/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/party/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	sess, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	dao := repository.NewDAO(&sess)

	partyService := service.NewPartyServie(dao)

	httpApp := handler.NewHttpApp(partyService)

	e := echo.New()

	e.Validator = &comutils.CustomValidator{Validator: validator.New()}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "[${time_custom}] ${status} ${method} ${path} ${latency_human} ${error}\n",
		CustomTimeFormat: "02.01.2006 15:04:05",
		Output:           e.Logger.Output(),
	}))

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
