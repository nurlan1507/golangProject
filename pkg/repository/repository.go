package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"testApp/pkg/models"
	"time"
)

type Authorization interface {
	CreateUser(email string, username string, password string, role string) (*models.UserModel, error)
	GetUsers() []models.UserModel
	GetUser(email string, password string) (*models.UserModel, error)
	UpdateRefreshToken(userId int, refreshToken string) error
	CreateRefreshToken(userId int, refreshToken string) error
	GetRefreshToken(userId int) (*models.RefreshToken, error)
	DeletePendingUser(userId int) (*models.UserModel, error)
}
type IAdminRepository interface {
	CreateTeacher(email string, username string) (*models.UserModel, error)
	CreateTeacherInviteToken(teacherId int, token string) (*models.TeacherInvite, error)
}

type TestRepository interface {
	CreateTest(title string, description string, authorId int, startAt time.Time) (*models.TestModel, error)
	AddQuestion(description string, questionType string, questionOrder int, testId int) (*models.QuestionModel, error)
}

type Repository struct {
	Authorization
	AdminRepository IAdminRepository
	TestRepository
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization:   NewAuthRepo(db),
		AdminRepository: NewAdminRepository(db),
		TestRepository:  NewTestRepository(db),
	}
}
