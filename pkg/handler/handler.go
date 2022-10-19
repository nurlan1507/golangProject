package handler

import (
	"html/template"
	"net/http"
)

type Handler struct {
	TemplateCache map[string]*template.Template
}

func (h *Handler) render(w http.ResponseWriter, name string, r *http.Request) {
	ts, ok := h.TemplateCache[name]
	if ok == false {
		http.Error(w, "page does not exist", 500)
		return
	}
	err := ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		return
	}
}
