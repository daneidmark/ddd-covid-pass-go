package inmemory

import (
	"testing"
	"time"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
	"github.com/daneidmark/ddd-covid-pass-go/eventbus"
	"github.com/daneidmark/ddd-covid-pass-go/issuing"
)

func TestCanSavePatients(t *testing.T) {
	p := covid.NewPatient(covid.PersonalNumber("123123-3232"))
	eb := eventbus.InMemEventBus{Subscribers: map[eventbus.Topic][]eventbus.EventHandler{}}
	repo := NewPatientRepository(&eb)

	i := issuing.IssuingEventHandler{Eh: make(chan cqrs.Event)}
	i.Register(&eb)
	go func() {
		i.Consume()
	}()

	repo.Store(p)

	time.Sleep(1000000)
	p1 := repo.Find(p.PersonalNumber)
	if p1.PersonalNumber != p.PersonalNumber {
		t.Fatal("The storage is not correct") //TODO: Is there a better way
	}
}
