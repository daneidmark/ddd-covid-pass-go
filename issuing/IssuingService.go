package issuing

import (
	"fmt"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
	"github.com/daneidmark/ddd-covid-pass-go/eventbus"
)

type IssuingSaga struct {
	Eh eventbus.EventHandler
}

func (e *IssuingSaga) Register(bus *eventbus.InMemEventBus) {
	bus.Subscribe("topic", e.Eh)
}

func (e *IssuingSaga) Consume() {
	for {
		select {
		case d := <-e.Eh:
			go handleEvent(d)
		}
	}
}

func handleEvent(event cqrs.Event) {
	fmt.Printf("Envelope: %v; DataEvent: %v\n", event, event.Data)
	switch event.Data.(type) {
	case *covid.Registered:
		createCovidPass(covid.PatientReference(event.AggregateId))
	case *covid.SecondVaccineTaken:
		markAsEligible(covid.PatientReference(event.AggregateId))
	}
}

func createCovidPass(ref covid.PatientReference) {
	fmt.Println("Should create non elibible covid pass")
}

func markAsEligible(ref covid.PatientReference) {
	fmt.Println("Should create mark covid pass ass elibible on second vaccination")
}
