package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/infra/web"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/infra/web/webserver"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/services"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/pkg/provider"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const portServer = ":8080"

func init() {
	viper.AutomaticEnv()
}

func main() {
	apiKey := viper.GetString("WEATHER_API_KEY")
	externalCallUrl := viper.GetString("EXTERNAL_CALL_URL")

	collectorURL := viper.GetString("OTEL_EXPORTER_OTLP_ENDPOINT")
	if collectorURL == "" {
		collectorURL = "localhost:4317"
	}

	tracerProvider, err := provider.InitProvider("service-a", collectorURL)
	if err != nil {
		log.Fatalf("failed to initialization provider: %v", err)
	}

	defer func() {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			log.Printf("Error disabling tracer provider.: %v", err)
		}
	}()

	cepService := services.NewCepService()
	weatherService := services.NewWeatherService(apiKey)
	externalCall := services.NewExternalCallService(externalCallUrl)
	webWeatherHandler := web.NewWebWeatherHandler(cepService, weatherService, externalCall)

	webserver := webserver.NewWebServer(portServer)
	otelHandler := otelhttp.NewHandler(http.HandlerFunc(webWeatherHandler.PostWeatherHandler), "PostWeatherHandler")
	webserver.AddHandler(http.MethodPost, "/weather", otelHandler.ServeHTTP)

	fmt.Println("Starting server on port", portServer)
	webserver.Start()
}
