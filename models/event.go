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
	Name string
	CreatedAt time.Time
	CollectedAt time.Time
	UserId string
	SessionId string
	Platform string
	Language string
	AppId string
	Properties map[string]interface{}
}

func (this *Event) String() string {
	return string(this.Bytes())
}

func (this *Event) Bytes() []byte {
	props := map[string]interface{}{
		"id":this.Id,
		"name":this.Name,
		"createdAt":this.CollectedAt,
		"collectedAt":this.CollectedAt,
		"userId":this.UserId,
		"sessionId":this.SessionId,
		"platform":this.Platform,
		"language":this.Language,
		"appId":this.AppId,
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
		name := "log"
		createdAt := time.Now()
		collectedAt := time.Now()
		userId := fmt.Sprintf("userId-%d", i)
		sessionId := fmt.Sprintf("sessionId-%d", i)
		platform := fmt.Sprintf("platform-%d", i)
		language := fmt.Sprintf("language-%d", i)
		appId := fmt.Sprintf("appId-%d", i)

		properties := map[string]interface{} {
			"int" : i,
			"string" : fmt.Sprintf("string-%3d", i),
			"time" : time.Now(),
		}

		Events[id] = &Event{id, name, createdAt, collectedAt, userId, sessionId, platform, language, appId, properties}
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

