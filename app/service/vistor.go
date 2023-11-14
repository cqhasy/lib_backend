package service

import "AILN/app/model/document"

type VisitorService struct{}

func (v *VisitorService) GetValue(key string) (string, error) {
	d, err := document.FindOne(key)
	if err != nil {
		return "", err
	}
	return d.Value, nil
}
