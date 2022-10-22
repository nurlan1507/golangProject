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

func (a *Auth) CreateUser(email string, username string, password string) (*models.UserModel, error) {
	stmt := `INSERT INTO users(email,username, password, refreshToken) VALUES ($1, $2, $3, '') RETURNING id, email,username,password`
	result := a.Db.QueryRow(context.Background(), stmt, email, username, password)
	newUser := &models.UserModel{}
	err := result.Scan(&newUser.Id, &newUser.Email, &newUser.Username, &newUser.Password)
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

func (a *Auth) UpdateRefreshToken(userId int, refreshToken string) error {
	stmt := `UPDATE users SET refreshToken = $1 WHERE id = $2`
	_, err := a.Db.Exec(context.Background(), stmt, refreshToken, userId)
	if err != nil {
		return err
	}
	return nil
}

func NewAuthRepo(db *pgxpool.Pool) *Auth {
	return &Auth{Db: db}
}
