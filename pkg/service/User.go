package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"testApp/pkg/models"
	"testApp/pkg/repository"
)

type User struct {
	repo repository.Authorization
	JWT  JWT
}

func NewUserService(repo repository.Authorization) *User {
	return &User{repo: repo}
}

func (u *User) SignIn(username string, password string) (string, error) {

	return "", nil
}

func (u *User) SignUp(username string, password string) (*models.UserModel, error) {
	hashedPassword, err := u.HashPassword(password)
	if err != nil {
		return nil, err
	}
	newUser, err := u.repo.CreateUser(username, string(hashedPassword))
	return newUser, nil
}

func (u *User) HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (u *User) CheckPassword(password string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, errors.New("passwords do not match")
	}
	return true, nil
}
