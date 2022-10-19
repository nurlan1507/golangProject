package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routers() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})
	return router
}
