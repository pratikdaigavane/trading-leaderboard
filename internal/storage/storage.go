package storage

import (
	"github.com/rs/zerolog"
	"leaderboard/internal/config"
	"leaderboard/internal/storage/operations"
)

// Storage is used to interact with the storage layer, for example, Redis
// This module is structured in such a way that it can be easily replaced with any other storage system
// Without changing the other modules
type Storage struct {
	logger     *zerolog.Logger
	Operations operations.Operations
}

// New creates a new Storage and initializes the Operations client (Redis in this case)
func New(log *zerolog.Logger, config *config.Manager) *Storage {
	log.Info().Str("Redis Address", config.GetServiceConfig().RedisAddr).Msg("Creating new redis client")
	return &Storage{
		logger:     log,
		Operations: operations.NewRedisOperations(config, log),
	}
}
