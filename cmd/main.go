package main

import (
	"log"
	"testApp"
	"testApp/pkg"
)

type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func main() {
	server := &testApp.Server{}
	server.RunServer(":4000", pkg.Routes())
}
