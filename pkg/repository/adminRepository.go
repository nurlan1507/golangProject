package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"testApp/pkg/helpers"
	"testApp/pkg/models"
)

type AdminRepository struct {
	Db      *pgxpool.Pool
	Loggers helpers.Loggers
}

func (a *AdminRepository) CreateTeacher(email string, username string) (*models.UserModel, error) {
	stmt := `INSERT INTO pending_users(email, username) VALUES ($1,$2) RETURNING *`
	newTeacher := &models.UserModel{}
	err := a.Db.QueryRow(context.Background(), stmt, email, username).Scan(&newTeacher.Id, &newTeacher.Email, &newTeacher.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, helpers.ErrNoRecord
		}
		a.Loggers.ErrorLogger.Println(err)
		return nil, err
	}
	return newTeacher, nil
}

func (a *AdminRepository) CreateTeacherInviteToken(teacherId int, token string) (*models.TeacherInvite, error) {
	newInvitation := &models.TeacherInvite{}
	stmt := `INSERT INTO invited_tokens(teacher_id, token) VALUES ($1, $2) RETURNING *`
	err := a.Db.QueryRow(context.Background(), stmt, teacherId, token).Scan(&newInvitation.InvitationId, &newInvitation.TeacherId, &newInvitation.Token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, helpers.ErrNoRecord
		}
		a.Loggers.ErrorLogger.Println(err)
		return nil, err
	}
	return newInvitation, nil
}

func NewAdminRepository(db *pgxpool.Pool) *AdminRepository {
	return &AdminRepository{Db: db, Loggers: *helpers.InitLoggers()}
}
