package request

type GetDocumentReq struct {
	Block string `json:"block" form:"block" binding:"required"`
	Group string `json:"group" form:"group" binding:"required"`
}

type GetDocumentDetailReq struct {
	ID uint `json:"id" form:"id" binding:"required"`
}
