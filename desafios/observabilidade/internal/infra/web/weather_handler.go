package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/dto"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/entity"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/services"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/observabilidade/internal/usecase"
)

type WebWeatherHandler struct {
	cepService      services.CepServiceInterface
	weatherService  services.WeatherServiceInterface
	externalCallUrl services.ExternalCallServiceInterface
}

func NewWebWeatherHandler(
	cepService services.CepServiceInterface,
	weatherService services.WeatherServiceInterface,
	externalCallUrl services.ExternalCallServiceInterface,
) *WebWeatherHandler {
	return &WebWeatherHandler{
		cepService:      cepService,
		weatherService:  weatherService,
		externalCallUrl: externalCallUrl,
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
	inputCepDto := dto.CepInputDto{Cep: cep}
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

func (h *WebWeatherHandler) PostWeatherHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var raw map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cepValue, ok := raw["cep"]
	if !ok {
		http.Error(w, entity.ErrInvalidCep.Error(), http.StatusBadRequest)
		return
	}

	cepStr, ok := cepValue.(string)
	if !ok {
		http.Error(w, entity.ErrInvalidCep.Error(), http.StatusBadRequest)
		return
	}

	cepInput := entity.Cep{Cep: cepStr}
	if err := cepInput.Validate(); errors.Is(err, entity.ErrInvalidCep) {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	h.externalCallUrl.SetCep(cepInput.Cep)
	weather, err := h.externalCallUrl.GetExternalCall(r.Context())
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
