package service

import (
	"ProjectCinema/config"
	"ProjectCinema/models"
	"strconv"
)

func FetchFilmsFromDB(page, limit int, genre string) ([]models.Film, error) {
	var films []models.Film
	query := config.DB.Model(models.Film{})

	if genre != "" {
		genreInt, _ := strconv.Atoi(genre)
		query = query.Where("genre=?", genreInt)
	}
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&films).Error; err != nil {
		return nil, err
	}
	return films, nil
}
