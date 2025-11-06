package entity

import "testing"

func TestNewWeather(t *testing.T) {
	weather := NewWeather()
	if weather == nil {
		t.Error("NewWeather() should return a non-nil Weather instance")
	}
	if weather.Current.TempC != 0 {
		t.Errorf("Expected TempC to be 0, got %f", weather.Current.TempC)
	}
}

func TestWeather_ToFahrenheit(t *testing.T) {
	tests := []struct {
		name  string
		tempC float64
		want  float64
	}{
		{
			name:  "Zero Celsius",
			tempC: 0,
			want:  32,
		},
		{
			name:  "100 Celsius",
			tempC: 100,
			want:  212,
		},
		{
			name:  "Negative temperature",
			tempC: -40,
			want:  -40,
		},
		{
			name:  "Room temperature",
			tempC: 25,
			want:  77,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWeather()
			w.Current.TempC = tt.tempC
			got := w.ToFahrenheit()
			if got != tt.want {
				t.Errorf("ToFahrenheit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeather_ToKelvin(t *testing.T) {
	tests := []struct {
		name  string
		tempC float64
		want  float64
	}{
		{
			name:  "Zero Celsius",
			tempC: 0,
			want:  273.15,
		},
		{
			name:  "100 Celsius",
			tempC: 100,
			want:  373.15,
		},
		{
			name:  "Negative temperature",
			tempC: -273.15,
			want:  0,
		},
		{
			name:  "Room temperature",
			tempC: 25,
			want:  298.15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWeather()
			w.Current.TempC = tt.tempC
			got := w.ToKelvin()
			if got != tt.want {
				t.Errorf("ToKelvin() = %v, want %v", got, tt.want)
			}
		})
	}
}
