package service

import (
	"github.com/kykl/goapi/models"
	"encoding/json"
	"encoding/base64"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	_ "google.golang.org/cloud/compute/metadata"
	"google.golang.org/cloud/pubsub"
	"net/http"
	"io/ioutil"
	"errors"
	"log"
	"fmt"
)

type GooglePubSubLogger struct {
	ctx context.Context
}

var jsonFile = "/Users/kykl/Downloads/goapi-464d98feb1e7.json"
var projectId = "goapi-991"

func NewContext() context.Context {
	client, err := NewClient(jsonFile)
	if err != nil {
		log.Fatalf("clientAndId failed, %v", err)
	}
	return cloud.NewContext(projectId, client)
}

func (this *GooglePubSubLogger) Context() (context.Context, error) {
	if this.ctx == nil {
		fmt.Printf("created new context")
		client, err := NewClient(jsonFile)
		if err != nil {
			log.Fatalf("clientAndId failed, %v", err)
			return nil, err
		}
		this.ctx = cloud.NewContext(projectId, client)
	}
	return this.ctx, nil
}

func (this *GooglePubSubLogger) Log(event models.Event) (id string, err error) {
	data, err := json.Marshal(event)
	if err != nil {
		return "", err
	}
	message := base64.StdEncoding.EncodeToString(data)

	fmt.Printf("message: %s\n", message)
	msgIds, err := pubsub.Publish(NewContext(), event.Type, &pubsub.Message{
		Data: []byte(message),
	})

	if err != nil {
		return "", err
	}

	return msgIds[0], nil
}

func(this *GooglePubSubLogger) publish(ctx context.Context, topic string, message string) (id string, err error) {
	msgIds, err := pubsub.Publish(ctx, topic, &pubsub.Message{
		Data: []byte(message),
	})
	if err != nil {
		log.Fatalf("Publish failed, %v", err)
		return "", err
	}

	return msgIds[0], nil
}

// newClient creates http.Client with a jwt service account when
// jsonFile flag is specified, otherwise by obtaining the GCE service
// account's access token.
func NewClient(jsonFile string) (*http.Client, error) {
	if jsonFile != "" {
		jsonKey, err := ioutil.ReadFile(jsonFile)
		if err != nil {
			return nil, err
		}
		conf, err := google.JWTConfigFromJSON(jsonKey, pubsub.ScopePubSub)
		if err != nil {
			return nil, err
		}
		return conf.Client(oauth2.NoContext), nil
	}

	return nil, errors.New("Could not create an authenticated client.")
}