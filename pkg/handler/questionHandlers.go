package handler

import (
	"encoding/json"
	"net/http"
	"testApp/pkg/models"
)

func (h *Handler) AddQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		h.Loggers.ErrorLogger.Println(err)
		return
	}
	var questions []*models.QuestionModel
	json.NewDecoder(r.Body).Decode(&questions)
	_, err = h.TestService.AddQuestions(questions)
	//time to validate
	if len(h.TestService.GetValidationErrorMap()) != 0 {
		w.WriteHeader(400)
		marshalValidationErrorMap, err := json.Marshal(h.TestService.GetValidationErrorMap())
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Write(marshalValidationErrorMap)
		return
	}

	marhsal, err := json.Marshal(questions)
	w.Write(marhsal)
	return
}
