package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"testApp/pkg/models"
)

type Authorization interface {
	CreateUser(email string, username string, password string) (*models.UserModel, error)
	GetUsers() []models.UserModel
	GetUser(email string, password string) (*models.UserModel, error)
	UpdateRefreshToken(userId int, refreshToken string) error
	CreateRefreshToken(userId int, refreshToken string) error
	GetRefreshToken(userId int) (*models.RefreshToken, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
	}
}
