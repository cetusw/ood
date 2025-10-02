package observerable

import (
	"weatherstation/pkg/weatherstation/model"
	"weatherstation/pkg/weatherstation/observers"
)

type Observable interface {
	RegisterObserver(observer observers.Observer)
	NotifyObservers(data model.WeatherInfo)
	RemoveObserver(observer observers.Observer)
}

type observable struct {
	observers map[observers.Observer]int
}

func NewObservable() Observable {
	return &observable{
		observers: make(map[observers.Observer]int),
	}
}

func (o *observable) RegisterObserver(observer observers.Observer) {
	o.observers[observer] = 0
}

func (o *observable) NotifyObservers(data model.WeatherInfo) {
	for observer := range o.observers {
		observer.Update(data)
	}
}

func (o *observable) RemoveObserver(observer observers.Observer) {
	delete(o.observers, observer)
}
