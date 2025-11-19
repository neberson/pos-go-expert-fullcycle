package main

import (
	"fmt"
	"net/http"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/infra/web"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/infra/web/webserver"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/services"
	"github.com/spf13/viper"
)

const portServer = ":8080"

func init() {
	viper.AutomaticEnv()
}

func main() {
	apiKey := viper.GetString("WEATHER_API_KEY")
	externalCallUrl := viper.GetString("EXTERNAL_CALL_URL")

	cepService := services.NewCepService()
	weatherService := services.NewWeatherService(apiKey)
	externalCall := services.NewExternalCallService(externalCallUrl)
	webWeatherHandler := web.NewWebWeatherHandler(cepService, weatherService, externalCall)

	webserver := webserver.NewWebServer(portServer)
	webserver.AddHandler(http.MethodPost, "/weather", webWeatherHandler.PostWeatherHandler)
	fmt.Println("Starting server on port", portServer)
	webserver.Start()
}
