package services

import (
	"context"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/dto"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/entity"
)

type CepServiceInterface interface {
	GetCepViaCep(ctx context.Context, cep string) (*entity.PostalAddress, error)
}

type WeatherServiceInterface interface {
	GetWeather(ctx context.Context, city string) (*entity.Weather, error)
}

type ExternalCallServiceInterface interface {
	GetExternalCall(ctx context.Context) (*dto.WeatherOutputDto, error)
	SetCep(cep string)
}
