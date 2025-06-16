package main

import (
	"fmt"

	"github.com/Nikita213-hub/simpleInMemStorage/internal/storage"
	"github.com/Nikita213-hub/simpleInMemStorage/server"
)

func main() {
	fmt.Println("Basic setup")
	server := server.NewServer("localhost", "8080")
	strg := storage.NewStorage()
	_ = strg
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
