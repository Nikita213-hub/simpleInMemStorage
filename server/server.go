package server

import (
	"fmt"
	"net/http"
	"strconv"
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

func (s *Server) Run(handlers ...http.Handler) error {
	fmt.Println("Listening ", s.addr, "...")
	s.mux = http.NewServeMux()
	for i, v := range handlers {
		s.mux.Handle("/test/"+strconv.Itoa(i), v)
	}
	err := http.ListenAndServe(s.addr, s.mux)
	if err != nil {
		return err
	}
	return nil
}
