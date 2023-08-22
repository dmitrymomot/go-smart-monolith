package events

import (
	"context"
	"log"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/common"
)

// natsClient is a client for the NATS messaging system.
// It is used to publish events.
// See adapters/events/nats.go.
type natsClient interface {
	PublishEvent(subject string, events ...interface{}) error
}

// EventSender is a decoration function that sends an event,
// after the command handler has been executed.
func EventSender[Cmd any](nc natsClient) common.CommandDecorator[Cmd] {
	return func(next common.CommandHandler[Cmd]) common.CommandHandler[Cmd] {
		return func(ctx context.Context, cmd Cmd) ([]interface{}, error) {
			e, err := next(ctx, cmd)
			if err != nil {
				return nil, err
			}

			// Publish the event.
			if len(e) > 0 {
				if err := nc.PublishEvent("events_topic_name", e...); err != nil {
					// log error, but do not return it
					// because the command handler has already been executed
					// and the error has already been returned.
					// This log is for example purposes only.
					log.Printf("error publishing event: %v", err)
				}
			}

			return e, nil
		}
	}
}
