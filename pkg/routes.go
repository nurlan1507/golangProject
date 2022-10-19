package pkg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
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
	db, err := OpenDb("postgres://postgres:admin@localhost:5432/testApp")
	handlers.DB = db
	handlers.TemplateCache = templateCache
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/signUp", handlers.SignUp)
	router.HandlerFunc(http.MethodGet, "/signIn", handlers.SignIn)

	return router
}

func OpenDb(connectionString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	defer pool.Close()
	return pool, nil
}
