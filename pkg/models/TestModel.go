package models

import (
	"time"
)

type QuestionModel struct {
	QuestionId   int
	Description  string
	Type         string
	Order        int
	TestId       int
	CorrectValue string //Ответ на вопрос, важен только в случае типа вопроса с инпутом
}

type AnswerModel struct {
	AnswerId   int
	Value      string
	Correct    bool
	QuestionId int
}

type UserAnswers struct {
	UserAnswerId int
	UserId       int
	AnswerId     int
	Value        string //Ответ на вопрос, важен только в случае типа вопроса с инпутом
}

type TestModel struct {
	Id           int
	Title        string
	AuthorId     int
	CreatedAt    time.Time
	ExpiresAt    time.Time
	Questions    []QuestionModel
	Participants []UserModel
}
