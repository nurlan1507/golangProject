package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"testApp/pkg/models"
)

type Authorization interface {
	CreateUser(email string, username string, password string) (*models.UserModel, error)
	GetUsers() []models.UserModel
	UpdateRefreshToken(userId int, refreshToken string) error
}

type Repository struct {
	Authorization
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
	}
}
