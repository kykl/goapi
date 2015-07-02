package models

import (
	"fmt"
	"errors"
	"time"
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
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

func (this *Event) String() string {
	return string(this.Bytes())
}

func (this *Event) Bytes() []byte {
	props := map[string]interface{}{
		"id":this.Id,
		"type":this.Type,
		"createdAt":this.CollectedAt,
		"collectedAt":this.CollectedAt,
	}
	if len(this.Properties) > 0 {
		js, err := json.Marshal(this.Properties)
		if err == nil {
			props["json"] = string(js)
		}
	}

	js, err := json.Marshal(props)
	if err != nil {
		return  []byte("{}")
	}
	return js
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

func AddOneEvent(event *Event) (EventId string) {
	if event.Id == "" {
		event.Id = uuid.New()
	}
	//event.Id = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Events[event.Id] = event
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

