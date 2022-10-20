package pkg

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"testApp/pkg/handler"
)

func Routes() http.Handler {
	handlers := new(handler.Handler)
	templateCache, err := NewTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	handlers.TemplateCache = templateCache
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/signUp", handlers.SignUp)
	router.HandlerFunc(http.MethodGet, "/signIn", handlers.SignIn)

	return router
}
