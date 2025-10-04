package tests

import (
	"testing"
	"weatherstation/pkg/weatherstation/model"
	"weatherstation/pkg/weatherstation/observerable"
)

func TestNotifyObserversPriorityOrder(t *testing.T) {
	obs := observerable.NewObservable()

	var callOrder []string

	lowPriority := &RecordingObserver{id: "low", log: &callOrder}
	medPriority := &RecordingObserver{id: "med", log: &callOrder}
	highPriority := &RecordingObserver{id: "high", log: &callOrder}

	obs.RegisterObserver(lowPriority, 0)
	obs.RegisterObserver(highPriority, 10)
	obs.RegisterObserver(medPriority, 5)

	callOrder = []string{}
	obs.NotifyObservers("station1", model.WeatherInfo{})

	expected := []string{"high", "med", "low"}

	if len(callOrder) != len(expected) {
		t.Fatalf("Expected %d calls, got %d", len(expected), len(callOrder))
	}

	for i, expectedID := range expected {
		if callOrder[i] != expectedID {
			t.Errorf("Position %d: expected %s, got %s", i, expectedID, callOrder[i])
		}
	}
}

func TestNotifyObserversSamePriorityOrderStable(t *testing.T) {
	obs := observerable.NewObservable()

	var callOrder []string

	a := &RecordingObserver{id: "A", log: &callOrder}
	b := &RecordingObserver{id: "B", log: &callOrder}

	obs.RegisterObserver(a, 5)
	obs.RegisterObserver(b, 5)

	callOrder = []string{}
	obs.NotifyObservers("station1", model.WeatherInfo{})

	if len(callOrder) != 2 {
		t.Errorf("Expected 2 calls, got %d", len(callOrder))
	}

	foundA, foundB := false, false
	for _, id := range callOrder {
		if id == "A" {
			foundA = true
		}
		if id == "B" {
			foundB = true
		}
	}
	if !foundA || !foundB {
		t.Error("Both observers should be notified even with same priority")
	}
}
