package issuing

import (
	"fmt"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
	"github.com/google/uuid"
)

type Service struct {
	Repo covid.CovidPassRepository
}

func (s *Service) CreateCovidPass(id covid.CovidPassId, ref covid.PatientReference) {
	fmt.Println("Should create non elibible covid pass")
	c := covid.CovidPass{}
	c.SetId(cqrs.AggregateId(uuid.UUID(id).String()))
	c.New(ref)
	s.Repo.Store(c)
}

func (s *Service) MarkAsEligible(id covid.CovidPassId) {
	fmt.Println("Should create mark covid pass ass elibible on second vaccination")
	c := s.Repo.Find(id)
	c.MarkAsEligible()
	s.Repo.Store(c)
}
