package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testApp/pkg/models"
)

type q struct {
	Questions int `json:"questions"`
}

func (h *Handler) CreateTest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.Loggers.ErrorLogger.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	newTest := &models.TestModel{}
	err = json.NewDecoder(r.Body).Decode(&newTest)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}
	fmt.Printf("%+v", newTest)
	result, err := h.TestService.CreateTest(newTest)
	if err != nil {
		w.WriteHeader(400)
		h.Loggers.ErrorLogger.Println(err)
	}

	mars, err := json.Marshal(result)
	if err != nil {
		return
	}
	_, err = w.Write(mars)
	if err != nil {
		w.WriteHeader(500)
		return
	}
}

func (h *Handler) CreateTestPost(w http.ResponseWriter, r *http.Request) {

}
