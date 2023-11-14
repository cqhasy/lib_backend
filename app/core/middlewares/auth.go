package middlewares

import (
	"AILN/app/common/jwt"
	"AILN/app/response"
	"github.com/gin-gonic/gin"
)

func UseJwt(r *gin.RouterGroup) {
	r.Use(func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		claims, err := jwt.Parse(token)
		if err != nil {
			response.FailMsg(c, "Unauthorized")
			c.Abort()
			return
		}
		c.Set("userID", claims["userID"])
		c.Next()
	})
}
