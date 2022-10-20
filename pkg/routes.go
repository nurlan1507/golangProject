package pkg

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testApp/pkg/handler"
)

func Routes() http.Handler {
	handlers := handler.InitilalizeHandler()
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/signUp", handlers.SignUp)
	router.HandlerFunc(http.MethodGet, "/signIn", handlers.SignIn)
	router.HandlerFunc(http.MethodPost, "/signUp", handlers.SignUpPost)
	return router
}
