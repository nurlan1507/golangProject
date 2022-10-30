package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"testApp/pkg/helpers"
	"testApp/pkg/models"
	"time"
)

type testRepository struct {
	Db *pgxpool.Pool
}

func (t *testRepository) CreateTest(title string, description string, authorId int, startAt time.Time) (*models.TestModel, error) {
	stmt := `INSERT INTO test (title, description, author_id, create_at, start_at, expires_at) 
    VALUES ($1,$2,$3,$4,$5,$6) RETURNING  id, title,description, author_id,created_at, start_at, expires_at`
	createdAt := time.Now().Format("2006-11-11")
	//startAt = startAt.Format("2006-11-11")
	newTest := &models.TestModel{}
	err := t.Db.QueryRow(context.Background(), stmt, title, description, authorId, createdAt, startAt, startAt).
		Scan(&newTest.Id, &newTest.Title, &newTest.AuthorId, &newTest.CreatedAt, &newTest.StartAt, &newTest.ExpiresAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, helpers.ErrNoRecord
		}
		return nil, err
	}
	return newTest, nil
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
