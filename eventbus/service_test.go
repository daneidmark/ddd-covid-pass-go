package eventbus

import (
	"fmt"
	"testing"
	"time"

	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
)

type Event1 struct {
	hej string
}

func TestCanSubscribeToTopic(t *testing.T) {
	e := &InMemEventBus{Subscribers: map[Topic][]EventHandler{}}
	ch1 := make(chan cqrs.Event)
	ch2 := make(chan cqrs.Event)
	ch3 := make(chan cqrs.Event)

	e.Subscribe("topic1", ch1)
	e.Subscribe("topic2", ch2)
	e.Subscribe("topic2", ch3)

	go e.Publish("topic1", cqrs.Event{AggregateId: "231233-1233", Version: 0, Timestamp: time.Now(), Data: &Event1{hej: "dan"}})
	go e.Publish("topic2", cqrs.Event{AggregateId: "231233-1233", Version: 0, Timestamp: time.Now(), Data: &Event1{hej: "Kalle"}})

	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)
		}
	}

}

func printDataEvent(ch string, data cqrs.Event) {
	fmt.Printf("Channel: %s; Envelope: %v; DataEvent: %v\n", ch, data, data.Data)
}
