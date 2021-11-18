package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/vaccination"
)

type registrationHandler struct {
	s vaccination.Service
}

func NewRegistrationHandler(s vaccination.Service) *registrationHandler {
	return &registrationHandler{s: s}
}

type RegisterPatientCommand struct {
	PersonalNumber string `json:"personal_number"`
}

func (h *registrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var cmd RegisterPatientCommand
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Missing body")
	}

	json.Unmarshal(reqBody, &cmd)
	h.s.RegisterPatient(covid.PersonalNumber(cmd.PersonalNumber))
}
