package main

import (
	"client-server-api/internal/services"
	"log"
)

func main() {
	cotacao, err := services.BuscaCotacaoClient()
	if err != nil {
		log.Fatal(err)
	}

	err = services.GravaArquivoCotacao(cotacao)
	if err != nil {
		log.Fatalf("Erro ao gravar arquico de cotação: %v", err)
	}
}
