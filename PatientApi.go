package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PatientApi struct {
	PatientReader PatientReader
	PatientStorer PatientStorer
}

type RegisterPatientCommand struct {
	PersonalNumber string `json:"personal_number"	`
}

func (api *PatientApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var cmd RegisterPatientCommand
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Missing body")
	}

	json.Unmarshal(reqBody, &cmd)
	api.handle(cmd)
}

func (api *PatientApi) handle(cmd RegisterPatientCommand) {
	patient := Patient{}
	patient.register(PersonalNumber(cmd.PersonalNumber))
	api.PatientStorer.Store(patient)
}
