package observerable

import (
	"sort"
	"weatherstation/pkg/weatherstation/model"
	"weatherstation/pkg/weatherstation/observers"
)

type Observable interface {
	RegisterObserver(observer observers.Observer, priority int)
	NotifyObservers(data model.WeatherInfo)
	RemoveObserver(observer observers.Observer)
}

type observerEntry struct {
	observer observers.Observer
	priority int
}

type observable struct {
	observers       map[observers.Observer]int
	sortedObservers []observerEntry
}

func NewObservable() Observable {
	return &observable{
		observers: make(map[observers.Observer]int),
	}
}

func (o *observable) RegisterObserver(observer observers.Observer, priority int) {
	if _, exists := o.observers[observer]; exists {
		return
	}
	o.observers[observer] = priority
	o.sortObservers()
}

func (o *observable) NotifyObservers(data model.WeatherInfo) {
	copyList := make([]observerEntry, len(o.sortedObservers))
	copy(copyList, o.sortedObservers)

	for _, entry := range copyList {
		entry.observer.Update(data)
	}
}

func (o *observable) RemoveObserver(observer observers.Observer) {
	if _, exists := o.observers[observer]; !exists {
		return
	}
	delete(o.observers, observer)
	o.sortObservers()
}

func (o *observable) sortObservers() {
	o.sortedObservers = make([]observerEntry, 0, len(o.observers))
	for observer, priority := range o.observers {
		o.sortedObservers = append(o.sortedObservers, observerEntry{
			observer: observer,
			priority: priority,
		})
	}

	sort.Slice(o.sortedObservers, func(i, j int) bool {
		return o.sortedObservers[i].priority > o.sortedObservers[j].priority
	})
}
