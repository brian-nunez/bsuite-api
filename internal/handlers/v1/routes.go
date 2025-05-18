package v1

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	v1Group := e.Group("/api/v1")
	v1Group.GET("/health", HealthHandler)
}
