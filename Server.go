package testApp

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
)

type config struct {
}
type Server struct {
	httpServer *http.Server
	db         *pgxpool.Pool
}

func (s *Server) RunServer(port string, handler http.Handler) error {
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
