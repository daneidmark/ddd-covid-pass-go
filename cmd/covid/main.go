package main

import (
	"net/http"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
	"github.com/daneidmark/ddd-covid-pass-go/eventbus"
	"github.com/daneidmark/ddd-covid-pass-go/inmemory"
	"github.com/daneidmark/ddd-covid-pass-go/issuing"
	"github.com/daneidmark/ddd-covid-pass-go/server"
	"github.com/daneidmark/ddd-covid-pass-go/vaccination"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	eb := eventbus.InMemEventBus{Subscribers: map[eventbus.Topic][]eventbus.EventHandler{}}

	i := issuing.Saga{Service: issuing.Service{Repo: inmemory.NewCovidPassRepository(&eventbus.NoopService{})}, Repo: inmemory.IssuingSagaRepository{DB: map[covid.PatientReference]covid.CovidPassId{}}, EventHandler: make(chan cqrs.Event)}
	i.Register(&eb)
	go i.Consume()

	s := vaccination.NewService(inmemory.NewPatientRepository(&eb))
	router.Handle("/covid-pass/patient/register", server.NewRegistrationHandler(s))
	router.Handle("/covid-pass/patient/vaccinate", server.NewVaccinationHandler(s))
	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
