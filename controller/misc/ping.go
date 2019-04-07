package misc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary PING-PONG
// @Description Ping health check
// @Tags miscellaneous
// @Accept json
// @Produce json
// @Success 200 {object} misc.Message
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, PingMessage{
		Message: "pong",
	})
}
