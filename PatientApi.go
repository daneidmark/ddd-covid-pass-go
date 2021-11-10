package main

import (
	"fmt"
	"net/http"
)

type PatientApi struct {
	PatientReader PatientReader
	PatientStorer PatientStorer
}

func (api *PatientApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.register("19571232-0066") //TODO: get from request
}

func (api *PatientApi) register(pn PersonalNumber) {
	patient := Patient{}
	patient.register(pn)
	api.PatientStorer.Store(patient)
}

type InmemoryPatientReader struct {
}

func (r *InmemoryPatientReader) Find(pn PersonalNumber) Patient {
	return Patient{PersonalNumber: "dsa"}
}

type InmemoryPatientStorer struct {
}

func (r *InmemoryPatientStorer) Store(p Patient) {
	for _, e := range p.UncommittedEvents {
		fmt.Printf("saving event: %v\n", e)
	}
	fmt.Printf("PersonalNumber: %v\n", p.PersonalNumber)
}
