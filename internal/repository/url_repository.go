package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/yourusername/url-shortener-svc/internal/models"
	db "github.com/yourusername/url-shortener-svc/storage"
)

// URLRepository presents the URLRepository interface
type URLRepository interface {
	SaveURL(url *models.URL) (string, error)
	FindByShortURL(shortURL string) (*models.URL, error)
}

// MongoURLRepository implements the URLRepository interface
type MongoURLRepository struct{}

// NewURLRepository creates a new instance of MongoURLRepository
func NewURLRepository() URLRepository {
	return &MongoURLRepository{}
}

// SaveURL saves a new URL
func (r *MongoURLRepository) SaveURL(url models.URL) (string, error) {
	url.ID = primitive.NewObjectID().Hex()
	url.CreatedAt = time.Now().Unix()

	collection := db.MI.DB.Collection("urls")
	_, err := collection.InsertOne(context.TODO(), url)
	if err != nil {
		return "", err
	}

	return url.ID, nil
}

// FindByShortURL serves a new URL
func (r *MongoURLRepository) FindByShortURL(shortURL string) (*models.URL, error) {
	var url models.URL

	collection := db.MI.DB.Collection("urls")
	err := collection.FindOne(context.TODO(), bson.M{"short_url": shortURL}).Decode(&url)
	if err != nil {
		return nil, err
	}

	return &url, nil
}
