package main

import (
	"reflect"
	"testing"
)

// Implement test Aggregate
type Name string

type Person struct {
	AggregateRoot
	Name Name
}

func (p *Person) birth(n Name) {
	// Enforce business invariants

	p.ApplyNew(p, &Born{Name: n})
}

func (person *Person) Transition(e Event) {
	switch e := e.Data.(type) {
	case *Born:
		person.Name = e.Name
	}
}

type Born struct {
	Name Name
}

// tests

func TestPersonBorn(t *testing.T) {
	person := Person{}
	person.birth("Pelle")
	if person.Name != "Pelle" {
		t.Fatal("The name is not correct") //TODO: Is there a better way
	}

	if len(person.UncommittedEvents) != 1 {
		t.Fatal("No events has been registered")
	}

	if reflect.TypeOf(person.UncommittedEvents[0].Data).String() != "*main.Born" {
		t.Fatal("Born is not the first event " + reflect.TypeOf(person.UncommittedEvents[0].Data).String())
	}
}
