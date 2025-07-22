package main

import (
	database "client-server-api/internal/infra/database/sqlite"
	"client-server-api/internal/services"
	"encoding/json"
	"log"
	"net/http"
)

func init() {
	if err := database.CreateTable(); err != nil {
		log.Println(err.Error())
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", CotacaoHandler)
	http.ListenAndServe(":8080", mux)
	defer database.DbClose()
}

func CotacaoHandler(w http.ResponseWriter, r *http.Request) {
	cotacao, err := services.BuscaCotacao()
	if err != nil {
		log.Println("Erro na requisição contate o suporte técnico", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = database.InsertCotacao(cotacao)
	if err != nil {
		log.Printf("Erro ao gravar a cotação: %v\n", err.Error())
		w.Write([]byte("Erro ao gravar a cotação: %v\n" + err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)
}
