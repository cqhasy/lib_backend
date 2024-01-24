package response

import "AILN/app/model/document"

type GetDocumentResponse struct {
	Docs []*document.Document
}

type SimpleDocument struct {
	ID       uint   `json:"id" form:"id" binding:"required"`
	Title    string `json:"title" form:"title" binding:"required"`
	CreateAt int64  `json:"create_at" form:"create_at" binding:"required"`
}

type GetSimpleDocumentResponse struct {
	Docs []*SimpleDocument
}

type GetDocumentDetailResponse struct {
	Docs *document.Document
}
