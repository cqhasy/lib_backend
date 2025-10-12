package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UseCors(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		// 允许来源
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			// OPTIONS 请求要返回 200 且带上上面的所有头
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})
}
