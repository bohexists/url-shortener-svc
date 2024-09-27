package model

import (
	"time"
)

// URL presents a short URL
type URL struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	OriginalURL string    `bson:"original_url" json:"original_url"`
	ShortURL    string    `bson:"short_url" json:"short_url"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	ExpiresAt   time.Time `bson:"expires_at" json:"expires_at,omitempty"`
}
