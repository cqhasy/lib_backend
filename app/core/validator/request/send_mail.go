package request

import (
	"AILN/app/controller"
	"AILN/app/lib/response"
	"github.com/gin-gonic/gin"
)

type SendEmailCodeReq struct {
	Email string `json:"email" form:"email" binding:"email"`
	Usage string `json:"usage" form:"usage"`
}

func (l SendEmailCodeReq) CheckParams(c *gin.Context) {

	if err := c.ShouldBind(&l); err != nil {
		response.FailMsg(c, err.Error())
		return
	}

	(&controller.Auth{SendEmailCodeReq: l}).SendEmailCode(c)

}
