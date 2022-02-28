package comutils

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewLogger(e *echo.Echo) echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "[${time_custom}] ${status} ${method} ${path} ${latency_human} ${error}\n",
		CustomTimeFormat: "02.01.2006 15:04:05",
		Output:           e.Logger.Output(),
	})
}
