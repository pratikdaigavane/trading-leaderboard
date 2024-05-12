package operations

import (
	"context"
	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/ratelimiter"
	"github.com/rs/zerolog"
	"leaderboard/internal/config"
	"leaderboard/models"
	"time"
)
import "github.com/redis/go-redis/v9"

var ctx = context.Background()

type RedisOperations struct {
	client   *redis.Client
	logger   *zerolog.Logger
	limiters map[string]ratelimiter.RateLimiter[any]
}

func NewRedisOperations(config *config.Manager, logger *zerolog.Logger) *RedisOperations {
	return &RedisOperations{
		client: redis.NewClient(&redis.Options{
			Addr:     config.GetServiceConfig().RedisAddr,
			Password: "",
		}),
		logger:   logger,
		limiters: make(map[string]ratelimiter.RateLimiter[any]),
	}
}

func (r *RedisOperations) Add(trade *models.Trade) error {
	err := r.client.ZIncrBy(ctx, getLeaderboardKey(trade.Symbol), trade.Quantity, trade.TraderId).Err()
	if err != nil {
		r.logger.Error().Err(err).Interface("obj", trade).Stack().Msg("Failed to add trade to redis")
		return err
	}
	err = failsafe.Run(func() error {
		r.expire(trade.Symbol)
		return nil
	}, r.getOrSetLimiter(trade.Symbol))
	r.logger.Info().Interface("obj", trade).Msg("Trade added to Redis")
	return nil
}

func (r *RedisOperations) expire(symbol string) {
	r.logger.Error().Msg("Setting Leaderboard Expiry")
	r.client.ExpireAt(ctx, getLeaderboardKey(symbol), getExpiryTime())
}

func (r *RedisOperations) getOrSetLimiter(symbol string) ratelimiter.RateLimiter[any] {
	if limiter, ok := r.limiters[symbol]; ok {
		return limiter
	}
	r.logger.Info().Str("symbol", symbol).Msg("Creating new rate limiter")
	r.limiters[symbol] = ratelimiter.Bursty[any](1, 1*time.Minute).Build()
	return r.limiters[symbol]
}

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

func getLeaderboardKey(symbol string) string {
	return "leaderboard:" + symbol
}

func getExpiryTime() time.Time {
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).
		AddDate(0, 0, 1)
	return midnight
}
