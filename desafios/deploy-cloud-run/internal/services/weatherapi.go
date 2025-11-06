package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/entity"
)

var (
	urlWeatherApi = "http://api.weatherapi.com/v1/current.json?key=%s&q=%s"
)

type WeatherService struct {
	apiKey string
}

func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{
		apiKey: apiKey,
	}
}

func (w *WeatherService) GetWeather(ctx context.Context, city string) (*entity.Weather, error) {

	weather := entity.NewWeather()

	url := fmt.Sprintf(urlWeatherApi, w.apiKey, city)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return weather, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return weather, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return weather, ErrToManyRequests
	}

	if resp.StatusCode != http.StatusOK {
		return weather, ErrGeneric
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return weather, err
	}

	if (strings.Contains(string(body), `code`)) && (strings.Contains(string(body), `1006`)) {
		return weather, ErrNotFound
	}

	err = json.Unmarshal(body, weather)
	if err != nil {
		return weather, err
	}

	return weather, nil
}
