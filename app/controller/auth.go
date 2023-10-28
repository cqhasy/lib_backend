package controller

import (
	"AILN/app/core/validator/request"
	"AILN/app/lib/email"
	"AILN/app/lib/response"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	SendEmailCodeReq request.SendEmailCodeReq
}

func (a *Auth) SendEmailCode(c *gin.Context) {
	req := a.SendEmailCodeReq
	objEmail := req.Email
	err := email.Push(&email.Msg{
		Email: objEmail,
		Type:  0,
	})
	if err != nil {
		response.FailMsg(c, err.Error())
		return
	}
	response.OkMsg(c, "发送成功")
}
