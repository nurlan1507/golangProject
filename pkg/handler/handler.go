package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
	"testApp/pkg/helpers"
	"testApp/pkg/repository"
	"testApp/pkg/service"
	"testApp/pkg/templateData"
)

type Handler struct {
	UserService   service.UserService
	TemplateCache map[string]*template.Template
	Loggers       *helpers.Loggers
}

func InitilalizeHandler(repos *repository.Repository) (*Handler, error) {
	templateCache, err := NewTemplateCache()
	if err != nil {
		return nil, err
	}
	return &Handler{UserService: service.NewService(repos).UserService, TemplateCache: templateCache, Loggers: helpers.InitLoggers()}, nil
}

func (h *Handler) render(w http.ResponseWriter, name string, data *templateData.TemplateData) {
	ts, ok := h.TemplateCache[name]
	if ok == false {
		http.Error(w, "page does not exist", 500)
		return
	}
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		return
	}
}

// NewTemplateCache to generate new template cache
func NewTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{
			"./ui/html/base.tmpl",
			page,
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
