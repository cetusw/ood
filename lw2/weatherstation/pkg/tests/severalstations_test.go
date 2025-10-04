package tests

import (
	"testing"
	"weatherstation/pkg/weatherstation/model"
	"weatherstation/pkg/weatherstation/observerable"
)

func TestMultipleStationsIndependentNotifications(t *testing.T) {
	stationA := observerable.NewWeatherData("StationA")
	stationB := observerable.NewWeatherData("StationB")

	recorder := &MultiSourceRecorder{}

	stationA.RegisterObserver(recorder, 1)
	stationB.RegisterObserver(recorder, 0)

	stationA.SetMeasurements(model.WeatherInfo{Temperature: 20.0, Humidity: 0.5, Pressure: 760})
	stationB.SetMeasurements(model.WeatherInfo{Temperature: -5.0, Humidity: 0.9, Pressure: 770})

	if len(recorder.log) != 2 {
		t.Fatalf("Expected 2 updates, got %d", len(recorder.log))
	}

	if recorder.log[0].SourceID != "StationA" || recorder.log[0].Temp != 20.0 {
		t.Errorf("First update should be from StationA with temp 20.0, got %+v", recorder.log[0])
	}
	if recorder.log[1].SourceID != "StationB" || recorder.log[1].Temp != -5.0 {
		t.Errorf("Second update should be from StationB with temp -5.0, got %+v", recorder.log[1])
	}
}

func TestMultipleStationsRemoveFromOneOnly(t *testing.T) {
	stationA := observerable.NewWeatherData("A")
	stationB := observerable.NewWeatherData("B")

	recorder := &MultiSourceRecorder{}

	stationA.RegisterObserver(recorder, 0)
	stationB.RegisterObserver(recorder, 0)

	stationA.RemoveObserver(recorder)

	stationA.SetMeasurements(model.WeatherInfo{Temperature: 100})
	stationB.SetMeasurements(model.WeatherInfo{Temperature: 200})

	if len(recorder.log) != 1 {
		t.Fatalf("Expected 1 update (from B), got %d", len(recorder.log))
	}
	if recorder.log[0].SourceID != "B" || recorder.log[0].Temp != 200 {
		t.Errorf("Expected update from B with temp 200, got %+v", recorder.log[0])
	}
}
