package main

import (
	"client-server-api/internal/services"
	"fmt"
	"log"
)

func main() {
	cotacao, err := services.BuscaCotacaoClient()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(cotacao)
}
