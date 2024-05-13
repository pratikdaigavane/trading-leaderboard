package operations

import (
	"context"
	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/ratelimiter"
	"github.com/rs/zerolog"
	"leaderboard/internal/config"
	"leaderboard/models"
	"sync"
	"time"
)
import "github.com/redis/go-redis/v9"

var ctx = context.Background()

// RedisOperations is used to interact with the Redis storage layer
type RedisOperations struct {
	client   *redis.Client
	logger   *zerolog.Logger
	limiters map[string]ratelimiter.RateLimiter[any]
	lock     sync.Mutex
}

// NewRedisOperations creates a new RedisOperations and initializes the Redis client
func NewRedisOperations(config *config.Manager, logger *zerolog.Logger) *RedisOperations {
	return &RedisOperations{
		client: redis.NewClient(&redis.Options{
			Addr:     config.GetServiceConfig().RedisAddr,
			Password: "",
			DB:       config.GetServiceConfig().RedisDB,
		}),
		logger:   logger,
		limiters: make(map[string]ratelimiter.RateLimiter[any]),
	}
}

// Add adds the trade to the Redis, it increments the quantity of the trade for the trader in the Redis leaderboard
// A Sorted Set is used to store the leaderboard, where the score is the total quantity of the trade for the trader
// The Sorted Set is set to be expired at midnight every day
// In case the score is same for two traders, the trader id is used to break the tie (lexicographical ordering)
func (r *RedisOperations) Add(pipe redis.Pipeliner, trade *models.Trade) error {
	err := pipe.ZIncrBy(ctx, getLeaderboardKey(trade.Symbol), trade.Quantity, trade.TraderId).Err()
	if err != nil {
		r.logger.Error().Err(err).Interface("obj", trade).Stack().Msg("Failed to add trade to redis")
		return err
	}
	// Limit the number of expiry calls to Redis by using a rate limiter
	// This ensures that the expiry is set at least once a time frame (24 hours in this case) for each symbol
	err = failsafe.Run(func() error {
		r.expire(pipe, trade.Symbol)
		return nil
	}, r.getOrSetLimiter(trade.Symbol))
	return nil
}

// BatchAdd adds the trades to the Redis in batch, it creates a pipeline and adds all the trades to the pipeline
// Hence reducing the number of round trips to the Redis server. The pipeline is wrapped in a transaction (MULTI and EXEC commands).
func (r *RedisOperations) BatchAdd(trades []*models.Trade) error {
	r.logger.Info().Interface("trades", trades).Msg("Adding trades to Redis in batch")
	_, err := r.client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, trade := range trades {
			_ = r.Add(pipe, trade)
		}
		return nil
	})
	if err != nil {
		r.logger.Error().Err(err).Interface("trades", trades).Stack().Msg("Failed to add trade to redis")
		return err
	}
	return nil
}

func (r *RedisOperations) expire(pipe redis.Pipeliner, symbol string) {
	r.logger.Warn().Msg("Setting Leaderboard Expiry")
	pipe.ExpireAt(ctx, getLeaderboardKey(symbol), getExpiryTime())
}

// getOrSetLimiter returns the rate limiter for the given symbol, if it doesn't exist it creates a new rate limiter
func (r *RedisOperations) getOrSetLimiter(symbol string) ratelimiter.RateLimiter[any] {
	if limiter, ok := r.limiters[symbol]; ok {
		return limiter
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	r.logger.Info().Str("symbol", symbol).Msg("Creating new rate limiter")
	// Create a new rate limiter with a burst of 1 and a rate of 1 per minute, the rate can be adjusted as needed
	r.limiters[symbol] = ratelimiter.Bursty[any](1, 1*time.Minute).Build()
	return r.limiters[symbol]
}

// GetSortedList returns the sorted list of traders for the given symbol and depth
func (r *RedisOperations) GetSortedList(symbol string, depth int) (*[]models.UserTradeStat, error) {
	data, err := r.client.ZRevRangeWithScores(ctx, getLeaderboardKey(symbol), 0, int64(depth)-1).Result()
	if err != nil {
		r.logger.Error().Err(err).Str("symbol", symbol).Int("depth", depth).Stack().Msg("Failed to get sorted list from Redis")
		return nil, err
	}
	r.logger.Info().Interface("data", data).Msg("Data from Redis")
	stats := make([]models.UserTradeStat, len(data))
	for i, item := range data {
		stats[i] = models.UserTradeStat{
			TraderId:    item.Member.(string),
			TotalVolume: item.Score,
			Rank:        int64(i + 1),
		}
	}
	return &stats, nil
}

func (r *RedisOperations) FlushDB() error {
	r.logger.Warn().Msg("Flushing Redis DB")
	r.client.FlushDB(ctx)
	return nil
}

// getLeaderboardKey returns the key for the leaderboard in Redis that is used to store the leaderboard for the given symbol
func getLeaderboardKey(symbol string) string {
	return "leaderboard:" + symbol
}

// getExpiryTime returns the expiry time for the leaderboard in Redis, the expiry time is set to midnight every day
func getExpiryTime() time.Time {
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).
		AddDate(0, 0, 1)
	return midnight
}
