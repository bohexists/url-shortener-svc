package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yourusername/url-shortener-svc/internal/handlers"
	"github.com/yourusername/url-shortener-svc/internal/repository"
	"github.com/yourusername/url-shortener-svc/internal/routers"
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

	// Initialize repository, service, and handler
	urlRepo := repository.NewURLRepository()
	urlService := services.NewURLService(*urlRepo)
	urlHandler := handlers.NewURLHandler(urlService)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize routes
	routers.InitRoutes(e, urlHandler)

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
