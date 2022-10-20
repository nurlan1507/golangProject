package main

import (
	"log"
	"testApp"
	"testApp/pkg"
	"testApp/pkg/repository"
)

type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func main() {
	server := &testApp.Server{}
	db, err := repository.OpenDb(&repository.Config{DbName: "testApp", User: "postgres", Password: "admin"})
	if err != nil {
		return
	}
	repository.NewRepository(db)
	err = server.RunServer(":4000", pkg.Routes())
	if err != nil {
		return
	}
}
