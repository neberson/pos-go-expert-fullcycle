package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/infra/web"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/infra/web/webserver"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/services"
)

const portServer = ":8181"

func main() {
	apiKey := os.Getenv("WEATHER_API_KEY")
	cepService := services.NewCepService()
	weatherService := services.NewWeatherService(apiKey)
	externalCall := services.NewExternalCallService("")
	webWeatherHandler := web.NewWebWeatherHandler(cepService, weatherService, externalCall)

	webserver := webserver.NewWebServer(portServer)
	webserver.AddHandler(http.MethodGet, "/weather/{id}", webWeatherHandler.GetWeatherHandler)
	fmt.Println("Starting server on port", portServer)
	webserver.Start()
}
