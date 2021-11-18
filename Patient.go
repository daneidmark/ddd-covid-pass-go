package covid

import (
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
)

// Domain entities
type PersonalNumber string

type Patient struct {
	cqrs.AggregateRoot
	PersonalNumber PersonalNumber
}

func NewPatient(pn PersonalNumber) Patient {
	p := Patient{}
	p.SetId(cqrs.AggregateId(pn))
	p.Register(pn)
	return p
}

func (p *Patient) Register(n PersonalNumber) {
	p.ApplyNew(p, &Registered{PersonalNumber: n})
}

func (patient *Patient) Transition(e cqrs.Event) {
	switch e := e.Data.(type) {
	case *Registered:
		patient.PersonalNumber = e.PersonalNumber
	}
}

// Events
type Registered struct {
	PersonalNumber PersonalNumber
}

type PatientRepository interface {
	Store(p Patient) error
	Find(pn PersonalNumber) Patient
}
