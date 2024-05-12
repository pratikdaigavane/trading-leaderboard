package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"leaderboard/internal/config"
	"leaderboard/internal/manager"
	"leaderboard/internal/symbols"
)

type handles struct {
	logger  *zerolog.Logger
	config  *config.Manager
	symbols *symbols.Manager
	manager *manager.Manager
}

func SetupHandlers(r *gin.Engine, log *zerolog.Logger, config *config.Manager, symbols *symbols.Manager, manager *manager.Manager) {
	h := &handles{log, config, symbols, manager}
	r.GET("/version", h.handleVersion())
	r.GET("/leaderboard/:symbol", h.getLeaderBoard())
	r.GET("/symbols", h.getSymbols())
}
