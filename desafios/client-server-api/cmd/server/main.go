package main

import (
	"client-server-api/internal/services"
	"encoding/json"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", CotacaoHandler)
	http.ListenAndServe(":8080", mux)
}

func CotacaoHandler(w http.ResponseWriter, r *http.Request) {
	cotacao, err := services.BuscaCotacao()
	if err != nil {
		w.Write([]byte("Erro na requisição contate o suporte técnico."))
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)
}
