package response

import (
	"AILN/app/model/document"
)

type GetDocumentResponse struct {
	Docs []*document.Document
}

type SimpleDocument struct {
	ID       uint   `json:"id" form:"id" binding:"required"`
	Title    string `json:"title" form:"title" binding:"required"`
	CreateAt int64  `json:"create_at" form:"create_at" binding:"required"`
}

type UserInfo struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}

type GetSimpleDocumentResponse struct {
	Docs []*SimpleDocument
}

type GetDocumentDetailResponse struct {
	Docs *document.Document
}

type GetUsersResponse struct {
	Users []*UserInfo
}
