package covid

import (
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
	"github.com/google/uuid"
)

// Domain entities
type CovidPassId uuid.UUID
type PatientReference string

type CovidPass struct {
	cqrs.AggregateRoot
	CovidPassId      CovidPassId
	PatientReference PatientReference
	Eligible         bool
}

func NewCovidPass(id CovidPassId, ref PatientReference) CovidPass {
	p := CovidPass{}
	p.SetId(cqrs.AggregateId(uuid.UUID(id).String()))
	p.New(ref)
	return p
}

func (c *CovidPass) New(ref PatientReference) {
	c.ApplyNew(c, &Created{PatientReference: ref, Eligible: false})
}

func (c *CovidPass) MarkAsEligible() {
	c.ApplyNew(c, &Eligible{})
}

func (c *CovidPass) Transition(e cqrs.Event) {
	switch e := e.Data.(type) {
	case *Created:
		c.PatientReference = e.PatientReference
		c.Eligible = e.Eligible
	case *Eligible:
		c.Eligible = true
	}
}

// Events
type Created struct {
	PatientReference PatientReference
	Eligible         bool
}

type Eligible struct {
}

// Repository
type CovidPassRepository interface {
	Store(p CovidPass) error
	Find(pn CovidPassId) CovidPass
}
