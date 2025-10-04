package tests

import (
	"testing"
	"weatherstation/pkg/weatherstation/model"
	"weatherstation/pkg/weatherstation/observerable"
)

func TestNotifyObserversSafeRemovalDuringUpdate(t *testing.T) {
	observer := observerable.NewObservable()
	selfRemover := &SelfRemovingObserver{observable: observer}

	observer.RegisterObserver(selfRemover, 0)

	observer.NotifyObservers("station1", model.WeatherInfo{
		Temperature: 20.0,
		Humidity:    0.5,
		Pressure:    760.0,
	})

	if !selfRemover.called {
		t.Error("Expected SelfRemovingObserver to be called")
	}
}
