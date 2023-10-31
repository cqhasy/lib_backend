package request

type SendEmailCodeReq struct {
	Email string `json:"email" form:"email" binding:"email"`
	Usage string `json:"usage" form:"usage"`
}
