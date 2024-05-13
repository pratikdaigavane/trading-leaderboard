package consumers

import "leaderboard/models"

type TradeConsumer struct {
	base *Consumer
}

// Consume consumes the trade event and adds the trade to the database
func (t *TradeConsumer) Consume(payload interface{}) bool {
	trade := payload.(models.Trade)
	t.base.logger.Info().Interface("trade", trade).Msg("Trade Consumed")
	err := t.base.manager.AddTrade(&trade)
	if err != nil {
		t.base.logger.Error().Err(err).Stack().Msg("Failed to add trade")
		return false
	}
	return true
}
