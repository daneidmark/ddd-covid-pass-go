package inmemory

import (
	"errors"
	"sync"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
	"github.com/daneidmark/ddd-covid-pass-go/eventbus"
	"github.com/google/uuid"
)

type covidPassRepository struct {
	db             map[cqrs.AggregateId]eventStorage
	lock           *sync.RWMutex
	eventPublisher eventbus.Service
}

func NewCovidPassRepository(eb eventbus.Service) covid.CovidPassRepository {
	return &covidPassRepository{db: map[cqrs.AggregateId]eventStorage{}, lock: &sync.RWMutex{}, eventPublisher: eb}
}

func (c *covidPassRepository) Store(p covid.CovidPass) error {
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
		c.eventPublisher.Publish(eventbus.Topic("topic"), e)
	}

	c.db[p.Id] = storage

	return nil
}

func (c *covidPassRepository) Find(id covid.CovidPassId) covid.CovidPass {
	c.lock.RLock()
	defer c.lock.RUnlock()
	storage := c.db[cqrs.AggregateId(uuid.UUID(id).String())]
	p := covid.CovidPass{}
	p.BuildFromHistory(&p, storage.Events)

	return p
}
