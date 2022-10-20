package pkg

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testApp/pkg/handler"
	"testApp/pkg/repository"
)

func Routes() http.Handler {
	db, _ := repository.OpenDb(&repository.Config{DbName: "testApp", User: "postgres", Password: "admin"})
	repos := repository.NewRepository(db)
	handler.InitilalizeHandler(repos)
	handlers := handler.InitilalizeHandler(repos)
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/signUp", handlers.SignUp)
	router.HandlerFunc(http.MethodGet, "/signIn", handlers.SignIn)
	router.HandlerFunc(http.MethodPost, "/signUp", handlers.SignUpPost)
	router.HandlerFunc(http.MethodGet, "/getUsers", handlers.GetUsers)
	return router
}
