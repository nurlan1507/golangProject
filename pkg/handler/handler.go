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
	TokenService  service.JWT
	AdminService  service.Admin
	TestService   service.TestService
	TemplateCache map[string]*template.Template
	Loggers       *helpers.Loggers
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func InitilalizeHandler(repos *repository.Repository) (*Handler, error) {
	templateCache, err := NewTemplateCache()
	if err != nil {
		return nil, err
	}
	var services = service.NewService(*repos)
	var adminServices = service.NewAdminService(*repos)
	return &Handler{
		UserService:   services.UserService,
		TemplateCache: templateCache,
		Loggers:       helpers.InitLoggers(),
		TokenService:  services.JWT,
		AdminService:  adminServices,
		TestService:   services.TestService,
	}, nil
}

func (h *Handler) render(w http.ResponseWriter, name string, data *templateData.TemplateData, code int) {
	ts, ok := h.TemplateCache[name]
	if ok == false {
		http.Error(w, "page does not exist", 500)
		return
	}
	//buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), 500)
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
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		cache[name] = ts
	}
	return cache, nil
}
