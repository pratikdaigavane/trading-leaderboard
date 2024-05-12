package storage

import (
	"github.com/rs/zerolog"
	"leaderboard/internal/config"
	"leaderboard/internal/storage/operations"
)

type Storage struct {
	logger     *zerolog.Logger
	Operations operations.Operations
}

func New(log *zerolog.Logger, config *config.Manager) *Storage {
	log.Info().Str("Redis Address", config.GetServiceConfig().RedisAddr).Msg("Creating new redis client")
	return &Storage{
		logger:     log,
		Operations: operations.NewRedisOperations(config, log),
	}
}
