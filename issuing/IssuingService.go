package issuing

import (
	"fmt"

	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
	"github.com/daneidmark/ddd-covid-pass-go/eventbus"
)

type IssuingEventHandler struct {
	Eh eventbus.EventHandler
}

func (e *IssuingEventHandler) Register(bus *eventbus.InMemEventBus) {
	bus.Subscribe("topic", e.Eh)
}

func (e *IssuingEventHandler) Consume() {
	fmt.Println("Before")
	for {
		select {
		case d := <-e.Eh:
			go printDataEvent("ch1", d)
		}
	}
	fmt.Println("After")


}

func printDataEvent(ch string, data cqrs.Event) {
	fmt.Printf("Channel: %s; Envelope: %v; DataEvent: %v\n", ch, data, data.Data)
}
