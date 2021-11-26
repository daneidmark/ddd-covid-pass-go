package inmemory

import (
	"sync"

	covid "github.com/daneidmark/ddd-covid-pass-go"
)

type IssuingSagaRepository struct {
	DB map[covid.PatientReference]covid.CovidPassId
	mx sync.RWMutex
}

func (r *IssuingSagaRepository) Store(pn covid.PatientReference, id covid.CovidPassId) {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.DB[pn] = id
}

func (r *IssuingSagaRepository) Find(pn covid.PatientReference) covid.CovidPassId {
	r.mx.RLock()
	defer r.mx.RUnlock()
	return r.DB[pn]
}
