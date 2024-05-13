package handlers

import (
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"leaderboard/docs"
	"leaderboard/internal/config"
	"leaderboard/internal/manager"
	"leaderboard/internal/symbols"
	"time"
)

type handles struct {
	logger  *zerolog.Logger
	config  *config.Manager
	symbols *symbols.Manager
	manager *manager.Manager
	store   *persistence.InMemoryStore
}

func SetupHandlers(r *gin.Engine, log *zerolog.Logger, config *config.Manager, symbols *symbols.Manager, manager *manager.Manager) {
	// Create a new in-memory store for caching the leaderboard
	docs.SwaggerInfo.BasePath = "/"
	store := persistence.NewInMemoryStore(1 * time.Minute)
	h := &handles{log, config, symbols, manager, store}
	r.GET("/version", h.handleVersion())
	r.GET("/leaderboard/:symbol", h.getLeaderBoard())
	r.GET("/symbols", h.getSymbols())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
