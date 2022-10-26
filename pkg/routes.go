package pkg

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"reflect"
	"testApp/pkg/handler"
	"testApp/pkg/repository"
)

func Routes() http.Handler {
	var db, _ = repository.OpenDb(&repository.Config{DbName: "testApp", User: "postgres", Password: "admin"})
	var Repos = repository.NewRepository(db)
	handlers, err := handler.InitilalizeHandler(Repos)
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
	router.HandlerFunc(http.MethodPost, "/signIn", handlers.SignInPost)
	router.HandlerFunc(http.MethodPost, "/signUp", handlers.SignUpPost)
	router.HandlerFunc(http.MethodGet, "/getUsers", handlers.GetUsers)

	//homepage
	fmt.Println(reflect.TypeOf(handlers.Home))
	router.Handle(http.MethodGet, "/home", handlers.AuthMiddleware(handlers.Home))
	return router
}

func lol(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("lol"))
}
