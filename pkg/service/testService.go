package service

import (
	"testApp/pkg/helpers"
	"testApp/pkg/models"
	"testApp/pkg/repository"
)

type testService struct {
	Loggers    *helpers.Loggers
	Validation *helpers.Validation
	repo       repository.TestRepository
}

func (t *testService) CreateTest() *models.TestModel {

	return nil
}

func NewTestService(repo repository.Repository) *testService {
	return &testService{
		Loggers:    helpers.InitLoggers(),
		Validation: helpers.NewValidation(),
		repo:       repo.TestRepository,
	}
}
