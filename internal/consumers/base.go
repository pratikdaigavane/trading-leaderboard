package consumers

import (
	"context"
	"github.com/rs/zerolog"
	"leaderboard/internal/storage"
	"leaderboard/models"
)

type Consumer struct {
	broker  chan models.Event
	logger  *zerolog.Logger
	ctx     context.Context
	events  map[string]EventConsumer
	storage *storage.Storage
}

type EventConsumer interface {
	Consume(interface{}) bool
}

func New(ctx context.Context, broker chan models.Event, logger *zerolog.Logger, storage *storage.Storage) *Consumer {
	c := &Consumer{broker: broker, logger: logger, ctx: ctx, storage: storage}
	c.registerEvents()
	return c
}

func (c *Consumer) registerEvents() {
	c.events = make(map[string]EventConsumer)
	c.events["trade"] = &TradeConsumer{c}
}

func (c *Consumer) Start() {
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				c.logger.Info().Msg("Consumer Stopped")
				return
			case event := <-c.broker:
				c.logger.Info().Interface("event", event).Msg("Event Consumed")
				if val, ok := c.events[event.Type]; ok {
					// pool go routines to handle the event https://github.com/panjf2000/ants
					val.Consume(event.Payload)
				} else {
					c.logger.Warn().Str("event_type", event.Type).Interface("event", event).Msg("Unknown Event Type")
				}
			}
		}
	}()
	c.logger.Info().Msg("Consumer started")
}
