package producers

import (
	"context"
	"github.com/rs/zerolog"
	"leaderboard/internal/symbols"
	"leaderboard/models"
)

type Producer struct {
	ctx       context.Context
	broker    chan models.Event
	logger    *zerolog.Logger
	symbols   *symbols.Manager
	producers map[string]EventProducer
}

type EventProducer interface {
	Produce() bool
	Start()
}

func New(ctx context.Context, broker chan models.Event, logger *zerolog.Logger, symbols *symbols.Manager) *Producer {
	base := &Producer{ctx: ctx, broker: broker, logger: logger, symbols: symbols}
	base.registerEventProducers()
	for _, producer := range base.producers {
		producer.Start()
	}
	return base
}

func (p *Producer) registerEventProducers() {
	p.producers = make(map[string]EventProducer)
	p.producers["trade"] = &TradeProducer{p}
}

func (p *Producer) Produce(event models.Event) {
	p.logger.Info().Interface("trade", event).Msg("Producing Event")
	p.broker <- event
}
