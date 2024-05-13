package main

import (
	"context"
	"github.com/rs/zerolog"
	"leaderboard/internal/consumers"
	"leaderboard/internal/producers"
	"leaderboard/models"
)

// EventManger is the manager for handling events. It initialises the consumers.Consumer, broker and producers.Producer
// This module is essentially a simulation of event bus system. This is needed to generate mock trade events so that other
// components can consume these events and perform their respective operations. In real world, this would be replaced by
// a real event bus system like Kafka
type EventManger struct {
	consumer *consumers.Consumer
	producer *producers.Producer
	broker   chan models.Event
	ctx      context.Context
	cancel   context.CancelFunc
	logger   *zerolog.Logger
}

// newEventManager creates a new EventManger instance
func newEventManager(app *App) *EventManger {
	app.logger.Info().Msg("Initialising Events Manager")
	// Create a new context from the root context, and a cancel function to cancel the context as and when needed.
	ctx, cancel := context.WithCancel(app.ctx)
	broker := make(chan models.Event)
	consumer := consumers.New(ctx, broker, app.logger, app.manager)
	producer := producers.New(ctx, broker, app.logger, app.symbols)
	return &EventManger{consumer, producer, broker, ctx, cancel, app.logger}
}

// start starts the consumer and producer, both of which are running in separate go routines(virtual threads)
func (m *EventManger) start() {
	m.consumer.Start()
	m.producer.Start()
}

// shutdown cancels the ctx context and closes the broker channel
// the context is used to signal the consumer and producer to stop
func (m *EventManger) shutdown() {
	m.logger.Info().Msg("Shutting down Events Manager")
	m.cancel()
	close(m.broker)
}
