package operations

import "leaderboard/models"

// Operations is an interface that is used to interact with the storage layer.
// It's currently implemented by RedisOperations that interacts with Redis
type Operations interface {
	GetSortedList(symbol string, depth int) (*[]models.UserTradeStat, error)
	BatchAdd(trades []*models.Trade) error
	FlushDB() error
}
