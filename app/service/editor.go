package service

import (
	"AILN/app/model/document"
	"AILN/app/request"
	"fmt"
	"time"
)

type EditorService struct{}

func (e *EditorService) CreateDocument(req *request.CreateDocumentReq) error {
	d := &document.Document{
		Block:    req.Block,
		Group:    req.Group,
		Title:    req.Title,
		CreateAt: time.Now().Unix(),
		Content:  req.Content,
	}
	if err := document.Set(d); err != nil {
		return fmt.Errorf("set document error: %v", err)
	}
	return nil
}

func (e *EditorService) DeleteDocument(id uint) error {
	if err := document.DeleteOne(id); err != nil {
		return fmt.Errorf("delete document error: %v", err)
	}
	return nil
}
