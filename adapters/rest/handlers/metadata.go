package handlers

import "github.com/gin-gonic/gin"

func (h *handles) handleVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": h.config.GetServiceConfig().Version,
		})
	}
}
