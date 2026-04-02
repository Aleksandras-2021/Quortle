package services

import (
	models "Quortle/internal/models"

	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(username, password string) error {
	if username == "" || password == "" {
		return errors.New("username and password required")
	}

	user := models.User{
		Username:     username,
		PasswordHash: password,
	}

	return s.db.Create(&user).Error
}

func (s *UserService) GetUser(username string) (*models.User, error) {
	if username == "" {
		return nil, errors.New("username required")
	}

	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
