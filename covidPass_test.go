package covid

import (
	"reflect"
	"testing"
	"time"

	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Implement test Aggregate

// tests

func TestCreateCovidPass(t *testing.T) {

	p := NewCovidPass(CovidPassId(uuid.New()), PatientReference("dan"))

	assert.Equal(t, len(p.UncommittedEvents), 1)
	assert.Equal(t, reflect.TypeOf(p.UncommittedEvents[0].Data).String(), "*covid.Created")

	e := p.UncommittedEvents[0].Data.(*Created)
	assert.Equal(t, e.Eligible, false)
	assert.Equal(t, e.PatientReference, PatientReference("dan"))
}

func TestMarkCovidPassAsEligible(t *testing.T) {
	id := uuid.New()
	// Given Created Covid pass
	es := []cqrs.Event{{AggregateId: cqrs.AggregateId(id.String()), Version: 1, Timestamp: time.Now(), Data: &Created{Eligible: false, PatientReference: "123123-1323"}}}

	p := CovidPass{}
	p.BuildFromHistory(&p, es)

	// When Vaccinate Patient

	p.MarkAsEligible()

	// Then Patient First Vaccine Taken
	assert.Equal(t, len(p.UncommittedEvents), 1)
	assert.Equal(t, reflect.TypeOf(p.UncommittedEvents[0].Data).String(), "*covid.Eligible")
}
