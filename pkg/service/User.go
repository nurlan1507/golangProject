package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testApp/pkg/helpers"
	"testApp/pkg/models"
	"testApp/pkg/repository"
)

type User struct {
	repo    repository.Authorization
	loggers *helpers.Loggers
	JWT
}

func NewUserService(repo repository.Authorization) *User {
	return &User{repo: repo, JWT: NewJWTManager(), loggers: helpers.InitLoggers()}
}

func (u *User) SignIn(username string, password string) (string, error) {

	return "", nil
}

func (u *User) SignUp(username string, password string) (*models.UserModel, error) {
	u.loggers.InfoLogger.Println(username + "- " + password)
	hashedPassword, err := u.HashPassword(password)
	if err != nil {
		return nil, err
	}
	newUser, err := u.repo.CreateUser(username, string(hashedPassword))
	jwt, err := u.JWT.NewJWT(newUser, 1)
	fmt.Println(jwt)
	if err != nil {
		return nil, err
	}
	token, err := u.JWT.NewRefreshToken(*newUser)
	if err != nil {
		return nil, err
	}
	err = u.repo.UpdateRefreshToken(newUser.Id, token)
	if err != nil {
		return nil, err
	}
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

func (u *User) GetUsers() []models.UserModel {
	arr := u.repo.GetUsers()
	return arr
}
