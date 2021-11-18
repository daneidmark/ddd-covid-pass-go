package vaccination

import (
	"testing"
	"time"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/inmemory"
)

// Implement test Aggregate

// tests

func TestRegisterPatient(t *testing.T) {

	r := inmemory.NewPatientRepository()
	s := NewService(r)

	s.RegisterPatient("123123-1233")

	p := r.Find("123123-1233")

	if p.PersonalNumber != "123123-1233" {
		t.Fatal("The personal number is not correct") //TODO: Is there a better way
	}
}

func TestFirstVaccinationOfPatient(t *testing.T) {

	r := inmemory.NewPatientRepository()
	s := NewService(r)

	s.RegisterPatient("123123-1233")
	s.Vaccinate("123123-1233", covid.Vaccine{VaccineType: "Moderna", TimeTaken: time.Now()})

	p := r.Find("123123-1233")

	if p.FirstVaccine.VaccineType != "Moderna" {
		t.Fatal("The VaccineType is not correct") //TODO: Is there a better way
	}
}

func TestSecondVaccinationOfPatient(t *testing.T) {

	r := inmemory.NewPatientRepository()
	s := NewService(r)

	s.RegisterPatient("123123-1233")
	s.Vaccinate("123123-1233", covid.Vaccine{VaccineType: "Moderna", TimeTaken: time.Now()})
	s.Vaccinate("123123-1233", covid.Vaccine{VaccineType: "Moderna", TimeTaken: time.Now()})

	p := r.Find("123123-1233")

	// Then Patient is Second time Vaccinated

	if p.SecondVaccine.VaccineType != "Moderna" {
		t.Fatal("The VaccineType is not correct") //TODO: Is there a better way
	}
}
