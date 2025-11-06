package main

import (
	"os"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/infra/web"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/infra/web/webserver"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/services"
)

func main() {
	apiKey := os.Getenv("WEATHER_API_KEY")
	cepService := services.NewCepService()
	weatherService := services.NewWeatherService(apiKey)
	webWeatherHandler := web.NewWebWeatherHandler(cepService, weatherService)

	webserver := webserver.NewWebServer(":8080")
	webserver.AddHandler("/weather/{id}", webWeatherHandler.GetWeatherHandler)
	webserver.Start()
}
