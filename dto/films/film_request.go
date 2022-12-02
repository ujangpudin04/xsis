package filmsdto

import "time"

type CreateFilmRequest struct {
	Title       string  `json:"title" form:"title" validate:"required" `
	Description string  `json:"description" form:"description" validate:"required"`
	Rating      float64 `json:"rating" form:"rating"`
	Image       string  `json:"image" form:"image" `
}

type UpdateFilmRequest struct {
	Title       string    `json:"title" form:"title" `
	Description string    `json:"description" form:"description" `
	Rating      float64   `json:"rating" form:"rating" `
	Image       string    `json:"image" form:"image"`
	UpdatedAt   time.Time `json:"update_at"`
}
