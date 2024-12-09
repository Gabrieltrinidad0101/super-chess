package utils

import "sync"

type FuncType func(interface{})

type EventBus struct {
	subscribers map[string][]FuncType
}

var once sync.Once
var eventBus *EventBus

func GetEventBus() *EventBus {
	once.Do(func() {
		eventBus = &EventBus{
			subscribers: make(map[string][]FuncType),
		}
	})
	return eventBus
}

func (e *EventBus) Subscribe(event string, subscriber FuncType) {
	e.subscribers[event] = append(e.subscribers[event], subscriber)
}

func (e *EventBus) Publish(event string, data interface{}) {
	for _, subscriber := range e.subscribers[event] {
		subscriber(data)
	}
}
