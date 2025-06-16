package main

import (
	"fmt"

	"github.com/Nikita213-hub/simpleInMemStorage/server"
)

func main() {
	fmt.Println("Basic setup")
	server := server.NewServer("localhost", "8080")
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
