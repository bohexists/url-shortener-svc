package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// URL presents a short URL
type URL struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	OriginalURL string             `bson:"original_url"`
	ShortCode   string             `bson:"short_code"`
	CreatedAt   int64              `bson:"created_at"`
}
