package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"testApp/pkg/helpers"
	"testApp/pkg/models"
)

type testRepository struct {
	Db *pgxpool.Pool
}

func (t *testRepository) FindStudents(groupId int) ([]string, error) {
	stmt := `SELECT email FROM users WHERE group_id LIKE $1`
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
	fmt.Println(newTest)
	test := &models.TestModel{}
	stmt := `INSERT INTO test(title,description,author_id,created_at,start_at,expires_at,group_id) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING * `
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

func (t *testRepository) AddQuestion(question *models.QuestionModel, order int, testId int) (*models.QuestionModel, error) {
	stmt := `INSERT INTO question(description,question_type, question_order, correct_answer, test_id, point) VALUES($1,$2,$3,$4,$5,$6) 
	RETURNING question_id,description,question_type, question_order, correct_answer`
	newQuestion := &models.QuestionModel{}
	err := t.Db.QueryRow(context.Background(), stmt, question.Description, question.Type, order, question.CorrectValue, testId, question.Point).
		Scan(&newQuestion.QuestionId, &newQuestion.Description, &newQuestion.Type, &newQuestion.Order, &newQuestion.CorrectValue)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, helpers.ErrNoRecord
		}
		//if pgError, ok := err.(pgx.Err); ok {
		// handle error for 1< score <3 constraint
		//}
		return nil, err
	}
	return newQuestion, nil
}

func (t *testRepository) AddAnswer(questionId int, answers map[string]*models.AnswerModel) ([]models.AnswerModel, error) {
	var addedAnswers = make([]models.AnswerModel, 0, 4)
	for key, _ := range answers {
		stmt := `INSERT INTO answer(value,correct,question_id) values($1,$2,$3) RETURNING answer_id, value,correct,question_id `
		answer := &models.AnswerModel{}
		err := t.Db.QueryRow(context.Background(), stmt, answers[key].Value, answers[key].Correct, questionId).
			Scan(&answer.AnswerId, &answer.Value, &answer.Correct, &answer.QuestionId)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil, helpers.ErrNoRecord
			}
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				return nil, pgErr
			}
			return nil, err
		}
		addedAnswers = append(addedAnswers, *answer)
	}
	fmt.Println(answers)
	return addedAnswers, nil
}

//	func (t *testRepository) AddAnswer(value string, correct bool, questionId int) (*models.AnswerModel,error){
//		stmt := `INSERT INTO answer(value,correct,question_id) VALUES`
//	}
//
// func (t *testRepository) DeleteQuestion() {
//
// }
func (t *testRepository) GetTest(testId int) (*models.TestModel, error) {
	stmt := `SELECT * FROM test t WHERE t.id like $1`
	newTest := &models.TestModel{}
	err := t.Db.QueryRow(context.Background(), stmt, testId).Scan(&newTest.Id, &newTest.Title, &newTest.Description, &newTest.AuthorId, &newTest.CreatedAt, &newTest.StartAt, &newTest.ExpiresAt, newTest.GroupId)
	var pgErr *pgconn.PgError
	if err != nil {
		if errors.As(err, &pgErr) {
			return nil, helpers.ErrNoRecord
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, helpers.ErrNoRecord
		}
		return nil, err
	}
	return nil, nil
}

func NewTestRepository(db *pgxpool.Pool) *testRepository {
	return &testRepository{
		Db: db,
	}
}
