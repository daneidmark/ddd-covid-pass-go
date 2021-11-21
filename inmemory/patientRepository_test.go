package inmemory

import (
	"testing"
	"time"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/eventbus"
)

func TestCanSavePatients(t *testing.T) {
	p := covid.NewPatient(covid.PersonalNumber("123123-3232"))
	eb := eventbus.NoopService{}
	repo := NewPatientRepository(&eb)

	repo.Store(p)

	time.Sleep(1000000)
	p1 := repo.Find(p.PersonalNumber)
	if p1.PersonalNumber != p.PersonalNumber {
		t.Fatal("The storage is not correct") //TODO: Is there a better way
	}
}
