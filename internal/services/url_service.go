package service

import (
	"errors"
	"math/rand"
	"time"

	"github.com/yourusername/url-shortener-svc/internal/models"
	"github.com/yourusername/url-shortener-svc/internal/repository"
)

// URLService defines the interface for URL shortening
type URLService interface {
	ShortenURL(originalURL string) (string, error)
	GetOriginalURL(shortURL string) (*models.URL, error)
}

// urlService struct implementing the interface
type urlService struct {
	urlRepo repository.URLRepository
}

// NewURLService creates a new URLService
func NewURLService(repo repository.URLRepository) URLService {
	return &urlService{urlRepo: repo}
}

// ShortenURL generates a short URL and saves the original one
func (s *urlService) ShortenURL(originalURL string) (string, error) {
	shortURL := generateShortCode()

	// Check if the short URL already exists
	_, err := s.urlRepo.FindByShortURL(shortURL)
	if err == nil {
		// If found, generate a new one
		shortURL = generateShortCode()
	}

	// Create a new URL models
	url := models.URL{
		OriginalURL: originalURL,
		ShortCode:   shortURL,
		CreatedAt:   time.Now().Unix(),
	}

	// Save the URL
	_, err = s.urlRepo.SaveURL(&url)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

// GetOriginalURL retrieves the original URL by short code
func (s *urlService) GetOriginalURL(shortURL string) (*models.URL, error) {
	url, err := s.urlRepo.FindByShortURL(shortURL)
	if err != nil {
		return nil, errors.New("URL not found")
	}
	return url, nil
}

// Helper function to generate a random short code
func generateShortCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, 6) // short code of length 6
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
