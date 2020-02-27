package receiver

import (
	"context"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go"
	cloudevents2 "github.com/cloudevents/sdk-go/pkg/cloudevents"
	"github.com/tangxusc/event-center/pkg/repository"
	"strings"
	"time"
)

const ExtKey string = "StreamId"

func Start(ctx context.Context) error {
	client, err := cloudevents.NewDefaultClient()
	if err != nil {
		return err
	}
	err = client.StartReceiver(ctx, receiver)
	if err != nil {
		return err
	}
	return nil
}

func receiver(ctx context.Context, event cloudevents.Event) error {
	v3 := event.Context.AsV03()
	aggType, aggId, err := getAggInfo(v3)
	if err != nil {
		return err
	}
	timeDuration := getTimeDuration()
	bytes, ok := event.Data.([]byte)
	if !ok {
		return fmt.Errorf("data type not []byte")
	}
	err = repository.NewEvent(aggType, aggId, timeDuration, v3.GetID(), v3.Type, bytes).Save(ctx)
	return err
}

func getTimeDuration() string {
	now := time.Now()
	return now.Format("2006-01-02 15")
}

func getAggInfo(event *cloudevents2.EventContextV03) (aggType string, aggId string, err error) {
	extension, err := event.GetExtension(ExtKey)
	if err != nil {
		return
	}
	streamId := extension.(string)
	split := strings.Split(streamId, "@")
	if len(split) != 2 {
		err = fmt.Errorf("StreamId[%v] split at [@] error,except result len 2", streamId)
	}
	aggType, aggId = split[0], split[1]
	return
}
