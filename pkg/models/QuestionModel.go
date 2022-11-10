package models

type QuestionModel struct {
	QuestionId   int                    `json:"questionId"`
	Description  string                 `json:"description"`
	Type         string                 `json:"type"`
	Order        int                    `json:"order"`
	TestId       int                    `json:"testId"`
	Answers      map[string]AnswerModel `json:"answers"`
	Score        float32                `json:"point"`
	CorrectValue string                 `json:"correctValue"` //Ответ на вопрос, важен только в случае типа вопроса с инпутом
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
	Value        string `json:"value"` //Ответ на вопрос, важен только в случае типа вопроса с инпутом
}
