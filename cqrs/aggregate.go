package cqrs

import "time"

type Version uint64
type AggregateId string
type AggregateRoot struct {
	Id                AggregateId
	Version           Version
	Timestamp         time.Time
	UncommittedEvents []Event
}

type Event struct {
	AggregateId AggregateId
	Version     Version
	Timestamp   time.Time
	Data        interface{}
}

type Aggregate interface {
	Transition(e Event)
}

func (ar *AggregateRoot) Apply(a Aggregate, e Event) {
	a.Transition(e)
	ar.Id = e.AggregateId
	ar.Version = e.Version
	ar.Timestamp = e.Timestamp
}

func (ar *AggregateRoot) ApplyNew(a Aggregate, Data interface{}) {
	e := Event{AggregateId: ar.Id, Version: ar.Version + 1, Timestamp: time.Now(), Data: Data}
	ar.UncommittedEvents = append(ar.UncommittedEvents, e)
	ar.Apply(a, e)
}

func (ar *AggregateRoot) BuildFromHistory(a Aggregate, events []Event) {
	for _, event := range events {
		ar.Apply(a, event)
	}
}
