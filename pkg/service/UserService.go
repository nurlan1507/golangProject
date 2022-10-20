package service

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"testApp/pkg/models"
)

type UserService struct {
	UserModel models.UserModel
	Db        *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) *UserService {
	return &UserService{Db: db}
}

func (u *UserService) SignIn(username string, password string) {

}

func (u *UserService) SignToken() {

}
