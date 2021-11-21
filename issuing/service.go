package issuing

import (
	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
	"github.com/google/uuid"
)

type Service struct {
	Repo covid.CovidPassRepository
}

func (s *Service) CreateCovidPass(id covid.CovidPassId, ref covid.PatientReference) {
	c := covid.CovidPass{}
	c.SetId(cqrs.AggregateId(uuid.UUID(id).String()))
	c.New(ref)
	s.Repo.Store(c)
}

func (s *Service) MarkAsEligible(id covid.CovidPassId) {
	c := s.Repo.Find(id)
	c.MarkAsEligible()
	s.Repo.Store(c)
}
