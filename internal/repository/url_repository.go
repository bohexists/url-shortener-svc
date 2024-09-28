package repository

import (
	"context"
	"github.com/yourusername/url-shortener-svc/internal/models"
	db "github.com/yourusername/url-shortener-svc/storage"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLRepository struct {
	Collection *mongo.Collection
}

func NewURLRepository() *URLRepository {
	return &URLRepository{
		Collection: db.MI.DB.Collection("urls"),
	}
}

// SaveURL saves a new URL
func (r *URLRepository) SaveURL(url *models.URL) error {
	url.CreatedAt = time.Now().Unix()
	_, err := r.Collection.InsertOne(context.TODO(), url)
	return err
}

// FindByShortURL finds a URL by its shortened version
func (r *URLRepository) FindByShortURL(shortURL string) (*models.URL, error) {
	var url models.URL
	err := r.Collection.FindOne(context.TODO(), bson.M{"short_url": shortURL}).Decode(&url)
	if err != nil {
		return nil, err
	}
	return &url, nil
}
