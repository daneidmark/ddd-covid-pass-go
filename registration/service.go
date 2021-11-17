package registration

import (
	covid "github.com/daneidmark/ddd-covid-pass-go"
)

type Service interface {
	RegisterPatient(pn covid.PersonalNumber)
}

type service struct {
	r covid.PatientRepository
}

func NewService(r covid.PatientRepository) *service {
	return &service{r: r}
}

func (s *service) RegisterPatient(pn covid.PersonalNumber) {
	patient := covid.Patient{}
	patient.Register(pn)
	s.r.Store(patient)
}
