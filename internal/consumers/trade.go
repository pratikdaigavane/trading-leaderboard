package consumers

import "leaderboard/models"

type TradeConsumer struct {
	base *Consumer
}

func (t *TradeConsumer) Consume(payload interface{}) bool {
	trade := payload.(models.Trade)
	t.base.logger.Info().Interface("trade", trade).Msg("Trade Consumed")
	err := t.base.storage.Operations.Add(&trade)
	if err != nil {
		t.base.logger.Error().Err(err).Stack().Msg("Failed to add trade")
		return false
	}
	return true
}
