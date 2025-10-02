package observers

import "weatherstation/pkg/weatherstation/model"

type Observer interface {
	Update(data model.WeatherInfo)
}
