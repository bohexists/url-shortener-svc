package dto

// ShortenURLRequestDTO defines the structure of the request to shorten a URL
type ShortenURLRequestDTO struct {
	OriginalURL string `json:"original_url" validate:"required,url"`
}

// ShortenURLResponseDTO defines the structure of the response after shortening a URL
type ShortenURLResponseDTO struct {
	ShortURL string `json:"short_url"`
}

// GetOriginalURLResponseDTO defines the structure of the response when fetching the original URL
type GetOriginalURLResponseDTO struct {
	OriginalURL string `json:"original_url"`
}
