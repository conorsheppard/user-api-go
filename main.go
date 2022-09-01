package main

import (
	"github.com/conorsheppard/user-api-go/internal/api"
	"log"
)

const (
	serverAddress = `0.0.0.0:8080`
)

func main() {
	server := api.NewServer()
	err := server.Start(serverAddress)
	if err != nil {
		log.Fatalf("unable to start server: %s\n", err)
	}
}
