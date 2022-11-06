package service

import (
	"bytes"
	"html/template"
	"testApp/pkg/helpers"
	"testApp/pkg/models"
	"testApp/pkg/repository"
)

type testService struct {
	Loggers    *helpers.Loggers
	Validation *helpers.Validation
	repo       repository.TestRepository
}

func (t *testService) CreateTest(newTest *models.TestModel) (*models.TestModel, error) {
	createdTest, err := t.repo.CreateTest(newTest)
	if err != nil {
		t.Loggers.ErrorLogger.Println(err)
		return nil, err
	}
	//рассылка ученикам
	emails, err := t.repo.FindStudents(createdTest.GroupId)
	if err != nil {
		return nil, err
	}
	//sending email to participants
	ts, err := template.ParseFiles("./ui/html/mailTemplates/invitationStudents.html")
	buff := new(bytes.Buffer)
	err = ts.Execute(buff, createdTest)
	if err != nil {
		return nil, helpers.EmailError
	}
	subject := "Invitation\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + buff.String())
	err = sendEmailToParticipants(message, emails)
	if err != nil {
		return nil, err
	}
	return createdTest, nil
}

func (t *testService) AddQuestions([]*models.TestModel) ([]models.TestModel, error) {

	return nil, nil
}
func NewTestService(repo repository.Repository) *testService {
	return &testService{
		Loggers:    helpers.InitLoggers(),
		Validation: helpers.NewValidation(),
		repo:       repo.TestRepository,
	}
}
