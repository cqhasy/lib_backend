package service

import (
	"AILN/app/model/document"
	"AILN/app/model/user"
	"AILN/app/request"
)

type VisitorService struct{}

func (v *VisitorService) GetDocument(req *request.GetDocumentReq) (docs []*document.Document, err error) {
	if req.Block == "新闻中心" || req.Block == "通知公告" || req.Block == "工作动态" {
		docs, err = document.FindByBlockGroup(req.Block, req.Group)
	} else {
		var doc *document.Document
		doc, err = document.FindOneByBlockGroup(req.Block, req.Group)
		docs = append(docs, doc)
	}
	if err != nil {
		return nil, err
	}
	return
}

func (v *VisitorService) GetDocumentDetail(id uint) (*document.Document, error) {
	return document.FindOne(id)
}

func (v *VisitorService) GetUsersWithOffset(page uint, pagesize uint) ([]*user.User, error) {
	if page == 0 {
		page = 1
	}
	if pagesize == 0 {
		pagesize = 10
	}
	users, err := user.FindInPage(page, pagesize)
	return users, err
}
