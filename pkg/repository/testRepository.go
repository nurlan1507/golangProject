package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"testApp/pkg/helpers"
	"testApp/pkg/models"
)

type testRepository struct {
	Db *pgxpool.Pool
}

func (t *testRepository) FindStudents(groupId string) ([]string, error) {
	stmt := `SELECT email FROM users WHERE group_name LIKE $1`
	query, err := t.Db.Query(context.Background(), stmt, groupId)
	if err != nil {
		if errors.Is(err, helpers.ErrNoRecord) {
			return nil, helpers.ErrNoRecord
		}
		return nil, err
	}
	var emails []string
	for query.Next() {
		var email string
		err = query.Scan(&email)
		if err != nil {
			return nil, err
		}
		emails = append(emails, email)
	}
	return emails, nil
}

func (t *testRepository) CreateTest(newTest *models.TestModel) (*models.TestModel, error) {
	test := &models.TestModel{}
	stmt := `INSERT INTO test(title,description,author_id,created_at,start_at,expires_at,group_name) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING *`
	err := t.Db.QueryRow(context.Background(), stmt, newTest.Title, newTest.Description, newTest.AuthorId, newTest.CreatedAt, newTest.StartAt, newTest.StartAt, newTest.GroupId).
		Scan(&test.Id, &test.Title, &test.Description, &test.AuthorId, &test.CreatedAt, &test.StartAt, &test.ExpiresAt, &test.GroupId)
	if err != nil {
		if errors.Is(err, helpers.ErrNoRecord) {
			return nil, helpers.ErrNoRecord
		}
		return nil, err
	}
	return test, nil
}

func (t *testRepository) AddQuestion(description string, questionType string, questionOrder int, testId int) (*models.QuestionModel, error) {
	stmt := `INSERT INTO question(description,question_type, question_order, test_id) VALUES($1,$2,$3,$4) 
	RETURNING question_id,description,question_type,question_order`
	newQuestion := &models.QuestionModel{}
	err := t.Db.QueryRow(context.Background(), stmt, description, questionType, questionOrder, testId).
		Scan(&newQuestion.QuestionId, &newQuestion.Description, &newQuestion.Type, &newQuestion.Order)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, helpers.ErrNoRecord
		}
		return nil, err
	}
	return newQuestion, nil
}

//func (t *testRepository) AddAnswer(value string, correct bool, questionId int) (*models.AnswerModel,error){
//	stmt := `INSERT INTO answer(value,correct,question_id) VALUES`
//}
//
//func (t *testRepository) DeleteQuestion() {
//
//}

func NewTestRepository(db *pgxpool.Pool) *testRepository {
	return &testRepository{
		Db: db,
	}
}
