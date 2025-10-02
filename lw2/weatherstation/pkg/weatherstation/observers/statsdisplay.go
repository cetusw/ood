package observers

import (
	"fmt"
	"weatherstation/pkg/weatherstation/model"
)

type StatsDisplay struct {
	temperatureStats model.Stats
	humidityStats    model.Stats
	pressureStats    model.Stats
	countAcc         uint
}

func NewStatsDisplay() *StatsDisplay {
	return &StatsDisplay{
		temperatureStats: model.NewStats(),
		humidityStats:    model.NewStats(),
		pressureStats:    model.NewStats(),
		countAcc:         0,
	}
}

func (s *StatsDisplay) Update(data model.WeatherInfo) {
	s.temperatureStats.Update(data.Temperature)
	s.humidityStats.Update(data.Humidity)
	s.pressureStats.Update(data.Pressure)
	s.countAcc++

	s.temperatureStats.Print("Temperature", s.countAcc)
	s.humidityStats.Print("Humidity", s.countAcc)
	s.pressureStats.Print("Pressure", s.countAcc)
	fmt.Println("----------------")
}
