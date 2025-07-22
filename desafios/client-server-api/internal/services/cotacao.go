package services

import (
	"client-server-api/pkg/entity"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	urlCotacao = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
)

func BuscaCotacao() (*entity.Cotacao, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", urlCotacao, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao realizar a requisição da cotação: %v", err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao realizar a requisição da cotação: %v", err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("erro ao realizar a requisição da cotação: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo da requisição: %v", err.Error())
	}

	var cotacao entity.Cotacao
	if err = json.Unmarshal(body, &cotacao); err != nil {
		return nil, fmt.Errorf("erro na conversão do corpo da requisição: %v", err.Error())
	}

	return &cotacao, nil
}
