package inmemory

import (
	"testing"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/eventbus"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCanSaveCovidPasses(t *testing.T) {
	id := covid.CovidPassId(uuid.New())
	p := covid.NewCovidPass(id, covid.PatientReference("dan"))
	eb := eventbus.NoopService{}
	repo := NewCovidPassRepository(&eb)

	repo.Store(p)

	p1 := repo.Find(id)
	assert.Equal(t, p1.PatientReference, p.PatientReference)
}
