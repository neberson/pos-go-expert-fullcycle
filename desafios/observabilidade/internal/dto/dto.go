package dto

type CepInputDto struct {
	Cep string `json:"cep"`
}

type WeatherOutputDto struct {
	TemperatureC float64 `json:"temp_C"`
	TemperatureF float64 `json:"temp_F"`
	TemperatureK float64 `json:"temp_K"`
}
