package main

import (
	"net/http"

	"github.com/daneidmark/ddd-covid-pass-go/inmemory"
	"github.com/daneidmark/ddd-covid-pass-go/registration"
	"github.com/daneidmark/ddd-covid-pass-go/server"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/covid-pass/patient/register", server.NewRegistrationHandler(registration.NewService(inmemory.NewPatientRepository())))
	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
