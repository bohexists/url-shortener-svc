package services

import (
	"context"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/yourusername/url-shortener-svc/internal/model"
	db "github.com/yourusername/url-shortener-svc/storage"
)

// URLService serves a new URL
type URLService struct{}

// CreateShortURL creates a new URL
func (s *URLService) CreateShortURL(originalURL string) (string, error) {
	// Generate short code
	shortCode := generateShortCode()

	url := model.URL{
		ID:          primitive.NewObjectID(),
		OriginalURL: originalURL,
		ShortCode:   shortCode,
		CreatedAt:   time.Now().Unix(),
	}

	// Place URL in database
	_, err := db.MI.DB.Collection("urls").InsertOne(context.Background(), url)
	if err != nil {
		return "", err
	}

	return shortCode, nil
}

// GetOriginalURL retrieves a new URL
func (s *URLService) GetOriginalURL(shortCode string) (string, error) {
	var url model.URL

	// Find URL in database
	err := db.MI.DB.Collection("urls").FindOne(context.Background(), bson.M{"short_code": shortCode}).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.OriginalURL, nil
}

// generateShortCode generates a new short code
func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
