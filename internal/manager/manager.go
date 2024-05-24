package manager

import (
	"context"
	"github.com/rs/zerolog"
	"leaderboard/internal/buffer"
	"leaderboard/internal/storage"
	"leaderboard/internal/symbols"
	"leaderboard/models"
	"time"
)

// Manager is used to manage the leaderboard and is the core component of the leaderboard system
type Manager struct {
	ctx     context.Context
	buffer  map[string]*buffer.Buffer[models.Trade]
	logger  *zerolog.Logger
	cancel  context.CancelFunc
	storage *storage.Storage
}

// New creates a new Manager and initializes the buffer for each symbol
func New(ctx context.Context, logger *zerolog.Logger, symbols *symbols.Manager, storage *storage.Storage) *Manager {
	// Create a new context and cancel function to stop the buffer when the context is done
	ctx, cancel := context.WithCancel(ctx)
	buf := make(map[string]*buffer.Buffer[models.Trade])
	symbolsList := symbols.GetSymbols()
	for symbol := range *symbolsList {
		logger.Info().Str("symbol", symbol).Msg("Creating Buffer")
		buf[symbol] = buffer.New(ctx, logger, symbol, 10, 5*time.Second, storage.Operations.BatchAdd)
	}
	return &Manager{ctx: ctx, buffer: buf, logger: logger, cancel: cancel, storage: storage}
}

// AddTrade adds the trade to the buffer for the given symbol
func (l *Manager) AddTrade(trade *models.Trade) error {
	err := l.buffer[trade.Symbol].Add(trade)
	if err != nil {
		return err
	}
	return nil
}

// GetLeaderboard returns the leaderboard for the given symbol
func (l *Manager) GetLeaderboard(symbol string, depth int) (*[]models.UserTradeStat, error) {
	lb, err := l.storage.Operations.GetSortedList(symbol, depth)
	if err != nil {
		return nil, err
	}
	return lb, nil
}

// Shutdown stops the LeaderboardManager Manager and cancels the context to stop the buffer
func (l *Manager) Shutdown() {
	l.logger.Info().Msg("Shutting down LeaderboardManager Manager")
	l.cancel()
}
