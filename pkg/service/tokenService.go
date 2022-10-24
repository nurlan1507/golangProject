package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"testApp/pkg/helpers"
	"testApp/pkg/models"
	"time"
)

var NotAuthorized = errors.New("You are not authorized")
var NoToken = errors.New("No token")
var InvalidToken = errors.New("Token is invalid")
var TokenError = errors.New("Server Error")
var ExpiredToken = errors.New("Token is expired")

type ErrorHandlerJwt struct {
	Payload jwt.MapClaims
	Err     error
}

func (e ErrorHandlerJwt) Error() string {
	panic("implement me")
}

func HandleJWTError(payload jwt.MapClaims, err error) *ErrorHandlerJwt {
	return &ErrorHandlerJwt{
		Err:     err,
		Payload: payload,
	}
}

type JWT interface {
	NewJWT(user *models.UserModel, ttl time.Duration) (string, error)
	VerifyToken(accessToken string) (jwt.MapClaims, *ErrorHandlerJwt)
	Parse(accessToken string) (string, error)
	NewRefreshToken(model models.UserModel) (string, error)
	RefreshAccessToken(claims jwt.MapClaims) (string, error)
}
type Claims struct {
	Username string
	Id       int
	jwt.StandardClaims
}

type Manager struct {
	Loggers *helpers.Loggers
	SignKey string
}

func NewJWTManager() *Manager {
	manager := &Manager{
		SignKey: "key",
		Loggers: helpers.InitLoggers(),
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  strconv.FormatInt(time.Now().Add(30).Unix(), 10),
		"user": user.Id,
	})
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
