package registration

import (
	covid "github.com/daneidmark/ddd-covid-pass-go"
)

type Service interface {
	RegisterPatient(pn covid.PersonalNumber)
	Vaccinate(pn covid.PersonalNumber, v covid.Vaccine)
}

type service struct {
	r covid.PatientRepository
}

func NewService(r covid.PatientRepository) Service {
	return &service{r: r}
}

func (s *service) RegisterPatient(pn covid.PersonalNumber) {
	patient := covid.NewPatient(pn)
	s.r.Store(patient)
}

func (s *service) Vaccinate(pn covid.PersonalNumber, v covid.Vaccine) {
	p := s.r.Find(pn)
	p.Vaccinate(v)
	s.r.Store(p)
}
