package module

import (
	"github.com/kykl/goapi/service"
	"github.com/astaxie/beego"
)

type Module interface {
	Services() []interface{}
}

var DefaultServices = func() []interface{} {
	return []interface{}{
		&service.PrintLogger{},
	}
}


var DevServices = DefaultServices

var TestingServices = DefaultServices

var ProdServices = func() []interface{} {
	return []interface{}{
		//service.NewGooglePubSubLogger(),
		service.NewFileLogger("/var/log/goapi/event.log"),
	}
}
var Services func() []interface{}

func init() {
	if beego.RunMode == "prod" {
		Services = ProdServices
	} else if beego.RunMode == "test" {
		Services = TestingServices
	} else if beego.RunMode == "dev" {
		//Services = DevServices
		Services = ProdServices
	} else {
		Services = DevServices
	}
}