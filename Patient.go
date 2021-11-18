package covid

import (
	"errors"
	"time"

	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
)

// Domain entities
type PersonalNumber string
type VaccineType string

type Vaccine struct {
	VaccineType VaccineType
	TimeTaken   time.Time
}

type Patient struct {
	cqrs.AggregateRoot
	PersonalNumber PersonalNumber
	FirstVaccine   Vaccine
	SecondVaccine  Vaccine
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

func (p *Patient) Vaccinate(v Vaccine) error {
	if p.FirstVaccine == (Vaccine{}) {
		p.ApplyNew(p, &FirstVaccineTaken{VaccineType: v.VaccineType, TimeTaken: v.TimeTaken})
	} else if p.SecondVaccine == (Vaccine{}) {
		p.ApplyNew(p, &SecondVaccineTaken{VaccineType: v.VaccineType, TimeTaken: v.TimeTaken})
	} else {
		return errors.New("Patient is already double vaccinated")
	}

	return nil
}

func (patient *Patient) Transition(e cqrs.Event) {
	switch e := e.Data.(type) {
	case *Registered:
		patient.PersonalNumber = e.PersonalNumber
	case *FirstVaccineTaken:
		patient.FirstVaccine = Vaccine{VaccineType: e.VaccineType, TimeTaken: e.TimeTaken}
	case *SecondVaccineTaken:
		patient.SecondVaccine = Vaccine{VaccineType: e.VaccineType, TimeTaken: e.TimeTaken}
	}
}

// Events
type Registered struct {
	PersonalNumber PersonalNumber
}
type FirstVaccineTaken struct {
	VaccineType VaccineType
	TimeTaken   time.Time
}
type SecondVaccineTaken struct {
	VaccineType VaccineType
	TimeTaken   time.Time
}

// Repository
type PatientRepository interface {
	Store(p Patient) error
	Find(pn PersonalNumber) Patient
}
