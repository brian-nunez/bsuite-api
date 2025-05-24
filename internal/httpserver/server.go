package httpserver

import (
	v1 "github.com/brian-nunez/bsuite-api/internal/handlers/v1"
	"github.com/labstack/echo/v4"
)

func Bootstrap() *echo.Echo {
	server := New().
		WithDefaultMiddleware().
		WithErrorHandler().
		WithRoutes(v1.RegisterRoutes).
		WithNotFound().
		Build()

	return server
}
