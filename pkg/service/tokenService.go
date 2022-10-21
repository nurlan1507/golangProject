package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
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
	signKey string
}

func NewJWTManager() *Manager {
	manager := new(Manager)
	manager.signKey = "signkey"
	return manager
}

func (m *Manager) NewJWT(user *models.UserModel, ttl time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(ttl).Unix()
	claims["username"] = user.Username
	claims["id"] = user.Id
	tokenString, err := token.SignedString(m.signKey)
	if err != nil {
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
		return []byte(m.signKey), nil
	}
	token, err := jwt.ParseWithClaims(accessToken, claims, keyFunc)
	if err != nil {
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
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(30).Unix()
	claims["id"] = user.Id
	tokenString, err := token.SignedString(m.signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
