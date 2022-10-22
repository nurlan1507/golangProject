package service

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"testApp/pkg/helpers"
	"testApp/pkg/models"
	"time"
)

type JWT interface {
	NewJWT(user *models.UserModel, ttl time.Duration) (string, error)
	VerifyToken(accessToken string) (jwt.MapClaims, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken(model models.UserModel) (string, error)
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(ttl).Unix(),
		"username": user.Username,
		"id":       user.Id,
	})
	tokenString, err := token.SignedString([]byte(m.SignKey))
	if err != nil {
		m.Loggers.ErrorLogger.Println(err)
		return "", err
	}
	return tokenString, nil
}

func (m *Manager) VerifyToken(accessToken string) (jwt.MapClaims, error) {
	var claims = jwt.MapClaims{}
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return m.SignKey, nil
	}
	token, err := jwt.ParseWithClaims(accessToken, claims, keyFunc)
	if err != nil {
		m.Loggers.ErrorLogger.Println(err)
		return nil, err
	}
	if token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("Token is invalid")
	}
	return claims, nil
}

func (m *Manager) Parse(accessToken string) (string, error) {
	return "", nil
}

func (m *Manager) NewRefreshToken(user models.UserModel) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": json.Number(strconv.FormatInt(time.Now().Add(30).Unix(), 10)),
	})
	tokenString, err := token.SignedString([]byte(m.SignKey))
	if err != nil {
		m.Loggers.ErrorLogger.Println(err)
		return "", err
	}
	return tokenString, nil
}
