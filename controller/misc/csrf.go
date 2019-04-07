package misc

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-csrf"
)

// Ping godoc
// @Summary CSRF
// @Description Get CSRF token and cookie
// @Tags miscellaneous
// @Accept json
// @Produce json
// @Success 200 {string} string "IN HEADER"
// @Header 200 {string} X-CSRF-TOKEN "CSRF Token hash value"
// @Router /csrf [get]
func Csrf(c *gin.Context) {
	c.Header("X-CSRF-TOKEN", csrf.GetToken(c))
	c.String(http.StatusOK, "IN HEADER")
	log.Println(c.ClientIP(), "response CSRF token", csrf.GetToken(c))
}
