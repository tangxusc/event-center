package receiver

import (
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go/pkg/cloudevents"
	"knative.dev/eventing-contrib/pkg/kncloudevents"
	"testing"
)

func TestReceiver(t *testing.T) {
	defaultClient, err := kncloudevents.NewDefaultClient("http://localhost:8080")
	if err != nil {
		panic(err.Error())
	}
	target := cloudevents.New(cloudevents.CloudEventsVersionV03)
	target.SetType("test")
	target.SetSource("https://knative.dev/eventing-contrib/cmd/heartbeats/")
	target.SetID("eventId")
	target.SetExtension("StreamId", "User@"+"1234")
	_ = target.SetData("test")

	send, event, err := defaultClient.Send(context.TODO(), target)
	fmt.Println(send, event, err)
}
