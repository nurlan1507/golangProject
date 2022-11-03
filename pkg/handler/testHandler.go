package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type q struct {
	Questions int `json:"questions"`
}

func (h *Handler) CreateTest(w http.ResponseWriter, r *http.Request) {

	h.render(w, "createTest.tmpl", nil, 200)
}

func (h *Handler) CreateTestPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	var k q
	json.Unmarshal(buf, &k)
	fmt.Println(k.Questions)
	//fmt.Printf("Body : %s\n ", buf)
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(k)
}
