package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/vaccination"
)

type vaccinationHandler struct {
	s vaccination.Service
}

func NewVaccinationHandler(s vaccination.Service) *vaccinationHandler {
	return &vaccinationHandler{s: s}
}

type VaccinatatePatientCommand struct {
	PersonalNumber string `json:"personal_number"`
	VaccineType    string `json:"vaccine_type`
}

func (h *vaccinationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var cmd VaccinatatePatientCommand
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Missing body")
	}

	json.Unmarshal(reqBody, &cmd)
	h.s.Vaccinate(covid.PersonalNumber(cmd.PersonalNumber), covid.Vaccine{VaccineType: covid.VaccineType(cmd.VaccineType), TimeTaken: time.Now()})
}
