package tests

import "weatherstation/pkg/weatherstation/model"

type MultiSourceRecorder struct {
	log []struct {
		SourceID string
		Temp     float64
	}
}

func (m *MultiSourceRecorder) Update(sourceID string, data model.WeatherInfo) {
	m.log = append(m.log, struct {
		SourceID string
		Temp     float64
	}{SourceID: sourceID, Temp: data.Temperature})
}
