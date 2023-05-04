package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/upikoth/starter-go/internal/app/apiserver"
)

// @title   Starter API
// @host    localhost:8080.
func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Println("No .env file")
	}

	config, configErr := apiserver.NewConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}

	server := apiserver.New(config)

	if serverErr := server.Start(); serverErr != nil {
		log.Fatal(serverErr)
	}
}
