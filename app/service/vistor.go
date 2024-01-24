package service

import (
	"AILN/app/model/document"
	"AILN/app/request"
)

type VisitorService struct{}

func (v *VisitorService) GetDocument(req *request.GetDocumentReq) (docs []*document.Document, err error) {
	if req.Group == "新闻中心" || req.Group == "通知公告" {
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
