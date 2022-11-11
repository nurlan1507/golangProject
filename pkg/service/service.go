package service

import (
	"github.com/golang-jwt/jwt/v4"
	"testApp/pkg/models"
	"testApp/pkg/repository"
	"time"
)

type UserService interface {
	SignIn(email string, password string) (*models.UserModel, error)
	SignUp(email string, username string, password string) (*models.UserModel, error)
	HashPassword(password string) ([]byte, error)
	CheckPassword(password string, hashedPassword string) (bool, error)
	GetUsers() []models.UserModel
	SignUpTeacher(id int, password string) (*models.UserModel, error)
}

type JWT interface {
	NewJWT(user *models.UserModel, ttl time.Duration) (string, error)
	VerifyToken(accessToken string) (jwt.MapClaims, *ErrorHandlerJwt)
	Parse(accessToken string) (string, error)
	NewRefreshToken(model models.UserModel) (string, error)
	RefreshAccessToken(claims jwt.MapClaims) (string, error)
	GetRefreshToken(userId int) (*models.RefreshToken, error)
	GetClaims(token string) (jwt.MapClaims, error)
}
type TestService interface {
	CreateTest(model *models.TestModel) (*models.TestModel, error)
	AddQuestions(model []*models.QuestionModel, testId int) ([]models.TestModel, error)
	GetValidationErrorMap() map[string]string
}

type Admin interface {
	InviteTeacher(email string, username string) (*models.TeacherInvite, error)
}

type Service struct {
	UserService
	JWT
	AdminService Admin
	TestService
}

func NewService(reps repository.Repository) *Service {
	return &Service{
		UserService:  NewUserService(reps),
		JWT:          NewJWTManager(reps),
		AdminService: NewAdminService(reps),
		TestService:  NewTestService(reps),
	}
}
