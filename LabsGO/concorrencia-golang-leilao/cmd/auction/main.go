package main

import (
	"concorrencia-golang-leilao/configuration/database/mongodb"
	"context"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	_, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
