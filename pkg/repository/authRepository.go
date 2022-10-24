package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"testApp/pkg/models"
	"time"
)

type Auth struct {
	Db *pgxpool.Pool
}

func (a *Auth) CreateUser(email string, username string, password string) (*models.UserModel, error) {
	stmt := `INSERT INTO users(email,username, password) VALUES ($1, $2, $3) RETURNING id, email,username,password`
	result := a.Db.QueryRow(context.Background(), stmt, email, username, password)
	fmt.Println(result)
	newUser := &models.UserModel{}
	err := result.Scan(&newUser.Id, &newUser.Email, &newUser.Username, &newUser.Password)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return newUser, ErrNoRecord
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

func (a *Auth) GetUser(email string) (*models.UserModel, error) {
	stmt := `SELECT email,username,password FROM users WHERE email LIKE $1`
	result, err := a.Db.Query(context.Background(), stmt, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}
	user := &models.UserModel{}
	err = result.Scan(&user.Id, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *Auth) UpdateRefreshToken(userId int, refreshToken string) error {
	stmt := `UPDATE refreshTokens SET refresh_token=$1 WHERE user_id= $2`
	res, err := a.Db.Exec(context.Background(), stmt, refreshToken, userId)
	res.Update()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNoRecord
		}
		return err
	}
	return nil
}

func (a *Auth) CreateRefreshToken(userId int, refreshToken string) error {
	stmt := `INSERT INTO refreshTokens (user_id, refresh_token, expires) VALUES ($1,$2, $3) RETURNING *`
	expiresDate := time.Now().AddDate(0, 0, 30)
	_, err := a.Db.Query(context.Background(), stmt, userId, refreshToken, expiresDate)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (a *Auth) GetRefreshToken(userId int) (*models.RefreshToken, error) {
	stmt := `SELECT refresh_token,expires FROM refreshTokens WHERE user_id=$1`
	res, err := a.Db.Query(context.Background(), stmt, userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}
	token := &models.RefreshToken{}
	err = res.Scan(&token.Token, &token.Expires)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func NewAuthRepo(db *pgxpool.Pool) *Auth {
	return &Auth{Db: db}
}
