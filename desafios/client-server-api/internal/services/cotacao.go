package services

import (
	"client-server-api/pkg/entity"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	urlCotacao = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
)

func BuscaCotacao() (*entity.Cotacao, error) {
	resp, err := http.Get(urlCotacao)
	if err != nil {
		return nil, fmt.Errorf("erro ao realizar a requisição da cotação: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo da requisição: %v", err)
	}
	var cotacao entity.Cotacao

	if err = json.Unmarshal(body, &cotacao); err != nil {
		return nil, fmt.Errorf("erro na conversão do corpo da requisição: %v", err)
	}
	return &cotacao, nil
}
