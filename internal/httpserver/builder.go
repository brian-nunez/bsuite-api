package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ServerBuilder struct {
	e *echo.Echo
}

func New() *ServerBuilder {
	return &ServerBuilder{
		e: echo.New(),
	}
}

func (b *ServerBuilder) WithDefaultMiddleware() *ServerBuilder {
	b.e.Use(middleware.Recover())
	b.e.Use(middleware.RequestID())
	b.e.Use(middleware.CORS())
	b.e.Use(middleware.Logger())

	return b
}

func (b *ServerBuilder) WithRoutes(register func(e *echo.Echo)) *ServerBuilder {
	register(b.e)
	return b
}

func (b *ServerBuilder) WithErrorHandler() *ServerBuilder {
	b.e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := echo.ErrInternalServerError.Code
		message := "Internal Server Error"

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			message = he.Message.(string)
		}

		c.Logger().Error(err)

		if !c.Response().Committed {
			_ = c.JSON(code, map[string]any{
				"error":   true,
				"message": message,
			})
		}
	}

	return b
}

func (b *ServerBuilder) Build() *echo.Echo {
	return b.e
}
