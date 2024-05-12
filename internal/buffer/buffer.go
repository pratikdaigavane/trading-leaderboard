package buffer

import (
	"context"
	"github.com/rs/zerolog"
	"leaderboard/models"
	"sync"
	"time"
)

// Buffer is used to store the buffer data of trade volume and is pushed to the database only when the
// buffer is full or after certain time interval
type Buffer struct {
	store       []*models.Trade
	currSize    int64
	maxSize     int64
	lastFlushed time.Time
	lock        sync.Mutex
	funcOnFlush func(trade []*models.Trade) error
	maxDuration time.Duration
	ctx         context.Context
	logger      *zerolog.Logger
	name        string
}

var maxCapacity int64 = 10

func NewBuffer(ctx context.Context, logger *zerolog.Logger, name string, funcOnFlush func(trade []*models.Trade) error) *Buffer {
	buf := &Buffer{
		ctx:         ctx,
		store:       []*models.Trade{},
		name:        name,
		lastFlushed: time.Now(),
		funcOnFlush: funcOnFlush,
		maxSize:     maxCapacity,
		maxDuration: 5 * time.Second,
		logger:      logger,
	}
	go buf.startFlushTicker()
	return buf
}

// Add adds the value to the buffer and flushes the buffer if the buffer is full or after certain time interval
func (b *Buffer) Add(trade *models.Trade) error {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.store = append(b.store, trade)
	b.currSize++
	if b.currSize >= b.maxSize || time.Since(b.lastFlushed) >= b.maxDuration {
		err := b.flush()
		if err != nil {
			return err
		}
	}
	return nil
}

// flush flushes the buffer, calls the callback and resets the buffer
func (b *Buffer) flush() error {
	b.logger.Info().Str("name", b.name).Interface("store", b.store).Msg("Flushing buffer")
	if len(b.store) > 0 {
		err := b.funcOnFlush(b.store)
		if err != nil {
			b.logger.Error().Str("name", b.name).Interface("store", b.store).Err(err).Msg("Failed to flush buffer")
			return err
		}
	}
	b.currSize = 0
	b.store = []*models.Trade{}
	b.lastFlushed = time.Now()
	return nil
}

// startFlushTicker starts a ticker to flush the buffer after certain time interval
func (b *Buffer) startFlushTicker() {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-b.ctx.Done():
			b.logger.Info().Str("name", b.name).Msg("Shutting down the buffer ticker")
			// Take the lock and flush so that the buffer is flushed before shutting down
			b.lock.Lock()
			b.flush()
			b.lock.Unlock()
			return
		case <-t.C:
			b.logger.Debug().
				Str("name", b.name).
				Int64("size", b.currSize).
				Time("lastFlushed", b.lastFlushed).
				Interface("store", b.store).
				Msg("Checking buffer if buffer is full")
			b.lock.Lock()
			if time.Since(b.lastFlushed) >= b.maxDuration && b.currSize > 0 {
				b.flush()
			}
			b.lock.Unlock()
		}
	}
}
