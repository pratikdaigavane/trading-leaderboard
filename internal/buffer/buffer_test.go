package buffer

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"leaderboard/models"
	"math/rand/v2"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestBufferAdd(t *testing.T) {
	t.Parallel()
	name := "bitcoin"
	buf := newFakeBuffer("buff-1", t)
	t.Run("Add", func(t *testing.T) {
		err := buf.Add(getRandomTrade(name))
		assert.Nil(t, err)
	})
}

func TestBufferAddFullFlush(t *testing.T) {
	t.Parallel()
	name := "bitcoin"
	buf := newFakeBuffer("buff-2", t)
	t.Run("Add Buffer Full", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			err := buf.Add(getRandomTrade(name))
			assert.Nil(t, err)
		}
		assert.Zero(t, buf.CurrSize)
	})
}

func TestBufferAddTimeoutFlush(t *testing.T) {
	t.Parallel()
	name := "bitcoin"
	buf := newFakeBuffer("buff-3", t)
	t.Run("Add Timeout Flush", func(t *testing.T) {
		err := buf.Add(getRandomTrade(name))
		assert.Nil(t, err)
		time.Sleep(6 * time.Second)
		assert.Zero(t, buf.CurrSize)
	})
}

func newFakeBuffer(name string, t *testing.T) *Buffer {
	ctx, cancel := context.WithCancel(context.Background())
	funcOnFlush := func(trade []*models.Trade) error {
		return nil
	}
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	buf := NewBuffer(ctx, &logger, name, 10, 5*time.Second, funcOnFlush)
	t.Cleanup(func() {
		cancel()
	})
	return buf
}

func getRandomTrade(symbol string) *models.Trade {
	traderId := uuid.Must(uuid.NewRandom()).String()
	id := uuid.Must(uuid.NewRandom()).String()
	quantity, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64()*1000+0.1), 64)
	return &models.Trade{
		ID:        id,
		Timestamp: time.Now(),
		TraderId:  traderId,
		Symbol:    symbol,
		Quantity:  quantity,
	}
}
