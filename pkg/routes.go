package pkg

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"testApp/pkg/handler"
	"testApp/pkg/repository"
)

func Routes() http.Handler {
	db, _ := repository.OpenDb(&repository.Config{DbName: "testApp", User: "postgres", Password: "admin"})
	repos := repository.NewRepository(db)
	handlers, err := handler.InitilalizeHandler(repos)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	//auth
	router.HandlerFunc(http.MethodGet, "/signUp", handlers.SignUp)
	router.HandlerFunc(http.MethodGet, "/signIn", handlers.SignIn)
	router.HandlerFunc(http.MethodPost, "/signUp", handlers.SignUpPost)
	router.HandlerFunc(http.MethodGet, "/getUsers", handlers.GetUsers)

	//homepage
	router.HandlerFunc(http.MethodGet, "/home", handlers.Home)
	return router
}
