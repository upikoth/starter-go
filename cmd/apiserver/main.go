package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/upikoth/starter-go/internal/app/apiserver"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := apiserver.NewConfig()
	server := apiserver.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
