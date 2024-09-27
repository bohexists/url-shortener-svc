package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

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

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the URL Shortener Service!")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
