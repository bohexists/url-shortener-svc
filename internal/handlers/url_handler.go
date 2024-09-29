package handlers

import (
	"github.com/bohexists/url-shortener-svc/pkg/logger"
	"go.uber.org/zap"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bohexists/url-shortener-svc/internal/dto"
	"github.com/bohexists/url-shortener-svc/internal/services"
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
	var request dto.ShortenURLRequestDTO

	logger.Logger.Info("Received request to shorten URL")

	if err := c.Bind(&request); err != nil {
		logger.Logger.Error("Failed to bind request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := c.Validate(&request); err != nil {
		logger.Logger.Error("Invalid URL format", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid URL format"})
	}

	// Shorten the URL
	shortURL, err := h.urlService.ShortenURL(request.OriginalURL)
	if err != nil {
		logger.Logger.Error("Failed to shorten URL", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to shorten URL"})
	}

	logger.Logger.Info("Successfully shortened URL", zap.String("shortURL", shortURL))

	response := dto.ShortenURLResponseDTO{
		ShortURL: shortURL,
	}
	return c.JSON(http.StatusOK, response)
}

// RedirectToOriginalURL handles the redirection to the original URL
func (h *URLHandler) RedirectToOriginalURL(c echo.Context) error {
	shortURL := c.Param("shortURL")

	logger.Logger.Info("Received request to redirect", zap.String("shortURL", shortURL))

	// Get the original URL
	url, err := h.urlService.GetOriginalURL(shortURL)
	if err != nil {
		logger.Logger.Error("URL not found", zap.Error(err))
		return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
	}

	logger.Logger.Info("Redirecting to original URL", zap.String("originalURL", url.OriginalURL))

	// Redirect to the original URL
	response := dto.GetOriginalURLResponseDTO{
		OriginalURL: url.OriginalURL,
	}
	return c.JSON(http.StatusOK, response)
}
