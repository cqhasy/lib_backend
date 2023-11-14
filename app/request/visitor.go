package request

type GetValueReq struct {
	Key string `json:"key" form:"key" binding:"required"`
}
