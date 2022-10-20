package main

import (
	"log"
	"testApp"
	"testApp/pkg"
	"testApp/pkg/handler"
)

type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func main() {
	server := &testApp.Server{}
	handler.InitilalizeHandler()
	err := server.RunServer(":4000", pkg.Routes())
	if err != nil {
		return
	}
}
