package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	routes "github.com/yourusername/url-shortener-svc/internal/routers"
	"github.com/yourusername/url-shortener-svc/internal/services"
	"log"

	"github.com/yourusername/url-shortener-svc/config"
	db "github.com/yourusername/url-shortener-svc/storage"
)

func main() {

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to MongoDB
	err := db.Connect(cfg.MongoURI, "url_shortener")
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// initialize URL service
	urlService := services.URLService{}

	// initialize routes
	routes.InitRoutes(e, &urlService)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
