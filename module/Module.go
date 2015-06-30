package module

import (
	"github.com/kykl/goapi/service"
)

type Module interface {
	Services() []interface{}
}

var DefaultServices = func() []interface{} {
	return []interface{}{
		&service.GooglePubSubLogger{},
		service.NewContext(),
	}
}

var DevServices = DefaultServices
var TestingServices = DefaultServices
var ProdServices = DefaultServices

var Services = DevServices
