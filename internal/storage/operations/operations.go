package operations

import "leaderboard/models"

type Operations interface {
	GetSortedList(symbol string, depth int) (*[]models.UserTradeStat, error)
	BatchAdd(trades []*models.Trade) error
}
