package issuing

import (
	"fmt"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
	"github.com/daneidmark/ddd-covid-pass-go/eventbus"
	"github.com/daneidmark/ddd-covid-pass-go/inmemory"
	"github.com/google/uuid"
)

type IssuingSaga struct {
	EventHandler eventbus.EventHandler
	Repo         inmemory.IssuingSagaRepository
	Service      Service
}

func (e *IssuingSaga) Register(bus *eventbus.InMemEventBus) {
	bus.Subscribe("topic", e.EventHandler)
}

func (i *IssuingSaga) Consume() {
	for {
		select {
		case e := <-i.EventHandler:
			go i.handleEvent(e)
		}
	}
}

func (e *IssuingSaga) handleEvent(event cqrs.Event) {
	fmt.Printf("Envelope: %v; DataEvent: %v\n", event, event.Data)
	switch event.Data.(type) {
	case *covid.Registered:
		e.createCovidPass(covid.PersonalNumber(event.AggregateId))
	case *covid.SecondVaccineTaken:
		e.markAsEligible(covid.PersonalNumber(event.AggregateId))
	}
}

func (e *IssuingSaga) createCovidPass(pn covid.PersonalNumber) {
	id := covid.CovidPassId(uuid.New())
	ref := covid.PatientReference(pn)
	e.Repo.Store(ref, id)
	e.Service.CreateCovidPass(id, ref)
}

func (e *IssuingSaga) markAsEligible(ref covid.PersonalNumber) {
	id := e.Repo.Find(covid.PatientReference(ref))
	e.Service.MarkAsEligible(id)
}

type Service struct {
	Repo covid.CovidPassRepository
}

func (s *Service) CreateCovidPass(id covid.CovidPassId, ref covid.PatientReference) {
	fmt.Println("Should create non elibible covid pass")
	c := covid.CovidPass{}
	c.SetId(cqrs.AggregateId(uuid.UUID(id).String()))
	c.New(ref)
	s.Repo.Store(c)
}

func (s *Service) MarkAsEligible(id covid.CovidPassId) {
	fmt.Println("Should create mark covid pass ass elibible on second vaccination")
	c := s.Repo.Find(id)
	c.MarkAsEligible()
	s.Repo.Store(c)
}
