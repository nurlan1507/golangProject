package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"testApp/pkg/helpers"
	"testApp/pkg/models"
	"testApp/pkg/repository"
	"time"
)

var NotAuthorized = errors.New("You are not authorized")
var NoToken = errors.New("No token")
var InvalidToken = errors.New("Token is invalid")
var TokenError = errors.New("Server Error")
var ExpiredToken = errors.New("Token is expired")
var ExpiredRefreshToken = errors.New("RefreshToken is expired")

type ErrorHandlerJwt struct {
	Payload jwt.MapClaims
	Err     error
}

func HandleJWTError(payload jwt.MapClaims, err error) *ErrorHandlerJwt {
	return &ErrorHandlerJwt{
		Err:     err,
		Payload: payload,
	}
}

type Claims struct {
	Username string
	Id       int
	jwt.StandardClaims
}

type Manager struct {
	Loggers *helpers.Loggers
	SignKey string
	repo    repository.Authorization
}

func NewJWTManager(repo repository.Authorization) *Manager {
	manager := &Manager{
		SignKey: "key",
		Loggers: helpers.InitLoggers(),
		repo:    repo,
	}
	return manager
}

func (m *Manager) NewJWT(user *models.UserModel, ttl time.Duration) (string, error) {
	m.Loggers.InfoLogger.Println(m.SignKey)
	claims := &Claims{
		Username: user.Username,
		Id:       user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(m.SignKey))
	if err != nil {
		m.Loggers.ErrorLogger.Println(err)
		return "", err
	}
	return tokenString, nil
}

func (m *Manager) VerifyToken(accessToken string) (jwt.MapClaims, *ErrorHandlerJwt) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, TokenError
		}
		return []byte(m.SignKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return claims, HandleJWTError(claims, ExpiredToken)
		}
		return nil, HandleJWTError(nil, NotAuthorized)
	}
	return claims, nil
}

func (m *Manager) Parse(accessToken string) (string, error) {
	return "", nil
}

func (m *Manager) NewRefreshToken(user models.UserModel) (string, error) {
	claims := &Claims{
		Username: user.Username,
		Id:       user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(m.SignKey))
	if err != nil {
		m.Loggers.ErrorLogger.Println(err)
		return "", err
	}
	return tokenString, nil
}

func (m *Manager) RefreshAccessToken(payload jwt.MapClaims) (string, error) {
	newUserModel := new(models.UserModel)
	newUserModel.Username = fmt.Sprint(payload["Username"])
	userId, ok := payload["Id"].(int)
	if !ok {

	}
	newUserModel.Id = userId
	token, err := m.NewJWT(newUserModel, 1)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (m *Manager) GetRefreshToken(userId int) (*models.RefreshToken, error) {
	token, err := m.repo.GetRefreshToken(userId)
	if err != nil {
		if errors.Is(err, repository.ErrNoRecord) {
			return nil, repository.ErrNoRecord
		}
		return nil, err
	}
	if time.Now().UTC().Hour() > token.Expires.Hour() {
		return nil, ExpiredRefreshToken
	}
	return token, nil
}
