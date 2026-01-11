package http

import (
	"github.com/gooveapp/goove/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RegisterRoutes sets up the HTTP routes for the Echo server.
func RegisterRoutes(e *echo.Echo) {
	homeHandler := handlers.HomeHandler{}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("/static"))

	e.Static("/static", "static")

	e.GET("/", homeHandler.HandleHome)

	e.GET("/records",)
	e.GET("/settings",)
}
