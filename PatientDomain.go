package main

// Domain entities
type PersonalNumber string

type Patient struct {
	AggregateRoot
	PersonalNumber PersonalNumber
}

func (p *Patient) register(n PersonalNumber) {
	p.ApplyNew(p, &Registered{PersonalNumber: n})
}

func (patient *Patient) Transition(e Event) {
	switch e := e.Data.(type) {
	case *Registered:
		patient.PersonalNumber = e.PersonalNumber
	}
}

// Events
type Registered struct {
	PersonalNumber PersonalNumber
}

// Persistance
type PatientReader interface {
	Find(pn PersonalNumber) (p Patient)
}

type PatientStorer interface {
	Store(p Patient)
}
