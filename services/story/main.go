package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/jonashiltl/sessions-backend/services/story/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "[${time_custom}] ${status} ${method} ${path} ${latency_human} ${error}\n",
		CustomTimeFormat: "02.01.2006 15:04:05",
		Output:           e.Logger.Output(),
	}))
}
