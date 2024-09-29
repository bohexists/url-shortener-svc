package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/bohexists/url-shortener-svc/config"
	"github.com/bohexists/url-shortener-svc/internal/handlers"
	"github.com/bohexists/url-shortener-svc/internal/repository"
	"github.com/bohexists/url-shortener-svc/internal/routers"
	"github.com/bohexists/url-shortener-svc/internal/services"
	"github.com/bohexists/url-shortener-svc/pkg/logger"
	"github.com/bohexists/url-shortener-svc/pkg/server"
	db "github.com/bohexists/url-shortener-svc/storage"
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

	// Initialize logger
	logger.InitLogger()
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize routes
	routers.InitRoutes(e, urlHandler)

	// Start server
	server.StartServer(e, ":8080")
}
