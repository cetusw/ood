package main

import (
	"weatherstation/pkg/weatherstation/model"
	"weatherstation/pkg/weatherstation/observerable"
	"weatherstation/pkg/weatherstation/observers"
)

func main() {
	wd := observerable.NewWeatherData("station1")
	wd2 := observerable.NewWeatherData("station2")

	var display observers.Display
	statsDisplay := observers.NewStatsDisplay()
	wd.RegisterObserver(&display, 1)
	wd.RegisterObserver(statsDisplay, 0)
	wd2.RegisterObserver(statsDisplay, 2)

	wd.SetMeasurements(model.WeatherInfo{
		Temperature: 3,
		Humidity:    0.7,
		Pressure:    760,
	})
	wd.SetMeasurements(model.WeatherInfo{
		Temperature: 4,
		Humidity:    0.8,
		Pressure:    761,
	})

	wd.RemoveObserver(statsDisplay)

	wd2.SetMeasurements(model.WeatherInfo{
		Temperature: 10,
		Humidity:    0.8,
		Pressure:    761,
	})
	wd.SetMeasurements(model.WeatherInfo{
		Temperature: -10,
		Humidity:    0.8,
		Pressure:    761,
	})
}
