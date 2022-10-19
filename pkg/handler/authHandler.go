package handler

import (
	"html/template"
	"net/http"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ASdasd"))
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/pages/signIn.tmpl",
	}

	tmpl, _ := template.ParseFiles(files...)
	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
