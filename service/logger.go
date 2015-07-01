package service

import (
	"github.com/kykl/goapi/models"
)

type Logger interface {
	Log(event models.Event) (id string, err error)
}


