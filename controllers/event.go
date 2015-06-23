package controllers

import (
	"github.com/kykl/goapi/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Event
type EventController struct {
	beego.Controller
}

// @Title create
// @Description create Event
// @Param	body		body 	models.Event	true		"The Event content"
// @Success 200 {string} models.Event.Id
// @Failure 403 body is empty
// @router / [post]
func (o *EventController) Post() {
	var ob models.Event
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	Eventid := models.AddOneEvent(ob)
	o.Data["json"] = map[string]string{"EventId": Eventid}
	o.ServeJson()
}

// @Title Get
// @Description find Event by Eventid
// @Param	EventId		path 	string	true		"the Eventid you want to get"
// @Success 200 {object} models.Event
// @Failure 403 :EventId is empty
// @router /:EventId [get]
func (o *EventController) Get() {
	EventId := o.Ctx.Input.Params[":EventId"]
	if EventId != "" {
		ob, err := models.GetOneEvent(EventId)
		if err != nil {
			o.Data["json"] = err
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJson()
}

// @Title GetAll
// @Description get all Events
// @Success 200 {object} models.Event
// @Failure 403 :EventId is empty
// @router / [get]
func (o *EventController) GetAll() {
	obs := models.GetAllEvents()
	o.Data["json"] = obs
	o.ServeJson()
}
