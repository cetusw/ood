package tests

import "weatherstation/pkg/weatherstation/model"

type RecordingObserver struct {
	id  string
	log *[]string
}

func (r *RecordingObserver) Update(data model.WeatherInfo) {
	*r.log = append(*r.log, r.id)
}
