package main

import (
	"fmt"

	"github.com/Nikita213-hub/simpleInMemStorage/handlers"
	"github.com/Nikita213-hub/simpleInMemStorage/internal/storage"
	"github.com/Nikita213-hub/simpleInMemStorage/server"
)

func main() {
	fmt.Println("Basic setup")
	server := server.NewServer("localhost", "8080")
	strg := storage.NewStorage()
	AddHandler := &handlers.BaseHandler{
		Path: "PUT /values/{key}",
		Handler: &handlers.AddHandler{
			Strg: strg,
		},
	}
	GetAllHandler := &handlers.BaseHandler{
		Path: "GET /values",
		Handler: &handlers.GetAllHandler{
			Strg: strg,
		},
	}
	GetHandler := &handlers.BaseHandler{
		Path: "GET /values/{key}",
		Handler: &handlers.GetHandler{
			Strg: strg,
		},
	}
	DeleteHandler := &handlers.BaseHandler{
		Path: "DELETE /values/{key}",
		Handler: &handlers.DeleteHandler{
			Strg: strg,
		},
	}
	err := server.Run(AddHandler, GetHandler, DeleteHandler, GetAllHandler)
	if err != nil {
		panic(err)
	}
}
