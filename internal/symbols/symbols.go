package symbols

import (
	"github.com/rs/zerolog"
	"leaderboard/internal/config"
	"leaderboard/models"
)

type Manager struct {
	logger *zerolog.Logger
	config *config.Manager
}

func New(log *zerolog.Logger, config *config.Manager) *Manager {
	return &Manager{
		logger: log,
		config: config,
	}
}

// GetSymbols returns the symbols from the config
func (m *Manager) GetSymbols() *map[string]models.Symbol {
	return &m.config.Get().Symbols
}
