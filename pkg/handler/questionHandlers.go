package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testApp/pkg/models"
)

type AddQuestionForm struct {
	TestId    string                  `json:"testId"`
	Questions []*models.QuestionModel `json:"questions"`
}

func (h *Handler) AddQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Form := &AddQuestionForm{}
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(500)
		h.Loggers.ErrorLogger.Println(err)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&Form)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(err)
		h.Loggers.ErrorLogger.Println(err)
		return
	}
	id, _ := strconv.Atoi(Form.TestId)
	_, err = h.TestService.AddQuestions(Form.Questions, id)
	//time to validate
	if err != nil {

		return
	}

	if len(h.TestService.GetValidationErrorMap()) != 0 {
		marshalValidationErrorMap, err := json.Marshal(h.TestService.GetValidationErrorMap())
		if err != nil {
			h.Loggers.ErrorLogger.Println(err)
			w.WriteHeader(500)
			w.Write(marshalValidationErrorMap)
			return
		}
		h.Loggers.ErrorLogger.Println(err)
		w.WriteHeader(400)
		w.Write(marshalValidationErrorMap)
		return
	}
	marhsal, err := json.Marshal(&Form)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(400)
	w.Write(marhsal)
	return
}
