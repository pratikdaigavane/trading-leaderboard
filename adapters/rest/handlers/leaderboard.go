package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// getLeaderBoard returns a handler function that returns the leaderboard for a given symbol.
func (h *handles) getLeaderBoard() gin.HandlerFunc {
	return func(c *gin.Context) {
		symbol := c.Param("symbol")
		h.logger.Info().Str("symbol", symbol).Msg("Getting leaderboard")
		leaderboard, err := h.manager.GetLeaderboard(symbol, h.config.Get().LeaderboardDepth)
		if err != nil {
			h.logger.Error().Err(err).Msg("Error getting leaderboard")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting leaderboard"})
			return
		}
		c.JSON(http.StatusOK, leaderboard)
	}
}

// getSymbols returns a handler function that returns the list of symbols.
func (h *handles) getSymbols() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, h.symbols.GetSymbols())
	}
}
