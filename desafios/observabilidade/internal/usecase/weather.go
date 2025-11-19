package usecase

import (
	"context"
	"time"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/dto"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/entity"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/services"
)

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

func (w *GetWeatherUseCase) Execute(input dto.CepInputDto) (dto.WeatherOutputDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	cep := entity.NewCep(input.Cep)
	if err := cep.Validate(); err != nil {
		return dto.WeatherOutputDto{}, err
	}

	postalAddress, err := w.cepService.GetCepViaCep(ctx, cep.Cep)
	if err != nil {
		return dto.WeatherOutputDto{}, err
	}

	weather, err := w.weatherService.GetWeather(ctx, postalAddress.Localidade)
	if err != nil {
		return dto.WeatherOutputDto{}, err
	}

	weatherOutputDto := dto.WeatherOutputDto{
		TemperatureC: weather.Current.TempC,
		TemperatureF: weather.ToFahrenheit(),
		TemperatureK: weather.ToKelvin(),
	}

	return weatherOutputDto, nil
}
