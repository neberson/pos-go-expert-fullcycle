package services

import (
	"client-server-api/pkg/entity"
	"client-server-api/pkg/entity/dto"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	urlCotacao    = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	urlBaseClient = "http://localhost:8080/cotacao"
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

func BuscaCotacaoClient() (*dto.CotacaoDto, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo da requisição: %v", err.Error())
	}

	var cotacao dto.CotacaoDto
	if err = json.Unmarshal(body, &cotacao); err != nil {
		return nil, fmt.Errorf("erro na conversão do corpo da requisição: %v", err.Error())
	}

	return &cotacao, nil
}
