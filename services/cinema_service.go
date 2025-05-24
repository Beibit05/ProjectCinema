package services

import (
	"ProjectCinema/config"
	"ProjectCinema/models"
)

func CreateCinema(cinema *models.Cinema) error {
	return config.DB.Create(cinema).Error
}

func GetAllCinemas() ([]models.Cinema, error) {
	var cinemas []models.Cinema
	err := config.DB.Find(&cinemas).Error
	return cinemas, err
}

func GetCinemaByID(id string) (*models.Cinema, error) {
	var cinema models.Cinema
	err := config.DB.First(&cinema, id).Error
	return &cinema, err
}

func UpdateCinema(cinema *models.Cinema) error {
	return config.DB.Save(cinema).Error
}

func DeleteCinema(id string) error {
	return config.DB.Delete(&models.Cinema{}, id).Error
}
