package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"testApp/pkg/helpers"
	"testApp/pkg/models"
	"testApp/pkg/repository"
	"time"
)

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
	Role     string
	Email    string
	jwt.StandardClaims
}

type Manager struct {
	Loggers   *helpers.Loggers
	SignKey   string
	AuthRepo  repository.Authorization
	AdminRepo repository.IAdminRepository
}

func NewJWTManager(repo repository.Repository) *Manager {
	manager := &Manager{
		SignKey:   "key",
		Loggers:   helpers.InitLoggers(),
		AuthRepo:  repo.Authorization,
		AdminRepo: repo.AdminRepository,
	}
	return manager
}

func (m *Manager) NewJWT(user *models.UserModel, ttl time.Duration) (string, error) {
	m.Loggers.InfoLogger.Println(m.SignKey)
	claims := &Claims{
		Username: user.Username,
		Id:       user.Id,
		Role:     user.Role,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(m.SignKey))
	if err != nil {
		fmt.Println(err)
		return "", helpers.TokenError
	}
	return tokenString, nil
}

func (m *Manager) VerifyToken(accessToken string) (jwt.MapClaims, *ErrorHandlerJwt) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, helpers.TokenError
		}
		return []byte(m.SignKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return claims, HandleJWTError(claims, helpers.ExpiredToken)
		}
		return nil, HandleJWTError(nil, helpers.NotAuthorized)
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
		Role:     user.Role,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(m.SignKey))
	if err != nil {
		fmt.Println(err)
		m.Loggers.ErrorLogger.Println(err)
		return "", helpers.TokenError
	}
	return tokenString, nil
}

func (m *Manager) RefreshAccessToken(payload jwt.MapClaims) (string, error) {
	newUserModel := new(models.UserModel)
	newUserModel.Username = fmt.Sprint(payload["Username"])
	newUserModel.Role = fmt.Sprint(payload["Role"])
	newUserModel.Email = fmt.Sprint(payload["Email"])
	id := fmt.Sprint(payload["Id"])
	fmt.Println(id)

	newUserModel.Id, _ = strconv.Atoi(id)
	token, err := m.NewJWT(newUserModel, 1)
	if err != nil {
		return "", helpers.TokenError
	}
	return token, nil
}

func (m *Manager) GetRefreshToken(userId int) (*models.RefreshToken, error) {
	token, err := m.AuthRepo.GetRefreshToken(userId)
	if err != nil {
		if errors.Is(err, helpers.ErrNoRecord) {
			return nil, helpers.ErrNoRecord
		}
		return nil, helpers.TokenError
	}
	if time.Now().UTC().Hour() > token.Expires.Hour() {
		return nil, helpers.ExpiredRefreshToken
	}
	return token, nil
}

func (m *Manager) GetClaims(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, helpers.TokenError
		}
		return []byte(m.SignKey), nil
	})
	return claims, nil
}
