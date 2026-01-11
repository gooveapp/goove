package http

import (
	"github.com/gooveapp/goove/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	homeHandler := handlers.HomeHandler{}
	recordHandler := handlers.RecordHandler{}
	settingsHandler := handlers.SettingsHandler{}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("/static"))

	e.Static("/static", "static")

	e.GET("/", homeHandler.HandleHome)

	records := e.Group("records")
	records.GET("/", recordHandler.HandleRecords)
	records.GET("/new", recordHandler.HandleRecordCreatePage)

	settings := e.Group("settings")
	settings.GET("/", settingsHandler.HandleSettings)
 }
