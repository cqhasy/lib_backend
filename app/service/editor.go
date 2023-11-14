package service

import (
	"AILN/app/model/document"
	"fmt"
	"time"
)

type EditorService struct{}

func (e *EditorService) CreateDocument(key string, value string) error {
	d := &document.Document{
		Key:       key,
		Value:     value,
		CreatedAt: time.Now().Unix(),
	}
	if err := document.Set(d); err != nil {
		return fmt.Errorf("set document error: %v", err)
	}
	return nil
}
