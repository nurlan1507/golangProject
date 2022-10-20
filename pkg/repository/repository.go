package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"testApp/pkg/service"
)

type Authorization interface {
}

type Repository struct {
	Authorization
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: service.NewUserService(db),
	}
}
