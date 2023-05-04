package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/upikoth/starter-go/internal/app/apiserver"
)

// @title   Starter API
// @host    localhost:8080
func main() {
	godotenv.Load()

	config, err := apiserver.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	server := apiserver.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
