package observers

import (
	"fmt"
	"weatherstation/pkg/weatherstation/model"
)

type Display struct{}

func (d *Display) Update(data model.WeatherInfo) {
	fmt.Println("Current Temp", data.Temperature)
	fmt.Println("Current Hum", data.Humidity)
	fmt.Println("Current Pressure", data.Pressure)
	fmt.Println("----------------")
}
