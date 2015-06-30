package service

import (
	"fmt"
	"github.com/kykl/goapi/models"
)

type Logger interface {
	Log(event models.Event) (id string, err error)
}

type PrintLogger struct {

}

func (this *PrintLogger) Log(event models.Event) (id string, err error) {
	fmt.Printf("my event: %+v\n", event)
	return event.Id, nil
}