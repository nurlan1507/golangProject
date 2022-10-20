package service

import (
	"testApp/pkg/models"
	"testApp/pkg/repository"
)

type UserService interface {
	SignIn(username string, password string) (string, error)
	SignUp(username string, password string) (*models.UserModel, error)
	HashPassword(password string) ([]byte, error)
	CheckPassword(password string, hashedPassword string) (bool, error)
}

type Service struct {
	UserService
}

func NewService(reps *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(reps.Authorization),
	}
}
