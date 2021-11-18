package eventbus

import (
	"fmt"

	"github.com/daneidmark/ddd-covid-pass-go/cqrs"
)

type Service interface {
	Publish(e cqrs.Event)
}

type NoopService struct {

}

func (* NoopService) Publish(e cqrs.Event) {
	fmt.Printf("Publishing %+v\n", e)
}