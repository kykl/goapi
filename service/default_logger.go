package service

import (
	"github.com/kykl/goapi/models"
	"encoding/json"
	"os"
	"fmt"
	"log"
)

type PrintLogger struct {
}

func (this *PrintLogger) Log(event models.Event) (id string, err error) {
	rawBytes, err := json.Marshal(event)
	if err != nil {
		return "", err
	}
	fmt.Println(string(rawBytes))
	return event.Id, nil
}

type FileLogger struct {
	name string
}

func NewFileLogger(filename string) Logger {
	flog := &FileLogger{}
	flog.name = filename
	return flog
}

func (this *FileLogger) Log(event models.Event) (id string, err error) {
	f, err := os.OpenFile(this.name, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	filelog := log.New(f, "", 0)
	filelog.Printf(event.String())
	return event.Id, nil
}
