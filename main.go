package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/covid-pass/patient/register", &PatientApi{PatientReader: &InmemoryPatientReader{}, PatientStorer: &InmemoryPatientStorer{}}).Methods("POST")

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
