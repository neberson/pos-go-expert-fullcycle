package main

import (
	"fmt"
	"os"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/services"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/usecase"
)

func main() {
	apiKey := os.Getenv("WEATHER_API_KEY")

	cepService := services.NewCepService()
	weatherService := services.NewWeatherService(apiKey)
	WeatherUseCase := usecase.NewGetWeatherUseCase(cepService, weatherService)
	inputCep := usecase.CepInputDto{Cep: "75915000"}
	_, err := WeatherUseCase.Execute(inputCep)
	if err != nil {
		panic(err)
	}
	fmt.Println("Weather fetched successfully")
}
