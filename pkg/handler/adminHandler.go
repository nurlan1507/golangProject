package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) AddTeacher(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/baseAdmin.tmpl",
		"./ui/html/adminPages/adminPanel.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	err = ts.ExecuteTemplate(w, "baseAdmin.tmpl", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

}
