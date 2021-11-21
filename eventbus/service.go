package eventbus

import (
	"fmt"
	"sync"

	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
)

type Service interface {
	Publish(t Topic, e cqrs.Event)
}

type NoopService struct {
}

func (*NoopService) Publish(t Topic, e cqrs.Event) {
	fmt.Printf("Publishing %+v\n", e)
}

type EventHandler chan cqrs.Event
type Topic string

type InMemEventBus struct {
	Subscribers map[Topic][]EventHandler
	rm sync.RWMutex
}

func (eb *InMemEventBus) Subscribe(t Topic, eh EventHandler) {
	eb.Subscribers[t] = append(eb.Subscribers[t], eh)
}

func (eb *InMemEventBus) Publish(t Topic, e cqrs.Event) {
	eb.rm.RLock()
	s := eb.Subscribers[t]
	go func(event cqrs.Event, handlers []EventHandler) {
		for _, eh := range handlers {
			eh <- e
		}
	}(e, s)
	eb.rm.RUnlock()
}
