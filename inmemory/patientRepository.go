package inmemory

import (
	"errors"
	"sync"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
	"github.com/daneidmark/ddd-covid-pass-go/eventbus"
)

type eventStorage struct {
	Id             cqrs.AggregateId
	CurrentVersion cqrs.Version
	Events         []cqrs.Event
}

type patientRepository struct {
	db             map[cqrs.AggregateId]eventStorage
	lock           *sync.RWMutex
	eventPublisher eventbus.Service
}

func NewPatientRepository() covid.PatientRepository {
	return &patientRepository{db: map[cqrs.AggregateId]eventStorage{}, lock: &sync.RWMutex{}, eventPublisher: &eventbus.NoopService{}}
}

func (c *patientRepository) Store(p covid.Patient) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if _, ok := c.db[p.Id]; !ok {
		c.db[p.Id] = eventStorage{Id: p.Id, CurrentVersion: 0, Events: []cqrs.Event{}}
	}

	storage := c.db[p.Id]

	if storage.CurrentVersion >= p.Version {
		return errors.New("wrong version")
	}

	storage.CurrentVersion = p.Version
	storage.Events = append(storage.Events, p.UncommittedEvents...)

	for _, e := range p.UncommittedEvents {
		c.eventPublisher.Publish(e)
	}

	c.db[p.Id] = storage

	return nil
}

func (c *patientRepository) Find(id covid.PersonalNumber) covid.Patient {
	c.lock.RLock()
	defer c.lock.RUnlock()
	storage := c.db[cqrs.AggregateId(id)]
	p := covid.Patient{}
	p.BuildFromHistory(&p, storage.Events)

	return p
}
