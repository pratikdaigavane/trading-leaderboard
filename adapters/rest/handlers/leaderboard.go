package handlers

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// getLeaderBoard godoc
// @Summary Get the leaderboard for a given symbol.
// @Schemes
// @Description Get the leaderboard for a given symbol.
// @Accept json
// @Produce json
// @Router /leaderboard/{symbol} [get]
// @Param symbol path string true "Symbol"
// @Success 200 {object} []models.UserTradeStat
func (h *handles) getLeaderBoard() gin.HandlerFunc {
	return cache.CachePage(h.store, 1*time.Minute, func(c *gin.Context) {
		symbol := c.Param("symbol")
		h.logger.Info().Str("symbol", symbol).Msg("Getting leaderboard")
		leaderboard, err := h.manager.GetLeaderboard(symbol, h.config.Get().LeaderboardDepth)
		if err != nil {
			h.logger.Error().Err(err).Msg("Error getting leaderboard")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting leaderboard"})
			return
		}
		c.JSON(http.StatusOK, leaderboard)
	})
}

// getSymbols godoc
// @Summary Get the list of symbols.
// @Schemes
// @Description Get the list of symbols.
// @Accept json
// @Produce json
// @Router /symbols [get]
// @Success 200 {object} map[string]models.Symbol
func (h *handles) getSymbols() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, h.symbols.GetSymbols())
	}
}
