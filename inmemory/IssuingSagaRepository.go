package inmemory

import covid "github.com/daneidmark/ddd-covid-pass-go"

type IssuingSagaRepository struct {
	DB map[covid.PatientReference]covid.CovidPassId
}

func (r *IssuingSagaRepository) Store(pn covid.PatientReference, id covid.CovidPassId){
	r.DB[pn] = id
}

func (r *IssuingSagaRepository) Find(pn covid.PatientReference) covid.CovidPassId {
	return r.DB[pn]
}