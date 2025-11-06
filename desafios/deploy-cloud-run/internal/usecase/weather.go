package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/entity"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/services"
)

type CepInputDto struct {
	Cep string `json:"cep"`
}

type WeatherOutputDto struct {
	TemperatureC float64 `json:"temp_C"`
	TemperatureF float64 `json:"temp_F"`
	TemperatureK float64 `json:"temp_K"`
}

type GetWeatherUseCase struct {
	cepService     services.CepServiceInterface
	weatherService services.WeatherServiceInterface
}

func NewGetWeatherUseCase(
	cepServiceInterface services.CepServiceInterface,
	weatherServiceInterface services.WeatherServiceInterface,
) *GetWeatherUseCase {
	return &GetWeatherUseCase{
		cepService:     cepServiceInterface,
		weatherService: weatherServiceInterface,
	}
}

func (w *GetWeatherUseCase) Execute(input CepInputDto) (WeatherOutputDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	cep := entity.NewCep(input.Cep)
	if err := cep.Validate(); err != nil {
		return WeatherOutputDto{}, err
	}

	postalAddress, err := w.cepService.GetCepViaCep(ctx, cep.Cep)
	if err != nil {
		return WeatherOutputDto{}, err
	}

	weather, err := w.weatherService.GetWeather(ctx, postalAddress.Localidade)
	if err != nil {
		return WeatherOutputDto{}, err
	}

	fmt.Println(weather)
	fmt.Println(postalAddress)

	return WeatherOutputDto{}, nil
}
