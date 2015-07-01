package service
import (
	"github.com/kykl/goapi/models"
	"encoding/json"
	"fmt"
)

type PrintLogger struct {}

func (this *PrintLogger) Log(event models.Event) (id string, err error) {
	rawBytes, err := json.Marshal(event)
	if err != nil {
		return "", err
	}
	fmt.Println(string(rawBytes))
	return event.Id, nil
}
