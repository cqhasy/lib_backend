package request

type LoginReq struct {
	Username string `json:"username" form:"username" binding:"alphanum"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}
