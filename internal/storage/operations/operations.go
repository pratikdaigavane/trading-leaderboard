package operations

import "leaderboard/models"

type Operations interface {
	Add(trade *models.Trade) error
	GetSortedList(symbol string, depth int) (*[]models.UserTradeStat, error)
}
