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
	var a []*models.QuestionModel
	json.NewDecoder(r.Body).Decode(&a)
	marhsal, err := json.Marshal(a)
	w.Write(marhsal)
	return
}
