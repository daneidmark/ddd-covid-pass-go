package covid

import (
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
)

// Domain entities
type CovidPassId string
type PatientReference string

type covidPass struct {
	cqrs.AggregateRoot
	CovidPassId      CovidPassId
	PatientReference PatientReference
}

func NewCovidPass(id CovidPassId, ref PatientReference) covidPass {
	p := covidPass{}
	p.SetId(cqrs.AggregateId(id))
	p.New(ref)
	return p
}

func (c *covidPass) New(ref PatientReference) {
	c.ApplyNew(c, &Created{PatientReference: ref})
}

func (c *covidPass) Transition(e cqrs.Event) {
	switch e := e.Data.(type) {
	case *Created:
		c.PatientReference = e.PatientReference
	}
}

// Events
type Created struct {
	PatientReference PatientReference
}

// Repository
type CovidPassRepository interface {
	Store(p covidPass) error
	Find(pn CovidPassId) covidPass
}
