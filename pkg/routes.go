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
	handlers := handler.InitilalizeHandler(repos)
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))
	router.HandlerFunc(http.MethodGet, "/signUp", handlers.SignUp)
	router.HandlerFunc(http.MethodGet, "/signIn", handlers.SignIn)
	router.HandlerFunc(http.MethodPost, "/signUp", handlers.SignUpPost)
	router.HandlerFunc(http.MethodGet, "/getUsers", handlers.GetUsers)
	return router
}
