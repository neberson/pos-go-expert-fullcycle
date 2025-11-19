package main

import (
	"fmt"
	"os"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/infra/web"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/infra/web/webserver"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/services"
)

const portServer = ":8080"

func main() {
	apiKey := os.Getenv("WEATHER_API_KEY")
	cepService := services.NewCepService()
	weatherService := services.NewWeatherService(apiKey)
	webWeatherHandler := web.NewWebWeatherHandler(cepService, weatherService)

	webserver := webserver.NewWebServer(portServer)
	webserver.AddHandler("/weather/{id}", webWeatherHandler.GetWeatherHandler)
	fmt.Println("Starting server on port", portServer)
	webserver.Start()
}
