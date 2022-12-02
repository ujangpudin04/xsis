package filmsdto

import "time"

type FilmResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      float64   `json:"rating"`
	UpdatedAt   time.Time `json:"updated_at"`
}
