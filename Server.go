package testApp

import (
	"net/http"
)

type config struct {
}
type Server struct {
	httpServer *http.Server
}

func (s *Server) RunServer(port string, handler http.Handler) error {
	//file, err := os.OpenFile("serverLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//if err != nil {
	//	return err
	//}
	//infoLog := log.New(file, "INFO:\t", log.Ldate|log.Ltime)
	//errorLog := log.New(file, "ERROR: \t", log.Ldate|log.Ltime|log.Lshortfile)
	//s.ErrorLog = errorLog
	//s.InfoLog = infoLog
	s.httpServer = &http.Server{
		Addr:    port,
		Handler: handler,
	}
	err := s.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
