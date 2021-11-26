package inmemory

import (
	"testing"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/eventbus"
	"github.com/stretchr/testify/assert"
)

func TestCanSavePatients(t *testing.T) {
	p := covid.NewPatient(covid.PersonalNumber("123123-3232"))
	eb := eventbus.NoopService{}
	repo := NewPatientRepository(&eb)

	repo.Store(p)

	p1 := repo.Find(p.PersonalNumber)
	assert.Equal(t, p1.PersonalNumber, p.PersonalNumber)
}
