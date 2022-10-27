package service

import (
	"testApp/pkg/models"
	"testApp/pkg/repository"
)

type AdminService struct {
	repo repository.IAdminRepository
	JWT
}

func (A *AdminService) InviteTeacher(email string, username string) (*models.TeacherInvite, error) {
	newTeacher, err := A.repo.CreateTeacher(email, username)
	if err != nil {
		return nil, err
	}
	newToken, err := A.JWT.NewJWT(newTeacher, 30)
	//fmt.Println(newToken)
	//fmt.Println(newTeacher)
	if err != nil {
		return nil, err
	}
	invitation, err := A.repo.CreateTeacherInviteToken(newTeacher.Id, newToken)
	if err != nil {
		return nil, err
	}
	return invitation, nil
}

func NewAdminService(repo repository.Repository) *AdminService {
	return &AdminService{repo: repo.AdminRepository, JWT: NewJWTManager(repo)}
}
