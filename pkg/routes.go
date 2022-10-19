package pkg

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testApp/pkg/handler"
)

func Routes() http.Handler {
	handlers := new(handler.Handler)
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/signUp", handlers.SignUp)
	router.HandlerFunc(http.MethodGet, "/signIn", handlers.SignIn)

	return router
}
