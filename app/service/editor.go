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

func (e *EditorService) UpdateDocument(d *document.Document) error {
	if err := document.Update(d); err != nil {
		return fmt.Errorf("update document error: %v", err)
	}
	return nil
}

func (e *EditorService) CheckIfDocumentExist(block, group, title string) (document.Document, bool, error) {
	re, err := document.FindOneByBlockGroupAndTitle(block, group, title)
	if err != nil {
		return document.Document{}, false, fmt.Errorf("find document when plan set error: %v", err)
	}
	return *re, true, nil
}

func (e *EditorService) DeleteDocument(id uint) error {
	if err := document.DeleteOne(id); err != nil {
		return fmt.Errorf("delete document error: %v", err)
	}
	return nil
}
