package helpers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// ServerError when error is related to server
func ServerError(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// BadRequest when error is related to user input
func BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

type Loggers struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	out         io.Closer
}

func InitLoggers() *Loggers {
	file, err := os.OpenFile("serverLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("SAS")
		log.Fatal(err)
	}

	InfoLogger := log.New(file, "INFO: \t", log.Ldate|log.Ltime)
	ErrorLogger := log.New(file, "ERROR: \t", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)

	newHelpers := &Loggers{
		ErrorLogger: ErrorLogger,
		InfoLogger:  InfoLogger,
		out:         file,
	}

	return newHelpers
}

func (l *Loggers) CloseFile(file os.File) error {
	l.out.Close()
	return nil
}
