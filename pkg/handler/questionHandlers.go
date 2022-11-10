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
	if err != nil {
		h.Loggers.ErrorLogger.Println(err)
	}
	//time to validate
	marhsal, err := json.Marshal(questions)
	w.Write(marhsal)
	return
}
