package consumers

import (
	"context"
	"github.com/rs/zerolog"
	"leaderboard/internal/manager"
	"leaderboard/models"
)

// Consumer is used to consume the events from the broker and process them
type Consumer struct {
	broker  chan models.Event
	logger  *zerolog.Logger
	ctx     context.Context
	events  map[string]EventConsumer
	manager *manager.Manager
}

// EventConsumer is an interface that is used to consume the events, this can be implemented by any struct that wants to consume the events
// For example, TradeConsumer is a struct that implements the EventConsumer interface and consumes the trade events
type EventConsumer interface {
	Consume(interface{}) bool
}

// New creates a new Consumer and registers the events that can be consumed
func New(ctx context.Context, broker chan models.Event, logger *zerolog.Logger, manager *manager.Manager) *Consumer {
	c := &Consumer{broker: broker, logger: logger, ctx: ctx, manager: manager}
	c.registerEvents()
	return c
}

// registerEvents registers the events that can be consumed by the consumer, for example, trade events
func (c *Consumer) registerEvents() {
	c.events = make(map[string]EventConsumer)
	c.events["trade"] = &TradeConsumer{c}
}

// Start starts the consumer in a separate goroutine and starts consuming the events
func (c *Consumer) Start() {
	go func() {
		for {
			select {
			// If the context is done (shutdown signal received), stop the consumer
			case <-c.ctx.Done():
				c.logger.Info().Msg("Consumer Stopped")
				return
			case event := <-c.broker:
				c.logger.Info().Interface("event", event).Msg("Event Consumed")
				if val, ok := c.events[event.Type]; ok {
					// Consume the event in a separate goroutine
					go val.Consume(event.Payload)
				} else {
					c.logger.Warn().Str("event_type", event.Type).Interface("event", event).Msg("Unknown Event Type")
				}
			}
		}
	}()
	c.logger.Info().Msg("Consumer started")
}
