package request

type GetDocumentReq struct {
	Block string `json:"block" form:"block" binding:"required"`
	Group string `json:"group" form:"group" binding:"required"`
}

type GetDocumentDetailReq struct {
	ID uint `json:"id" form:"id" binding:"required"`
}

type GetUsersReq struct {
	PageSize   uint `json:"pageSize" form:"pageSize" binding:"required"`
	PageNumber uint `json:"pageNumber" form:"pageNumber" binding:"required"`
}
