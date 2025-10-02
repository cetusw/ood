package observers

import (
	"fmt"
	"math"
	"weatherstation/pkg/weatherstation/model"
)

type StatsDisplay struct {
	minTemperature float64
	maxTemperature float64
	accTemperature float64
	countAcc       uint
}

func NewStatsDisplay() *StatsDisplay {
	return &StatsDisplay{
		minTemperature: math.MaxFloat64,
		maxTemperature: -math.MaxFloat64,
	}
}

func (s *StatsDisplay) Update(data model.WeatherInfo) {
	if s.minTemperature > data.Temperature {
		s.minTemperature = data.Temperature
	}
	if s.maxTemperature < data.Temperature {
		s.maxTemperature = data.Temperature
	}
	s.accTemperature += data.Temperature
	s.countAcc++

	fmt.Println("Max Temp", s.maxTemperature)
	fmt.Println("Min Temp", s.minTemperature)
	fmt.Println("Average Temp", (s.accTemperature / float64(s.countAcc)))
	fmt.Println("----------------")
}
