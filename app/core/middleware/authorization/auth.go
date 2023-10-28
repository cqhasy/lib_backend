package authorization

import (
	"AILN/app/lib/jwt"
	"AILN/app/lib/response"
	"github.com/gin-gonic/gin"
)

func CheckJwt(c *gin.Context) {
	token := c.GetHeader("Authorization")
	valueMap, err := jwt.VerifyJWT(token)
	if err != nil {
		response.FailMsg(c, "Error verifying token")
		c.Abort()
		return
	}
	c.Set("ValueMap", valueMap)
}
