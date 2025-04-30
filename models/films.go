package models

import "time"

type Film struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	GenreID         int       `json:"genre_id"`
	DirectorID      int       `json:"director_id"`
	DurationMinutes int       `json:"duration_minutes"`
	ReleaseYear     int       `json:"release_year"`
	CreatedAt       time.Time `json:"created_at"`
}
