package repository

import (
	"context"
	"github.com/tangxusc/event-center/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//{"agg_type":"",agg_id:"",date_time:"",events:[{event_id:"",data:"",type:""},{event_id:"",data:""}]}
type Event struct {
	AggType   string `json:"agg_type"`
	AggId     string `json:"agg_id"`
	DateTime  string `json:"date_time"`
	EventId   string `json:"event_id"`
	EventType string `json:"event_type"`
	Data      []byte `json:"data"`
}

func NewEvent(aggType string, aggId string, dateTime string, eventId string, eventType string, data []byte) *Event {
	return &Event{AggType: aggType, AggId: aggId, DateTime: dateTime, EventId: eventId, EventType: eventType, Data: data}
}

func (event *Event) Save(ctx context.Context) error {
	upsert := true
	after := options.After
	result := client.Database(config.Instance.Mongo.DbName).Collection(event.AggType).FindOneAndUpdate(ctx,
		bson.M{
			"agg_type":  event.AggType,
			"agg_id":    event.AggId,
			"date_time": event.DateTime,
		},
		bson.M{
			"$push": bson.M{
				"events": bson.M{
					"event_id":   event.EventId,
					"event_type": event.EventType,
					"data":       event.Data,
				},
			},
		},
		&options.FindOneAndUpdateOptions{
			Upsert:         &upsert,
			ReturnDocument: &after,
		})
	return result.Err()
}
