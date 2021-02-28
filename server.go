package clientsCRUD

import (
	"net/http"
	"time"
)

type Server struct {
	http *http.Server
}

func (s *Server) Run(port string) {
	s.http = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
}
