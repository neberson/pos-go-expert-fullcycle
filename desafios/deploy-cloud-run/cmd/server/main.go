package main

import (
	"fmt"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/services"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/usecase"
)

func main() {
	cepService := services.NewCepService()
	WeatherUseCase := usecase.NewGetWeatherUseCase(cepService)
	inputCep := usecase.CepInputDto{Cep: "75915000"}
	_, err := WeatherUseCase.Execute(inputCep)
	if err != nil {
		panic(err)
	}
	fmt.Println("Weather fetched successfully")
}
