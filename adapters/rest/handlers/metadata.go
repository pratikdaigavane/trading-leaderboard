package handlers

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// handleVersion godoc
// @Summary Get Version
// @Schemes
// @Description Get Version
// @Accept json
// @Produce json
// @Router /version [get]
func (h *handles) handleVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": h.config.GetServiceConfig().Version,
		})
	}
}
