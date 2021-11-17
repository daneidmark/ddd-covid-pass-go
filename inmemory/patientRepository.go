package inmemory

import (
	"fmt"
	"sync"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
)

type patientRepository struct {
	db   map[covid.PersonalNumber][]cqrs.Event
	lock *sync.RWMutex
}

func NewPatientRepository() covid.PatientRepository {
	return &patientRepository{db: map[covid.PersonalNumber][]cqrs.Event{}, lock: &sync.RWMutex{}}
}

func (c *patientRepository) Store(p covid.Patient) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	fmt.Printf("Storing Patient %s", p.PersonalNumber)
	fmt.Printf("Storing Patient %+v\n", p.UncommittedEvents)

	//TODO
}

func (c *patientRepository) Find(id covid.PersonalNumber) covid.Patient {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return covid.Patient{}
}
