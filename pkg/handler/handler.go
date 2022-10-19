package handler

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"html/template"
	"net/http"
)

type itemRep interface {
	saveItem()
}
type Handler struct {
	DB            *pgxpool.Pool
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
