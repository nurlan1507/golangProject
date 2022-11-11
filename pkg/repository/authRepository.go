package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"testApp/pkg/helpers"
	"testApp/pkg/models"
	"time"
)

type Auth struct {
	Db      *pgxpool.Pool
	loggers helpers.Loggers
}

func (a *Auth) CreateUser(email string, username string, password string, role string) (*models.UserModel, error) {
	stmt := `INSERT INTO users(email,username, password,role ) VALUES ($1, $2, $3,$4) RETURNING id, email,username,password,'student'`
	newUser := &models.UserModel{}
	err := a.Db.QueryRow(context.Background(), stmt, email, username, password, role).Scan(&newUser.Id, &newUser.Email, &newUser.Username, &newUser.Password, &newUser.Role)
	var pgErr *pgconn.PgError
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, helpers.ErrNoRecord
		}
		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "23505" {
				fmt.Println(err)
				return nil, helpers.ErrDuplicate
			}
		}
		a.loggers.ErrorLogger.Println(err)
		return nil, err
	}
	return newUser, nil
}
func (a *Auth) GetUsers() []models.UserModel {
	stmt := `SELECT * FROM people`
	result, err := a.Db.Query(context.Background(), stmt)
	if err != nil {
		a.loggers.ErrorLogger.Println(err)
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

func (a *Auth) GetUser(email string, password string) (*models.UserModel, error) {
	stmt := `SELECT id,email,username,password,role FROM users WHERE email = $1`
	result := a.Db.QueryRow(context.Background(), stmt, email)

	user := &models.UserModel{}
	err := result.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Role)
	fmt.Printf("%+v", *user)
	fmt.Println(user.Password)
	if err != nil {
		a.loggers.ErrorLogger.Println(err)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, helpers.NoSuchUser
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, helpers.PasswordIncorrect
	}
	return user, nil
}

func (a *Auth) UpdateRefreshToken(userId int, refreshToken string) error {
	stmt := `UPDATE refreshTokens SET refresh_token=$1 WHERE user_id= $2`
	res, err := a.Db.Exec(context.Background(), stmt, refreshToken, userId)
	res.Update()
	if err != nil {
		a.loggers.ErrorLogger.Println(err)
		if errors.Is(err, pgx.ErrNoRows) {
			return helpers.ErrNoRecord
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
			return nil, helpers.ErrNoRecord
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

func (a *Auth) DeletePendingUser(userId int) (*models.UserModel, error) {
	stmt := `DELETE FROM pending_users WHERE teacher_id=$1 RETURNING email,username`
	newUser := &models.UserModel{}
	err := a.Db.QueryRow(context.Background(), stmt, userId).Scan(&newUser.Email, &newUser.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, helpers.ErrNoRecord
		}
		a.loggers.ErrorLogger.Println(err)
		return nil, err
	}
	return newUser, nil
}
func (a *Auth) AddToGroup(userId int, groupId int) error {
	stmt := `INSERT INTO groups_students (student_id,group_id) 	VALUES ($1,$2)`
	res, err := a.Db.Exec(context.Background(), stmt, userId, groupId)
	if err != nil {
		return err
	}
	res.Insert()
	return nil
}
func NewAuthRepo(db *pgxpool.Pool) *Auth {
	return &Auth{Db: db, loggers: *helpers.InitLoggers()}
}
