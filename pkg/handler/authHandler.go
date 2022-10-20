package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testApp/pkg/helpers"
)

func (h *Handler) SignUpPost(w http.ResponseWriter, r *http.Request) {
	//username := r.PostForm.Get("username")
	//password := r.PostForm.Get("password")
	username := "nurlan"
	password := "admin"
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("ASDASDASDASDASDASDASDASDASDASDASDASDASDASDASDASA")
	up, err := h.UserService.SignUp(username, password)
	if err != nil {
		fmt.Println("ASDASDASDASDASDASDASDASDASDASDASDASDASDASDASDASA")
		helpers.BadRequest(w, r, err)
		return
	}
	json.NewEncoder(w).Encode(up)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	up := h.UserService.GetUsers()
	json.NewEncoder(w).Encode(up)
}
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	h.render(w, "signIn.tmpl", r)
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("asdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasda")
	h.render(w, "signIn.tmpl", r)
}
