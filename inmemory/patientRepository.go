package inmemory

import (
	"fmt"
	"sync"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
)

type patientRepository struct {
	db   map[cqrs.AggregateId][]cqrs.Event
	lock *sync.RWMutex
}

func NewPatientRepository() covid.PatientRepository {
	return &patientRepository{}
}

func (c *patientRepository) Store(p covid.Patient) {
	//c.lock.RLock()
	//defer c.lock.RUnlock()

	fmt.Printf("Storing Patient %s", p.PersonalNumber)
	//TODO
}

func (c *patientRepository) Find(id covid.PersonalNumber) covid.Patient {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return covid.Patient{}
}
