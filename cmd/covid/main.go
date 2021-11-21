package main

import (
	"fmt"
	"net/http"

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

	i := issuing.IssuingEventHandler{Eh: make(chan cqrs.Event)}
	i.Register(&eb)
	go i.Consume()
	
	fmt.Println("hej")
	router.Handle("/covid-pass/patient/register", server.NewRegistrationHandler(vaccination.NewService(inmemory.NewPatientRepository(&eb))))
	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
