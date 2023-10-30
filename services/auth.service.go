package services

import (
	"Auth-API/models"
)

type AuthService interface {
	Register(user *models.UserModel) error
	// Login(user *models.UserModel) error
	GetUser(Username string) (models.UserModel, error)
}

type AuthServiceImpl struct {
	UserService UserService
}

func NewAuthService(userService UserService) *AuthServiceImpl {
	return &AuthServiceImpl{UserService: userService}
}

func (authService *AuthServiceImpl) Register(user *models.UserModel) models.UserModel {
	User := authService.UserService.Create(*user)
	return User
}

func (authService *AuthServiceImpl) GetUser(Username string) (models.UserModel, error) {
	// User := authService.UserService.GetByEmail(Username)
	User,err := authService.UserService.GetByEmail(Username)
	if err != nil {
		return User,err
	}
	return User,nil
}
