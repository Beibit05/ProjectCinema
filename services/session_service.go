package services

import (
	"ProjectCinema/config"
	"ProjectCinema/models"
)

func CreateSession(session *models.Session) error {
	return config.DB.Create(session).Error
}

func GetAllSessions() ([]models.Session, error) {
	var sessions []models.Session
	err := config.DB.Find(&sessions).Error
	return sessions, err
}

func GetSessionByID(id string) (*models.Session, error) {
	var session models.Session
	err := config.DB.First(&session, id).Error
	return &session, err
}

func UpdateSession(session *models.Session) error {
	return config.DB.Save(session).Error
}

func DeleteSession(id string) error {
	return config.DB.Delete(&models.Session{}, id).Error
}
