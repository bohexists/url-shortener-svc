package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yourusername/url-shortener-svc/internal/services"
)

// InitRoutes initializes the router
func InitRoutes(e *echo.Echo, urlService *services.URLService) {
	// Route for shortening URL
	e.POST("/shorten", func(c echo.Context) error {
		type request struct {
			URL string `json:"url" binding:"required"`
		}

		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		shortCode, err := urlService.CreateShortURL(req.URL)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to shorten URL"})
		}

		return c.JSON(http.StatusOK, map[string]string{"short_code": shortCode})
	})

	// Route for getting original URL
	e.GET("/:shortCode", func(c echo.Context) error {
		shortCode := c.Param("shortCode")

		originalURL, err := urlService.GetOriginalURL(shortCode)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
		}

		return c.Redirect(http.StatusMovedPermanently, originalURL)
	})
}
