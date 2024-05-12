package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"leaderboard/internal/config"
	"leaderboard/internal/storage"
	"leaderboard/internal/symbols"
)

type handles struct {
	logger  *zerolog.Logger
	config  *config.Manager
	storage *storage.Storage
	symbols *symbols.Manager
}

func SetupHandlers(r *gin.Engine, log *zerolog.Logger, config *config.Manager, storage *storage.Storage, symbols *symbols.Manager) {
	h := &handles{log, config, storage, symbols}
	r.GET("/version", h.handleVersion())
	r.GET("/leaderboard/:symbol", h.getLeaderBoard())
	r.GET("/symbols", h.getSymbols())
}
