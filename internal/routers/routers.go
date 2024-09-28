package routers

import (
	"github.com/labstack/echo/v4"

	"github.com/yourusername/url-shortener-svc/internal/handlers"
)

// InitRoutes initializes the router
func InitRoutes(e *echo.Echo, urlHandler *handlers.URLHandler) {
	// Route for shortening URL
	e.POST("/shorten", urlHandler.ShortenURL)

	// Route for getting original URL
	e.GET("/:shortCode", urlHandler.RedirectToOriginalURL)
}
