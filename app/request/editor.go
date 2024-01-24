package request

type CreateDocumentReq struct {
	Block    string `json:"block" form:"block" binding:"required"`
	Group    string `json:"group" form:"group" binding:"required"`
	Title    string `json:"title" form:"title" binding:"required"`
	CreateAt int64  `json:"create_at" form:"create_at" binding:"required"`
	Content  string `json:"content" form:"content" binding:"required"`
}
type DeleteDocumentReq struct {
	ID uint `json:"id" form:"id" binding:"required"`
}
