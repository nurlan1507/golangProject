package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testApp/pkg/helpers"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _ := strconv.Atoi(r.URL.Query().Get("userId"))
	fmt.Println(userId)
	tests, err := h.TestService.GetTests(userId)
	if err != nil {
		if errors.Is(err, helpers.DbError) {
			w.WriteHeader(500)
			h.Loggers.ErrorLogger.Println(err)
			return
		}
		h.Loggers.ErrorLogger.Println(err)
		w.WriteHeader(404)
		w.Write([]byte("Notfound"))
		return
	}
	marshalTests, err := json.Marshal(tests)
	if err != nil {
		h.Loggers.ErrorLogger.Println(err)
		w.WriteHeader(500)
		return
	}
	w.Write(marshalTests)
	return
}
