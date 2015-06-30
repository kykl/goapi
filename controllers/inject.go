package controllers

import (
	"github.com/astaxie/beego"
	"github.com/facebookgo/inject"
	"github.com/astaxie/beego/context"
	"github.com/kykl/goapi/module"
)

type InjectableController struct {
	beego.Controller
}

func (this *InjectableController) Init(ct *context.Context, controllerName, actionName string, app interface{}) {
	err := inject.Populate(append(module.Services(), app)...)
	if err != nil {
		panic(err)
	}
	this.Controller.Init(ct, controllerName, actionName, app)
}
