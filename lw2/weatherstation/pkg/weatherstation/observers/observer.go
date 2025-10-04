package observers

import "weatherstation/pkg/weatherstation/model"

type Observer interface {
	Update(sourceID string, data model.WeatherInfo)
}
