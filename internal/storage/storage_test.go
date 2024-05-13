package storage

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"leaderboard/internal/config"
	"leaderboard/models"
	"math/rand/v2"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestBatchAdd(t *testing.T) {
	symbol := "bitcoin"
	storage := newFakeStorage(t)
	t.Run("Batch Add", func(t *testing.T) {
		trades := make([]*models.Trade, 10)
		for i := 0; i < 10; i++ {
			trades[i] = getRandomTrade(symbol)
		}
		err := storage.Operations.BatchAdd(trades)
		assert.Nil(t, err)
	})
}

func TestGetSortedList(t *testing.T) {
	symbol := "bitcoin"
	storage := newFakeStorage(t)
	t.Run("Get Sorted List", func(t *testing.T) {
		var trades []*models.Trade
		trades = append(trades, &models.Trade{ID: "1", Timestamp: time.Now(), TraderId: "1", Symbol: symbol, Quantity: 100})
		trades = append(trades, &models.Trade{ID: "2", Timestamp: time.Now(), TraderId: "2", Symbol: symbol, Quantity: 100})
		trades = append(trades, &models.Trade{ID: "3", Timestamp: time.Now(), TraderId: "1", Symbol: symbol, Quantity: 100})
		trades = append(trades, &models.Trade{ID: "4", Timestamp: time.Now(), TraderId: "3", Symbol: symbol, Quantity: 100})
		trades = append(trades, &models.Trade{ID: "5", Timestamp: time.Now(), TraderId: "3", Symbol: symbol, Quantity: 100})
		trades = append(trades, &models.Trade{ID: "6", Timestamp: time.Now(), TraderId: "4", Symbol: symbol, Quantity: 100})
		trades = append(trades, &models.Trade{ID: "7", Timestamp: time.Now(), TraderId: "5", Symbol: symbol, Quantity: 100})
		trades = append(trades, &models.Trade{ID: "8", Timestamp: time.Now(), TraderId: "6", Symbol: symbol, Quantity: 100})
		trades = append(trades, &models.Trade{ID: "9", Timestamp: time.Now(), TraderId: "7", Symbol: symbol, Quantity: 100})
		trades = append(trades, &models.Trade{ID: "10", Timestamp: time.Now(), TraderId: "8", Symbol: symbol, Quantity: 100})
		err := storage.Operations.BatchAdd(trades)
		assert.Nil(t, err)
		lb, err := storage.Operations.GetSortedList(symbol, 10)
		assert.Nil(t, err)
		assert.Equal(t, 8, len(*lb))
		assert.Equal(t, "3", (*lb)[0].TraderId)
		assert.Equal(t, "1", (*lb)[1].TraderId)
		assert.Equal(t, float64(200), (*lb)[0].TotalVolume)
		assert.Equal(t, float64(200), (*lb)[1].TotalVolume)
	})
}

func newFakeStorage(t *testing.T) *Storage {
	cwd, err := os.Getwd()
	t.Log(err)
	t.Log(cwd)
	err = godotenv.Load("../../.env") // Load the .env file from the root directory
	if err != nil {
		t.Fatal("Error loading .env file", err)
	}
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	cfg := config.New(&logger, "../../config.yaml")

	// Set the RedisDB to 8 for testing purposes, so that we don't accidentally flush the production DB
	// In a real-world scenario, ideally we can have a separate Redis instance for testing
	cfg.GetServiceConfig().RedisDB = 8
	storage := New(&logger, cfg)
	_ = storage.Operations.FlushDB() // Flush the test DB before running the tests
	t.Cleanup(func() {
		err := storage.Operations.FlushDB()
		if err != nil {
			return
		}
	})
	return storage
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
