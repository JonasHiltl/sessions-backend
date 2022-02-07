package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/jonashiltl/sessions-backend/packages/comutils"
	"github.com/jonashiltl/sessions-backend/services/party/internal/handlers"
	"github.com/jonashiltl/sessions-backend/services/party/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/party/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title User Microservice
// @description This Microservice manages user entities
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	partyCol, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	dao := repository.NewDAO(partyCol)

	partyService := service.NewPartyServie(dao)

	httpApp := handlers.NewHttpApp(partyService)

	e := echo.New()

	e.Validator = &comutils.CustomValidator{Validator: validator.New()}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "[${time_custom}] ${status} ${method} ${path} ${latency_human} ${error}\n",
		CustomTimeFormat: "02.01.2006 15:04:05",
		Output:           e.Logger.Output(),
	}))

	e.GET("/:id", httpApp.GetParty)
	e.POST("/", httpApp.CreateParty)
	e.DELETE("/:id", httpApp.DeleteParty)
	e.PATCH("/:id", httpApp.UpdateParty)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
