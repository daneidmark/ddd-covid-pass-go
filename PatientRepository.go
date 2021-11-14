package main

import "fmt"

// Persistance
type PatientReader interface {
	Find(pn PersonalNumber) (p Patient)
}

type PatientStorer interface { //TODO: Ser lite lustigt ut. Effective go
	Store(p Patient)
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
