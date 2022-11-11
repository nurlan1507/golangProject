package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testApp/pkg/models"
)

type AddQuestionForm struct {
	Questions []*models.QuestionModel `json:"questions"`
}

func (h *Handler) AddQuestions(w http.ResponseWriter, r *http.Request) {
	testId, _ := strconv.Atoi(r.URL.Query().Get("testId"))
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
		h.Loggers.ErrorLogger.Println(err)
		return
	}
	_, err = h.TestService.AddQuestions(Form.Questions, testId)
	//time to validate
	if err != nil {

		return
	}

	var errors = h.TestService.GetValidationErrorMap()
	if len(errors) != 0 {
		marshalValidationErrorMap, err := json.Marshal(&errors)
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
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			h.Loggers.ErrorLogger.Println(err)
			return
		}
	}
	w.Write(marhsal)
	return
}

func (h *Handler) GetMyTests(w http.ResponseWriter, r *http.Request) {
	//id, _ := strconv.Atoi(r.URL.Query().Get("userId"))

}
