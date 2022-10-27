package handler

import (
	"fmt"
	"log"
	"net/http"
	"testApp/pkg/helpers"
)

func (h *Handler) AddTeacher(w http.ResponseWriter, r *http.Request) {

	h.render(w, "adminPanel.tmpl", nil, 200)
}

func (h *Handler) AddTeacherPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("pizda")
		return
	}
	form := &AuthForm{
		Email:     r.PostForm.Get("email"),
		Username:  r.PostForm.Get("username"),
		Validator: &helpers.Validation{Errors: map[string]string{}},
	}
	_, err = h.AdminService.InviteTeacher(form.Email, form.Username)
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Redirect(w, r, "/home", 303)
}
