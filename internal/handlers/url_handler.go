package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	services "github.com/yourusername/url-shortener-svc/internal/services"
)

type URLHandler struct {
	urlService services.URLServiceinterface
}

func NewURLHandler(urlService services.URLServiceinterface) *URLHandler {
	return &URLHandler{
		urlService: urlService,
	}
}

// ShortenURL handles the request for creating a short URL
func (h *URLHandler) ShortenURL(c echo.Context) error {
	var request struct {
		OriginalURL string `json:"original_url" validate:"required,url"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := c.Validate(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid URL format"})
	}

	// Shorten the URL
	shortURL, err := h.urlService.ShortenURL(request.OriginalURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to shorten URL"})
	}

	return c.JSON(http.StatusOK, map[string]string{"short_url": shortURL})
}

// RedirectToOriginalURL handles the redirection to the original URL
func (h *URLHandler) RedirectToOriginalURL(c echo.Context) error {
	shortURL := c.Param("shortURL")

	// Get the original URL
	url, err := h.urlService.GetOriginalURL(shortURL)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
	}

	return c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}
