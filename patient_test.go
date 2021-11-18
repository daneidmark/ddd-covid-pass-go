package covid

import (
	"reflect"
	"testing"
	"time"

	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
)

// Implement test Aggregate

// tests

func TestRegisterPatient(t *testing.T) {

	p := NewPatient("123123-1233")

	if p.PersonalNumber != "123123-1233" {
		t.Fatal("The personal number is not correct") //TODO: Is there a better way
	}

	if len(p.UncommittedEvents) != 1 {
		t.Fatal("No events has been registered")
	}

	if reflect.TypeOf(p.UncommittedEvents[0].Data).String() != "*covid.Registered" {
		t.Fatal("Registered is not the first event " + reflect.TypeOf(p.UncommittedEvents[0].Data).String())
	}
}

func TestFirstVaccinationOfPatient(t *testing.T) {

	// Given Patient Registered
	e := []cqrs.Event{{AggregateId: "123123-1323", Version: 1, Timestamp: time.Now(), Data: &Registered{PersonalNumber: "123123-1323"}}}

	p := Patient{}
	p.BuildFromHistory(&p, e)

	// When Vaccinate Patient

	p.Vaccinate(Vaccine{VaccineType: "Moderna", TimeTaken: time.Now()})

	// Then Patient Vaccinated

	if p.FirstVaccine.VaccineType != "Moderna" {
		t.Fatal("The VaccineType is not correct") //TODO: Is there a better way
	}

	if len(p.UncommittedEvents) != 1 {
		t.Fatal("No events has been registered")
	}

	if reflect.TypeOf(p.UncommittedEvents[0].Data).String() != "*covid.FirstVaccineTaken" {
		t.Fatal("Registered is not the first event " + reflect.TypeOf(p.UncommittedEvents[0].Data).String())
	}
}

func TestSecondVaccinationOfPatient(t *testing.T) {

	// Given Patient Registered and First Vaccinated
	e := []cqrs.Event{
		{AggregateId: "123123-1323", Version: 1, Timestamp: time.Now(), Data: &Registered{PersonalNumber: "123123-1323"}},
		{AggregateId: "123123-1323", Version: 2, Timestamp: time.Now(), Data: &FirstVaccineTaken{VaccineType: "Moderna", TimeTaken: time.Now()}},
	}

	p := Patient{}
	p.BuildFromHistory(&p, e)

	// When Vaccinate Patient

	p.Vaccinate(Vaccine{VaccineType: "Moderna", TimeTaken: time.Now()})

	// Then Patient is Second time Vaccinated

	if p.SecondVaccine.VaccineType != "Moderna" {
		t.Fatal("The VaccineType is not correct") //TODO: Is there a better way
	}

	if len(p.UncommittedEvents) != 1 {
		t.Fatal("No events has been registered")
	}

	if reflect.TypeOf(p.UncommittedEvents[0].Data).String() != "*covid.SecondVaccineTaken" {
		t.Fatal("Registered is not the first event " + reflect.TypeOf(p.UncommittedEvents[0].Data).String())
	}
}
