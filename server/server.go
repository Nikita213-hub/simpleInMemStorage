package server

import (
	"fmt"
	"net/http"

	"github.com/Nikita213-hub/simpleInMemStorage/handlers"
)

type Server struct {
	addr string
	mux  *http.ServeMux
}

func NewServer(host, port string) *Server {
	addr := host + ":" + port
	return &Server{
		addr: addr,
		mux:  nil,
	}
}

func (s *Server) Run(handlers ...*handlers.BaseHandler) error {
	fmt.Println("Listening ", s.addr, "...")
	s.mux = http.NewServeMux()
	for _, v := range handlers {
		s.mux.Handle(v.Path, v.Handler)
	}
	err := http.ListenAndServe(s.addr, s.mux)
	if err != nil {
		return err
	}
	return nil
}
