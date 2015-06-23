package models

import (
	"fmt"
	"errors"
	"time"
	"code.google.com/p/go-uuid/uuid"
)

var (
	Events map[string]*Event
)

type Event struct {
	Id string
	Type string
	CreatedAt time.Time
	CollectedAt time.Time
	Properties map[string]interface{}
}

func init() {
	Events = make(map[string]*Event)
	for i := 0; i < 10; i++ {
		id := fmt.Sprintf("id%3d", i)
		eventType := "log"
		createdAt := time.Now()
		collectedAt := time.Now()
		properties := map[string]interface{} {
			"int" : i,
			"string" : fmt.Sprintf("string-%3d", i),
			"time" : time.Now(),
		}

		Events[id] = &Event{id, eventType, createdAt, collectedAt, properties}
	}
}

func AddOneEvent(event Event) (EventId string) {
	if event.Id == "" {
		event.Id = uuid.New()
	}
	//event.Id = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Events[event.Id] = &event
	return event.Id
}

func GetOneEvent(id string) (Event *Event, err error) {
	if v, ok := Events[id]; ok {
		return v, nil
	}
	return nil, errors.New("EventId Not Exist")
}

func GetAllEvents() map[string]*Event {
	return Events
}

