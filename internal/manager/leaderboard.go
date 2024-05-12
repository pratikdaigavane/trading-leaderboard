package manager

import (
	"context"
	"github.com/rs/zerolog"
	"leaderboard/internal/buffer"
	"leaderboard/internal/storage"
	"leaderboard/internal/symbols"
	"leaderboard/models"
)

type Manager struct {
	ctx     context.Context
	buffer  map[string]*buffer.Buffer
	logger  *zerolog.Logger
	cancel  context.CancelFunc
	storage *storage.Storage
}

func New(ctx context.Context, logger *zerolog.Logger, symbols *symbols.Manager, storage *storage.Storage) *Manager {
	ctx, cancel := context.WithCancel(ctx)
	buf := make(map[string]*buffer.Buffer)
	symbolsList := symbols.GetSymbols()
	for symbol := range *symbolsList {
		logger.Info().Str("symbol", symbol).Msg("Creating Buffer")
		buf[symbol] = buffer.NewBuffer(ctx, logger, symbol, storage.Operations.BatchAdd)
	}
	return &Manager{ctx: ctx, buffer: buf, logger: logger, cancel: cancel, storage: storage}
}

func (l *Manager) AddTrade(trade *models.Trade) error {
	err := l.buffer[trade.Symbol].Add(trade)
	if err != nil {
		return err
	}
	return nil
}

func (l *Manager) GetLeaderboard(symbol string, depth int) (*[]models.UserTradeStat, error) {
	lb, err := l.storage.Operations.GetSortedList(symbol, depth)
	if err != nil {
		return nil, err
	}
	return lb, nil
}

func (l *Manager) Shutdown() {
	l.logger.Info().Msg("Shutting down LeaderboardManager Manager")
	l.cancel()
}
