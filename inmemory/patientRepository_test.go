package inmemory

import (
	"testing"

	covid "github.com/daneidmark/ddd-covid-pass-go"
)

func TestCanSavePatients(t *testing.T) {
	p := covid.NewPatient(covid.PersonalNumber("123123-3232"))
	repo := NewPatientRepository()

	repo.Store(p)

	p1 := repo.Find(p.PersonalNumber)
	if p1.PersonalNumber != p.PersonalNumber {
		t.Fatal("The storage is not correct") //TODO: Is there a better way
	}
}
