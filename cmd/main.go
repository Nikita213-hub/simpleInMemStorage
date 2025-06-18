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
	AddHandler := &handlers.AddHandler{
		Path: "/values/{key}",
		Strg: strg,
	}
	GetHandler := &handlers.GetHandler{
		Path: "/values/{key}",
		Strg: strg,
	}
	DeleteHandler := &handlers.DeleteHandler{
		Path: "/values/{key}",
		Strg: strg,
	}
	err := server.Run(AddHandler, GetHandler, DeleteHandler)
	if err != nil {
		panic(err)
	}
}
