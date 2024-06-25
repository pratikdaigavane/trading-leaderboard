package producers

import (
	"fmt"
	petname "github.com/dustinkirkland/golang-petname"
	"github.com/google/uuid"
	"leaderboard/models"
	"math/rand/v2"
	"strconv"
	"time"
)

type TradeProducer struct {
	base *Producer
}

// Produce produces the trade events and sends them to the broker
// To simulate the trade events, we generate random trade events with registered symbols, random traders, and  random quantities
func (t *TradeProducer) Produce() bool {
	symbols := t.base.symbols.GetSymbols()
	symbolIds := make([]string, 0, len(*symbols))
	for key, _ := range *symbols {
		symbolIds = append(symbolIds, key)
	}
	traders := make([]string, 0, 100)
	for i := 0; i < 100; i++ {
		traders = append(traders, petname.Generate(2, "-"))
	}
	for i := 0; i < 0; i++ {
		symbolIndex := rand.N(len(symbolIds))
		quantity, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64()*1000+0.1), 64)
		now := time.Now()
		id := uuid.Must(uuid.NewRandom()).String()
		select {
		// If the context is done (shutdown signal received), stop generating the trade events
		case <-t.base.ctx.Done():
			t.base.logger.Info().Msg("Producer Stopped")
			return true
		default:
			trade := models.Trade{
				Symbol:    symbolIds[symbolIndex],
				Quantity:  quantity,
				TraderId:  traders[rand.N(len(traders))],
				ID:        id,
				Timestamp: now,
			}
			event := models.Event{ID: id, Timestamp: now, Payload: trade, Type: "trade"}
			go t.base.Produce(event)
			time.Sleep(1000 * time.Millisecond)
		}
	}
	return true
}

// Start starts the trade producer in a separate goroutine and start producing the trade events
func (t *TradeProducer) Start() {
	go func() {
		t.Produce()
	}()
}
