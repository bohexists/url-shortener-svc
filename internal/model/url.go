package model

// URL presents a short URL
type URL struct {
	ID          string `bson:"_id,omitempty"`
	OriginalURL string `bson:"original_url"`
	ShortCode   string `bson:"short_code"`
	CreatedAt   int64  `bson:"created_at"`
}
