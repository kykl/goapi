package service

import (
	"github.com/kykl/goapi/models"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/compute/metadata"
	"google.golang.org/cloud/pubsub"
	"net/http"
	"io/ioutil"
	"errors"
	"log"
	"github.com/astaxie/beego"
)

type GooglePubSubLogger struct {
	ctx context.Context
	topics map[string]bool
}

func NewGooglePubSubLogger() *GooglePubSubLogger {
	if logger == nil {
		logger = &GooglePubSubLogger{}
		logger.topics = map[string]bool{}
		logger.createContext()
	}

	return logger
}

func (this *GooglePubSubLogger) Log(event models.Event) (id string, err error) {
	if !this.topics[event.Type] {
		// lookup and create the topic
		hasTopic, err := pubsub.TopicExists(this.ctx, event.Type)
		if err != nil {
			return "", err
		}
		if !hasTopic {
			if pubsub.CreateTopic(this.ctx, event.Type) != nil {
				return "", err
			}
		}
		this.topics[event.Type] = true
	}

	data := event.Bytes()

	msgIds, err := pubsub.Publish(this.ctx, event.Type, &pubsub.Message{
		Data: data,
	})

	if err != nil {
		return "", err
	}

	return msgIds[0], nil
}

func (this *GooglePubSubLogger) createContext() (context.Context, error) {
	if this.ctx == nil {
		jsonFile := beego.AppConfig.DefaultString("gcp.service.account.json", "goapi.json")
		client, err := newClient(jsonFile)
		if err != nil {
			log.Fatalf("clientAndId failed, %v", err)
			return nil, err
		}
		projectId := beego.AppConfig.DefaultString("gcp.project.id", "goapi-991")
		this.ctx = cloud.NewContext(projectId, client)
	}
	return this.ctx, nil
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
func newClient(jsonFile string) (*http.Client, error) {
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

	if metadata.OnGCE() {
		c := &http.Client{
			Transport: &oauth2.Transport{
				Source: google.ComputeTokenSource(""),
			},
		}
		/*if *projID == "" {
			projectID, err := metadata.ProjectID()
			if err != nil {
				return nil, fmt.Errorf("ProjectID failed, %v", err)
			}
			*projID = projectID
		}*/
		return c, nil
	}
	return nil, errors.New("Could not create an authenticated client.")
}

var logger *GooglePubSubLogger
