package observers

import (
	"fmt"
	"weatherstation/pkg/weatherstation/model"
)

type Display struct{}

func (d *Display) Update(sourceID string, data model.WeatherInfo) {
	fmt.Printf("[%s] Current Temperature: %2.f\n", sourceID, data.Temperature)
	fmt.Printf("[%s] Current Humidity: %2.f\n", sourceID, data.Humidity)
	fmt.Printf("[%s] Current Pressure: %2.f\n", sourceID, data.Pressure)
	fmt.Printf("----------------\n")
}
