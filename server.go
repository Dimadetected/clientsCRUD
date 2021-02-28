package clientsCRUD

import "net/http"

type Server struct {
	port string
}

func (s *Server) Run() {
	http.ListenAndServe(s.port, nil)
}
