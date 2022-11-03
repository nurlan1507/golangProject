package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"testApp/pkg/helpers"
	"testApp/pkg/service"
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
	teacher, err := h.AdminService.InviteTeacher(form.Email, form.Username)
	if err != nil {
		fmt.Println(err)
		return
	}
	//html template of email
	ts, err := template.ParseFiles("./ui/html/mailTemplates/invite.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	buff := new(bytes.Buffer)
	type emailData struct {
		UserName string
		Token    string
	}
	err = ts.Execute(buff, &emailData{UserName: form.Username, Token: teacher.Token})
	if err != nil {
		fmt.Println(err)
		return
	}
	subject := "Invitation\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + buff.String())
	err = service.SendMessage(message, []string{form.Email})
	if err != nil {
		return
	}
	http.Redirect(w, r, "/home", 303)
}
