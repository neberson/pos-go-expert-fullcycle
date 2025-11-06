package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/entity"
)

const urlViaCep = "https://viacep.com.br/ws/%s/json/"

type CepService struct{}

func NewCepService() *CepService {
	return &CepService{}
}

func (c CepService) GetCepViaCep(ctx context.Context, cep string) (*entity.PostalAddress, error) {
	postalAddress := entity.NewPostalAddress()

	url := fmt.Sprintf(urlViaCep, cep)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return postalAddress, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return postalAddress, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return postalAddress, ErrToManyRequests
	}

	if resp.StatusCode != http.StatusOK {
		return postalAddress, ErrGeneric
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return postalAddress, err
	}

	if strings.Contains(string(body), `erro`) {
		return postalAddress, ErrNotFound
	}

	err = json.Unmarshal(body, postalAddress)
	if err != nil {
		return postalAddress, err
	}
	return postalAddress, nil
}
