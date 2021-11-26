package covid

import (
	"reflect"
	"testing"
	"time"

	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
	"github.com/stretchr/testify/assert"
)

// Implement test Aggregate

// tests

func TestRegisterPatient(t *testing.T) {

	p := NewPatient("123123-1233")

	assert.Equal(t, len(p.UncommittedEvents), 1)
	assert.Equal(t, reflect.TypeOf(p.UncommittedEvents[0].Data).String(), "*covid.Registered")

	e := p.UncommittedEvents[0].Data.(*Registered)
	assert.Equal(t, e.PersonalNumber, PersonalNumber("123123-1233"))
}

func TestFirstVaccinationOfPatient(t *testing.T) {

	// Given Patient Registered
	es := []cqrs.Event{{AggregateId: "123123-1323", Version: 1, Timestamp: time.Now(), Data: &Registered{PersonalNumber: "123123-1323"}}}

	p := Patient{}
	p.BuildFromHistory(&p, es)

	// When Vaccinate Patient

	p.Vaccinate(Vaccine{VaccineType: "Moderna", TimeTaken: time.Now()})

	// Then Patient First Vaccine Taken
	assert.Equal(t, len(p.UncommittedEvents), 1)
	assert.Equal(t, reflect.TypeOf(p.UncommittedEvents[0].Data).String(), "*covid.FirstVaccineTaken")
	e := p.UncommittedEvents[0].Data.(*FirstVaccineTaken)
	assert.Equal(t, e.VaccineType, VaccineType("Moderna"))
}

func TestSecondVaccinationOfPatient(t *testing.T) {

	// Given Patient Registered and First Vaccinated
	es := []cqrs.Event{
		{AggregateId: "123123-1323", Version: 1, Timestamp: time.Now(), Data: &Registered{PersonalNumber: "123123-1323"}},
		{AggregateId: "123123-1323", Version: 2, Timestamp: time.Now(), Data: &FirstVaccineTaken{VaccineType: "Moderna", TimeTaken: time.Now()}},
	}

	p := Patient{}
	p.BuildFromHistory(&p, es)

	// When Vaccinate Patient

	p.Vaccinate(Vaccine{VaccineType: "Moderna", TimeTaken: time.Now()})

	// Then Patient is Second time Vaccinated
	assert.Equal(t, len(p.UncommittedEvents), 1)
	assert.Equal(t, reflect.TypeOf(p.UncommittedEvents[0].Data).String(), "*covid.SecondVaccineTaken")
	e := p.UncommittedEvents[0].Data.(*SecondVaccineTaken)
	assert.Equal(t, e.VaccineType, VaccineType("Moderna"))
}

func TestSecondVaccinationWithDifferentTypesIsNotAllowed(t *testing.T) {

	// Given Patient Registered and First Vaccinated
	es := []cqrs.Event{
		{AggregateId: "123123-1323", Version: 1, Timestamp: time.Now(), Data: &Registered{PersonalNumber: "123123-1323"}},
		{AggregateId: "123123-1323", Version: 2, Timestamp: time.Now(), Data: &FirstVaccineTaken{VaccineType: "Moderna", TimeTaken: time.Now()}},
	}

	p := Patient{}
	p.BuildFromHistory(&p, es)

	// When Vaccinate Patient

	err := p.Vaccinate(Vaccine{VaccineType: "NOT_Moderna", TimeTaken: time.Now()})

	// Then it breaks
	assert.NotNil(t, err)
}

func TestThirdVaccinationIsNotAllowed(t *testing.T) {

	// Given Patient Registered and First Vaccinated and Second vaccination
	es := []cqrs.Event{
		{AggregateId: "123123-1323", Version: 1, Timestamp: time.Now(), Data: &Registered{PersonalNumber: "123123-1323"}},
		{AggregateId: "123123-1323", Version: 2, Timestamp: time.Now(), Data: &FirstVaccineTaken{VaccineType: "Moderna", TimeTaken: time.Now()}},
		{AggregateId: "123123-1323", Version: 3, Timestamp: time.Now(), Data: &SecondVaccineTaken{VaccineType: "Moderna", TimeTaken: time.Now()}},
	}

	p := Patient{}
	p.BuildFromHistory(&p, es)

	// When Vaccinate Patient

	err := p.Vaccinate(Vaccine{VaccineType: "Moderna", TimeTaken: time.Now()})

	// Then it breaks
	assert.NotNil(t, err)
}
