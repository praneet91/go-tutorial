package services

import (
	"fmt"
	internal "go-tutorial/internal/model"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func InitAuthService(db *gorm.DB) *AuthService {

	db.AutoMigrate(&internal.User{})
	return &AuthService{db: db}
}

func (as *AuthService) Login(email *string, password *string) (*internal.User, error) {
	if (email == nil) || (password == nil) {
		return nil, fmt.Errorf("email and password are required")
	}

	var user internal.User

	if err := as.db.Find(&user, "email = ? AND password = ?", *email, *password).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (as *AuthService) Register(email *string, password *string) (*internal.User, error) {
	if (email == nil) || (password == nil) {
		return nil, fmt.Errorf("email and password are required")
	}

	user := &internal.User{
		Email:    *email,
		Password: *password,
	}

	if err := as.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
