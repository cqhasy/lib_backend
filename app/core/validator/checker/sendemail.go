package checker

import (
	"AILN/app/controller"
	"AILN/app/core/validator/request"
	"AILN/app/lib/response"
	"github.com/gin-gonic/gin"
)

type SendEmailChecker struct {
	request.SendEmailCodeReq
}

func (l SendEmailChecker) CheckParams(c *gin.Context) {
	if err := c.ShouldBind(&l); err != nil {
		response.FailMsg(c, err.Error())
		return
	}

	(&controller.Auth{SendEmailCodeReq: l.SendEmailCodeReq}).SendEmailCode(c)

}
