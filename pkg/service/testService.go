package service

import (
	"bytes"
	"fmt"
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
		t.Loggers.ErrorLogger.Println(err)
		return nil, err
	}
	//sending email to participants
	ts, err := template.ParseFiles("./ui/html/mailTemplates/invitationStudents.html")
	buff := new(bytes.Buffer)
	err = ts.Execute(buff, createdTest)
	if err != nil {
		t.Loggers.ErrorLogger.Println(err)
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

func (t *testService) AddQuestions(questions []*models.QuestionModel, testId int) ([]models.TestModel, error) {

	for i, v := range questions {
		fmt.Printf("%+v", v)
		fmt.Println(" ")

		fmt.Println(" ")
		d := t.Validation.CheckQuestions(helpers.NoDescription(questions[i]), fmt.Sprintf("question-description-%v", i), "the description is empty")
		a := t.Validation.CheckQuestions(helpers.QuestionWithoutAnswers(questions[i]), fmt.Sprintf("question-answers-%v", i), "please select a correct answer to this question")
		if a == false || d == false {
			continue
		}
		ind := i + 1
		question, err := t.repo.AddQuestion(v, ind, testId)
		if err != nil {
			t.Loggers.ErrorLogger.Println(err)
			return nil, err
		}
		_, err = t.repo.AddAnswer(question.QuestionId, questions[i].Answers)
		if err != nil {
			t.Loggers.ErrorLogger.Println(err)
			return nil, err
		}
		if err != nil {
			t.Loggers.ErrorLogger.Println(err)
			return nil, err
		}
	}
	return nil, nil
}

func (t *testService) GetTest(testId int) (*models.TestModel, error) {
	test, err := t.repo.GetTest(testId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return test, nil
}

func (t *testService) GetValidationErrorMap() map[string]string {
	return t.Validation.Errors
}
func NewTestService(repo repository.Repository) *testService {
	return &testService{
		Loggers:    helpers.InitLoggers(),
		Validation: helpers.NewValidation(),
		repo:       repo.TestRepository,
	}
}
