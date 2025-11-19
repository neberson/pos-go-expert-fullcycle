package services

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/dto"
)

type ExternalCallService struct {
	url string
	cep string
}

func (e *ExternalCallService) SetCep(cep string) {
	e.cep = cep
}

func NewExternalCallService(url string) *ExternalCallService {
	return &ExternalCallService{
		url: url,
	}
}

func (e *ExternalCallService) GetUrl() *string {
	return &e.url
}

func (e *ExternalCallService) GetExternalCall(ctx context.Context) (*dto.WeatherOutputDto, error) {

	weather := dto.WeatherOutputDto{}

	urlCall := e.url + "/" + e.cep

	log.Println(urlCall)
	req, err := http.NewRequestWithContext(ctx, "GET", urlCall, nil)
	if err != nil {
		return &weather, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &weather, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return &weather, ErrToManyRequests
	}

	if resp.StatusCode != http.StatusOK {
		return &weather, ErrGeneric
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &weather, err
	}

	if strings.Contains(string(body), `error`) && strings.Contains(string(body), `"code":1006`) {
		return &weather, ErrNotFound
	}

	err = json.Unmarshal(body, &weather)
	if err != nil {
		return &weather, err
	}

	log.Println(string(body))

	return &weather, nil
}
