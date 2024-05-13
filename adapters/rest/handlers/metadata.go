package handlers

import "github.com/gin-gonic/gin"

// handleVersion returns a handler function that returns the version of the service.
func (h *handles) handleVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": h.config.GetServiceConfig().Version,
		})
	}
}
