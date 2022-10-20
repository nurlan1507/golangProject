package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"testApp/pkg/models"
)

type Auth struct {
	Db *pgxpool.Pool
}

func (a *Auth) CreateUser(username string, password string) (*models.UserModel, error) {
	fmt.Println("LOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOl")
	stmt := `INSERT INTO people(username, password) VALUES ($1, $2) RETURNING *`
	result := a.Db.QueryRow(context.Background(), stmt, username, password)
	newUser := &models.UserModel{}
	err := result.Scan(&newUser.Id, &newUser.Username, &newUser.Password)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return newUser, errors.New("POSHEL NAXUI")
		}
	}
	return newUser, nil
}
func (a *Auth) GetUsers() []models.UserModel {
	stmt := `SELECT * FROM people`
	result, err := a.Db.Query(context.Background(), stmt)
	if err != nil {
		return nil
	}
	var arr []models.UserModel
	fmt.Println(result.RawValues())
	for result.Next() {
		user := &models.UserModel{}
		err = result.Scan(&user.Id, &user.Username, &user.Password)
		arr = append(arr, *user)
	}
	return arr
}

func NewAuthRepo(db *pgxpool.Pool) *Auth {
	return &Auth{Db: db}
}
