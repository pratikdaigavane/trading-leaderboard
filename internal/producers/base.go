package producers

import (
	"context"
	"github.com/rs/zerolog"
	"leaderboard/internal/symbols"
	"leaderboard/models"
)

// Producer is used to produce the events and send them to the broker
type Producer struct {
	ctx       context.Context
	broker    chan models.Event
	logger    *zerolog.Logger
	symbols   *symbols.Manager
	producers map[string]EventProducer
}

// EventProducer is an interface that is used to produce the events, this can be implemented by any struct that wants to produce the events
// For example, TradeProducer is a struct that implements the EventProducer interface and produces the trade events
// Composition over inheritance, as always :)
type EventProducer interface {
	Produce() bool
	Start()
}

// New creates a new Producer and registers the event producers that can produce the events
func New(ctx context.Context, broker chan models.Event, logger *zerolog.Logger, symbols *symbols.Manager) *Producer {
	base := &Producer{ctx: ctx, broker: broker, logger: logger, symbols: symbols}
	base.registerEventProducers()
	return base
}

// Start starts the event producers in a separate goroutine and starts producing the events
func (p *Producer) Start() {
	p.logger.Info().Msg("Starting Event Producers")
	for _, producer := range p.producers {
		producer.Start()
	}
}

// registerEventProducers registers the event producers that can produce the events, for example, trade events
func (p *Producer) registerEventProducers() {
	p.producers = make(map[string]EventProducer)
	p.producers["trade"] = &TradeProducer{p}
}

// Produce sends the event to the broker channel to be consumed by the consumer
func (p *Producer) Produce(event models.Event) {
	p.logger.Info().Interface("trade", event).Msg("Producing Event")
	p.broker <- event
}
