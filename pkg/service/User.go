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

func NewUserService(repo repository.Repository) *User {
	return &User{repo: repo.Authorization, JWT: NewJWTManager(repo), loggers: helpers.InitLoggers()}
}

func (u *User) SignIn(email string, password string) (*models.UserModel, error) {
	user, err := u.repo.GetUser(email, password)
	if err != nil {
		return nil, err
	}
	refreshToken, err := u.JWT.NewRefreshToken(*user)
	err = u.repo.UpdateRefreshToken(user.Id, refreshToken)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	accessToken, err := u.JWT.NewJWT(user, 1)
	if err != nil {
		return nil, err
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return user, nil
}

func (u *User) SignUp(email string, username string, password string, groupId int) (*models.UserModel, error) {
	u.loggers.InfoLogger.Println(username + "- " + password)
	hashedPassword, err := u.HashPassword(password)
	if err != nil {
		return nil, err
	}
	newUser, err := u.repo.CreateUser(email, username, string(hashedPassword), "student", groupId)
	if err != nil {
		return nil, err
	}
	err = u.repo.AddToGroup(newUser.Id, groupId)
	if err != nil {
		u.loggers.ErrorLogger.Println(err)
		return nil, err
	}
	jwt, err := u.JWT.NewJWT(newUser, 1)
	if err != nil {
		return nil, err
	}
	refreshToken, err := u.JWT.NewRefreshToken(*newUser)
	if err != nil {
		return nil, err
	}
	err = u.repo.CreateRefreshToken(newUser.Id, refreshToken)
	if err != nil {
		return nil, err
	}
	newUser.AccessToken = jwt
	newUser.RefreshToken = refreshToken
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

func (u *User) SignUpTeacher(id int, password string) (*models.UserModel, error) {
	hashedPass, err := u.HashPassword(password)
	if err != nil {
		u.loggers.ErrorLogger.Println(err)
		return nil, err
	}
	userData, err := u.repo.DeletePendingUser(id)
	if err != nil {
		return nil, err
	}
	newUser, err := u.repo.CreateUser(userData.Email, userData.Username, string(hashedPass), "Teacher", 0)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
