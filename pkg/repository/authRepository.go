package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"testApp/pkg/models"
)

type Auth struct {
	Db *pgxpool.Pool
}

func (a *Auth) CreateUser(username string, password string) (*models.UserModel, error) {
	stmt := `INSERT INTO users (username, password) values ($1,$2)`
	result := a.Db.QueryRow(context.Background(), stmt, username, password)
	newUser := &models.UserModel{}
	err := result.Scan(&newUser.Id, &newUser.Username, &newUser.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Fatal("POSHEL NAX")
			return newUser, errors.New("POSHEL NAXUI")
		}
	}
	return newUser, nil
}

func NewAuthRepo(db *pgxpool.Pool) *Auth {
	return &Auth{Db: db}
}
