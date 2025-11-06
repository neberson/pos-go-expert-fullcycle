package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/entity"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/services"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/usecase"
)

type WebWeatherHandler struct {
	cepService     services.CepServiceInterface
	weatherService services.WeatherServiceInterface
}

func NewWebWeatherHandler(
	cepService services.CepServiceInterface,
	weatherService services.WeatherServiceInterface,
) *WebWeatherHandler {
	return &WebWeatherHandler{
		cepService:     cepService,
		weatherService: weatherService,
	}
}

func (h *WebWeatherHandler) GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cep := chi.URLParam(r, "id")
	cepInput := entity.NewCep(cep)

	if err := cepInput.Validate(); errors.Is(err, entity.ErrInvalidCep) {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	WeatherUseCase := usecase.NewGetWeatherUseCase(h.cepService, h.weatherService)
	inputCepDto := usecase.CepInputDto{Cep: cep}
	weather, err := WeatherUseCase.Execute(inputCepDto)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(weather)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
