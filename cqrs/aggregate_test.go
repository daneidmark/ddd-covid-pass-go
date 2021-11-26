package cqrs

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, person.Name, Name("Pelle"))
	assert.Equal(t, len(person.UncommittedEvents), 1)
	assert.Equal(t, reflect.TypeOf(person.UncommittedEvents[0].Data).String(), "*cqrs.Born")
}
