package observerable

import (
	"weatherstation/pkg/weatherstation/model"
)

type WeatherData struct {
	Observable
	sourceID    string
	temperature float64
	humidity    float64
	pressure    float64
}

func NewWeatherData(sourceID string) *WeatherData {
	return &WeatherData{
		Observable:  NewObservable(),
		sourceID:    sourceID,
		temperature: 0.0,
		humidity:    0.0,
		pressure:    760.0,
	}
}

func (wd *WeatherData) MeasurementsChanged() {
	wd.NotifyObservers(wd.sourceID, wd.getCurrentData())
}

func (wd *WeatherData) SetMeasurements(measurements model.WeatherInfo) {
	wd.humidity = measurements.Humidity
	wd.temperature = measurements.Temperature
	wd.pressure = measurements.Pressure
	wd.NotifyObservers(wd.sourceID, wd.getCurrentData())
}

func (wd *WeatherData) getCurrentData() model.WeatherInfo {
	return model.WeatherInfo{
		Temperature: wd.temperature,
		Humidity:    wd.humidity,
		Pressure:    wd.pressure,
	}
}
