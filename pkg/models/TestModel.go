package models

import (
	"time"
)

type TestModel struct {
	Id          int       `json:"Id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SubjectId   int       `json:"subjectId"`
	AuthorId    int       `json:"authorId"`
	GroupId     string    `json:"group"`
	CreatedAt   time.Time `json:"create_at"`
	StartAt     time.Time `json:"st"`
	ExpiresAt   time.Time `json:"expires_at"`
}
type QuestionModel struct {
	QuestionId   int    `json:"questionId"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	Order        int    `json:"order"`
	TestId       int    `json:"testId"`
	CorrectValue string `json:"correctValue"` //Ответ на вопрос, важен только в случае типа вопроса с инпутом
}

type AnswerModel struct {
	AnswerId   int    `json:"answerId"`
	Value      string `json:"value"`
	Correct    bool   `json:"correct"`
	QuestionId int    `json:"questionId"`
}

type UserAnswers struct {
	UserAnswerId int    `json:"userAnswerId"`
	UserId       int    `json:"userId"`
	AnswerId     int    `json:"answerId"`
	Value        string `json:"value"`
	//Ответ на вопрос, важен только в случае типа вопроса с инпутом
}
