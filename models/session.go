package models

import "time"

type Session struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FilmID    uint      `json:"film_id"`
	CinemaID  uint      `json:"cinema_id"`
	StartTime time.Time `json:"start_time"`
	HallName  string    `json:"hall_name"`
}
