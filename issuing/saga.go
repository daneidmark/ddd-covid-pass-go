package issuing

import (
	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
	"github.com/daneidmark/ddd-covid-pass-go/eventbus"
	"github.com/daneidmark/ddd-covid-pass-go/inmemory"
	"github.com/google/uuid"
)

type Saga struct {
	EventHandler eventbus.EventHandler
	Repo         inmemory.IssuingSagaRepository
	Service      Service
}

func (e *Saga) Register(bus *eventbus.InMemEventBus) {
	bus.Subscribe("topic", e.EventHandler)
}

func (i *Saga) Consume() {
	for {
		select {
		case e := <-i.EventHandler:
			go i.handleEvent(e)
		}
	}
}

func (e *Saga) handleEvent(event cqrs.Event) {
	switch event.Data.(type) {
	case *covid.Registered:
		e.createCovidPass(covid.PersonalNumber(event.AggregateId))
	case *covid.SecondVaccineTaken:
		e.markAsEligible(covid.PersonalNumber(event.AggregateId))
	}
}

func (e *Saga) createCovidPass(pn covid.PersonalNumber) {
	id := covid.CovidPassId(uuid.New())
	ref := covid.PatientReference(pn)
	e.Repo.Store(ref, id)
	e.Service.CreateCovidPass(id, ref)
}

func (e *Saga) markAsEligible(ref covid.PersonalNumber) {
	id := e.Repo.Find(covid.PatientReference(ref))
	e.Service.MarkAsEligible(id)
}
