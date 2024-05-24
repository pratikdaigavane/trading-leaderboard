package buffer

import (
	"context"
	"github.com/rs/zerolog"
	"sync"
	"time"
)

// Buffer is used to store the buffer data of any type and funcOnFlush is called when the
// buffer is full or after a certain time interval
type Buffer[T any] struct {
	store       []*T
	CurrSize    int64
	maxSize     int64
	lastFlushed time.Time
	lock        sync.Mutex
	funcOnFlush func(data []*T) error
	maxDuration time.Duration
	logger      *zerolog.Logger
	name        string
}

func New[T any](ctx context.Context, logger *zerolog.Logger, name string, maxCapacity int64, maxDuration time.Duration, funcOnFlush func(data []*T) error) *Buffer[T] {
	buf := &Buffer[T]{
		store:       []*T{},
		name:        name,
		lastFlushed: time.Now(),
		funcOnFlush: funcOnFlush,
		maxSize:     maxCapacity,
		maxDuration: maxDuration,
		logger:      logger,
	}
	go buf.startFlushTicker(ctx)
	return buf
}

// Add adds the value to the buffer and flushes the buffer if the buffer is full or after certain time interval
func (b *Buffer[T]) Add(entry *T) error {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.store = append(b.store, entry)
	b.CurrSize++
	if b.CurrSize >= b.maxSize || time.Since(b.lastFlushed) >= b.maxDuration {
		err := b.flush()
		if err != nil {
			return err
		}
	}
	return nil
}

// flush flushes the buffer, calls the callback and resets the buffer
func (b *Buffer[T]) flush() error {
	b.logger.Info().Str("name", b.name).Interface("store", b.store).Msg("Flushing buffer")
	if len(b.store) > 0 {
		err := b.funcOnFlush(b.store)
		if err != nil {
			b.logger.Error().Str("name", b.name).Interface("store", b.store).Err(err).Msg("Failed to flush buffer")
			return err
		}
	}
	b.CurrSize = 0
	b.store = []*T{}
	b.lastFlushed = time.Now()
	return nil
}

// startFlushTicker starts a ticker to flush the buffer after certain time interval
func (b *Buffer[T]) startFlushTicker(ctx context.Context) {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	for {
		select {
		// If the context is done (shutdown signal received), flush the buffer and return
		case <-ctx.Done():
			b.logger.Info().Str("name", b.name).Msg("Shutting down the buffer ticker")
			// Take the lock and flush so that the buffer is flushed before shutting down
			func() {
				b.lock.Lock()
				defer b.lock.Unlock()
				err := b.flush()
				if err != nil {
					b.logger.Error().Str("name", b.name).Err(err).Msg("Failed to flush buffer")
					return
				}
			}()
			return
		case <-t.C:
			b.logger.Debug().
				Str("name", b.name).
				Int64("size", b.CurrSize).
				Time("lastFlushed", b.lastFlushed).
				Interface("store", b.store).
				Msg("Checking buffer if buffer is full")
			func() {
				b.lock.Lock()
				defer b.lock.Unlock()
				if time.Since(b.lastFlushed) >= b.maxDuration && b.CurrSize > 0 {
					err := b.flush()
					if err != nil {
						b.logger.Error().Str("name", b.name).Err(err).Msg("Failed to flush buffer")
						return
					}
				}
			}()
		}
	}
}
