package tests

import (
	"weatherstation/pkg/weatherstation/model"
	"weatherstation/pkg/weatherstation/observerable"
)

type SelfRemovingObserver struct {
	observable observerable.Observable
	called     bool
}

func (o *SelfRemovingObserver) Update(data model.WeatherInfo) {
	if !o.called {
		o.observable.RemoveObserver(o)
		o.called = true
	}
}
