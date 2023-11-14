package controller

import (
	"AILN/app/request"
	"AILN/app/response"
	"AILN/app/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Auth struct{}

var authService *service.AuthService

// @Summary 登录
// @Description Authenticates a user and returns a token upon successful login.
// @Accept json;multipart/form-data
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} response.LoginResponse "Authentication successful"
// @Router /api/v1/auth/login [post]
func (a *Auth) Login(c *gin.Context) {
	req := &request.LoginReq{}
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("params invalid error: %v", err))
		return
	}
	token, err := authService.Login(req.Username, req.Password)
	if err != nil {
		response.FailMsg(c, err.Error())
		return
	}
	response.OkMsgData(c, "登录成功", response.LoginResponse{Token: token})
}
