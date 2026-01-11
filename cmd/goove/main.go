package main

import (
	"log"

	"github.com/gooveapp/goove/internal/config"
	"github.com/gooveapp/goove/internal/database"
	"github.com/gooveapp/goove/internal/http"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	cfg := config.FromEnv()

	// Initialize database
	db, err := database.New(cfg.DatabasePath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize Echo web server
	app := echo.New()

	http.RegisterRoutes(app)

	log.Println("Goove is running on http://localhost:3000")
	if err := app.Start(":3000"); err != nil {
		log.Fatal(err)
	}

}
