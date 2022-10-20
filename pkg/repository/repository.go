package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"testApp/pkg/models"
)

type Authorization interface {
	CreateUser(username string, password string) (*models.UserModel, error)
	GetUsers() []models.UserModel
}

type Repository struct {
	Authorization
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
	}
}
