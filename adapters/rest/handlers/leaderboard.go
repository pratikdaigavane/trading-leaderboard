package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handles) getLeaderBoard() gin.HandlerFunc {
	return func(c *gin.Context) {
		symbol := c.Param("symbol")
		h.logger.Info().Str("symbol", symbol).Msg("Getting leaderboard")
		leaderboard, err := h.storage.Operations.GetSortedList(symbol, h.config.Get().LeaderboardDepth)
		if err != nil {
			h.logger.Error().Err(err).Msg("Error getting leaderboard")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting leaderboard"})
			return
		}
		c.JSON(http.StatusOK, leaderboard)
	}
}

func (h *handles) getSymbols() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, h.symbols.GetSymbols())
	}
}
