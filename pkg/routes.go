package pkg

import (
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"testApp/pkg/handler"
	"testApp/pkg/repository"
)

func Routes() http.Handler {
	err := godotenv.Load(".env")

	var db, _ = repository.OpenDb(&repository.Config{DbName: os.Getenv("DbName"), User: os.Getenv("User"), Password: os.Getenv("Password")})
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
	//router.HandlerFunc(http.MethodGet, "/signUp", handlers.SignUp)
	router.HandlerFunc(http.MethodGet, "/signIn", handlers.SignIn)
	router.HandlerFunc(http.MethodPost, "/signIn", handlers.SignInPost)
	router.HandlerFunc(http.MethodPost, "/signUp", handlers.EnableCors(handlers.SignUpPost))
	router.HandlerFunc(http.MethodGet, "/getUsers", handlers.GetUsers)
	router.HandlerFunc(http.MethodGet, "/signUpTeacher", handlers.SignUpTeacher)
	router.HandlerFunc(http.MethodPost, "/signUpTeacher", handlers.SignUpTeacherPost)
	//homepage
	router.HandlerFunc(http.MethodGet, "/", handlers.AuthMiddleware(handlers.Home))

	//test
	router.HandlerFunc(http.MethodGet, "/createTest", handlers.CreateTest)
	router.HandlerFunc(http.MethodPost, "/createTest", handlers.CreateTestPost)
	//admin route
	router.HandlerFunc(http.MethodGet, "/addTeacher", handlers.AuthMiddleware(handlers.IsAdmin(handlers.AddTeacher)))
	router.HandlerFunc(http.MethodPost, "/addTeacher", handlers.AuthMiddleware(handlers.IsAdmin(handlers.AddTeacherPost)))

	router.HandlerFunc(http.MethodGet, "/q", func(writer http.ResponseWriter, request *http.Request) {
		//json.NewEncoder(writer).Encode(5)
		writer.Write([]byte("2"))
	})
	return router
}

func lol(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("lol"))
}
