package main

import (
	"fmt"
	"net/http"

	"github.com/Nikita213-hub/simpleInMemStorage/server"
)

type dummyHandler struct {
	responseMessage string
}

func (dh dummyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(dh.responseMessage))
}

func main() {
	fmt.Println("Basic setup")
	dummyHandler := dummyHandler{responseMessage: "Hello, basic setup\n"}
	server := server.NewServer("localhost", "8080")
	err := server.Run(dummyHandler)
	if err != nil {
		panic(err)
	}
}
