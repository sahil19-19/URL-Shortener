package models

// represents a URL object
type URL struct {
	ID          int    `json:"id,omitempty"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	CustomURL   string `json:"custom_url,omitempty"`
}
