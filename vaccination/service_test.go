package vaccination

import (
	"reflect"
	"testing"
	"time"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
	"github.com/daneidmark/ddd-covid-pass-go/eventbus"
	"github.com/daneidmark/ddd-covid-pass-go/inmemory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedEventBus struct {
	mock.Mock
}

func (m *MockedEventBus) Publish(t eventbus.Topic, e cqrs.Event) {
	m.Called(t, e)
}

func TestRegisterPatient(t *testing.T) {
	eb := new(MockedEventBus)
	samePersonalNumber := func(e cqrs.Event) bool {
		assert.Equal(t, reflect.TypeOf(e.Data).String(), "*covid.Registered")
		d := e.Data.(*covid.Registered)
		assert.Equal(t, d.PersonalNumber, covid.PersonalNumber("123123-1233"))
		return true
	}

	eb.On("Publish", eventbus.Topic("topic"), mock.MatchedBy(samePersonalNumber))

	r := inmemory.NewPatientRepository(eb)
	s := NewService(r)

	s.RegisterPatient("123123-1233")

	p := r.Find("123123-1233")
	assert.NotNil(t, p)

	eb.AssertExpectations(t)
}

func TestFirstVaccinationOfPatient(t *testing.T) {
	/*eb := new(MockedEventBus)
	correctVaccineType := func(e cqrs.Event) bool {
		assert.Equal(t, reflect.TypeOf(e.Data).String(), "*covid.FirstVaccineTaken")
		d := e.Data.(*covid.FirstVaccineTaken)
		assert.Equal(t, d.VaccineType, covid.VaccineType("Moderna"))
		return true
	}

	eb.On("Publish", eventbus.Topic("topic"), mock.MatchedBy(correctVaccineType))
	*/

	r := inmemory.NewPatientRepository(&eventbus.NoopService{})
	s := NewService(r)

	s.RegisterPatient("123123-1233")

	s.Vaccinate("123123-1233", covid.Vaccine{VaccineType: "Moderna", TimeTaken: time.Now()})

	p := r.Find("123123-1233")

	if p.FirstVaccine.VaccineType != "Moderna" {
		t.Fatal("The VaccineType is not correct") //TODO: Is there a better way
	}
}

func TestSecondVaccinationOfPatient(t *testing.T) {

	r := inmemory.NewPatientRepository(&eventbus.NoopService{})
	s := NewService(r)

	s.RegisterPatient("123123-1233")
	s.Vaccinate("123123-1233", covid.Vaccine{VaccineType: "Moderna", TimeTaken: time.Now()})
	s.Vaccinate("123123-1233", covid.Vaccine{VaccineType: "Moderna", TimeTaken: time.Now()})

	p := r.Find("123123-1233")

	// Then Patient is Second time Vaccinated

	if p.SecondVaccine.VaccineType != "Moderna" {
		t.Fatal("The VaccineType is not correct") //TODO: Is there a better way
	}
}
