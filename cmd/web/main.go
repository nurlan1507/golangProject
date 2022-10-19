package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func main() {
	addr := flag.String("address", ":4000", "network adders")
	//postgres := flag.String("db", "postgres://postgres:admin@localhost:5432/postgres", "database addres")

	file, err := os.OpenFile("serverLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	infoLog := log.New(file, "INFOLOG:\t", log.Ldate|log.Ltime)
	errorLog := log.New(file, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	app := &application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routers(),
	}
	err = srv.ListenAndServe()
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}
}
