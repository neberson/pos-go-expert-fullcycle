package entity

type Weather struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func NewWeather() *Weather {
	return &Weather{}
}

func (w *Weather) ToFahrenheit() float64 {
	return (w.Current.TempC * 1.8) + 32
}

func (w *Weather) ToKelvin() float64 {
	return w.Current.TempC + 273.15
}
