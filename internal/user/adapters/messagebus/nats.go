package messagebus

import "encoding/json"

type (

	// EventSender is an adapter that sends an event for user service.
	EventSender struct {
		nc natsClient
	}

	// natsClient is a client for the NATS messaging system.
	// It is used to decouple the user service from the NATS messaging system.
	// natsClient must be low-level implementation of the nats client, without
	// any business logic, or dependencies on other packages.
	natsClient interface {
		Publish(subject string, body []byte) error
	}
)

// NewEventSender creates a new EventSender.
func NewEventSender(nc natsClient) *EventSender {
	return &EventSender{nc: nc}
}

// Send sends an event.
func (es *EventSender) PublishEvent(subject string, events ...interface{}) error {
	for _, event := range events {
		body, err := json.Marshal(event)
		if err != nil {
			return err
		}
		if err := es.nc.Publish(subject, body); err != nil {
			return err
		}
	}
	return nil
}
