package main

import (
	"context"
	"github.com/rs/zerolog"
	"leaderboard/internal/consumers"
	"leaderboard/internal/producers"
	"leaderboard/models"
)

type EventManger struct {
	consumer *consumers.Consumer
	producer *producers.Producer
	broker   chan models.Event
	ctx      context.Context
	cancel   context.CancelFunc
	logger   *zerolog.Logger
}

func newEventManager(app *App) *EventManger {
	app.logger.Info().Msg("Initialising Events Manager")
	ctx, cancel := context.WithCancel(app.ctx)
	broker := make(chan models.Event)
	consumer := consumers.New(ctx, broker, app.logger, app.manager)
	producer := producers.New(ctx, broker, app.logger, app.symbols)
	return &EventManger{consumer, producer, broker, ctx, cancel, app.logger}
}

func (m *EventManger) start() {
	m.consumer.Start()
	m.producer.Start()
}

func (m *EventManger) shutdown() {
	m.logger.Info().Msg("Shutting down Events Manager")
	m.cancel()
	close(m.broker)
}
