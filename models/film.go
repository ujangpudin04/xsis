package models

import "time"

type Film struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Title       string    `json:"title" gorm:"type: varchar(255)"`
	Description string    `json:"description"`
	Rating      float64   `json:"rating" gorm:"type: varchar(255)"`
	Image       string    `json:"image" gorm:"type: varchar(255)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// for association relation with another table (user)
type FilmResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" `
	Description string    `json:"description"`
	Rating      float64   `json:"rating"`
	Image       string    `json:"image" `
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (FilmResponse) TableName() string {
	return "films"
}
