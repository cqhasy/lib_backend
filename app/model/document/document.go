package document

import "AILN/app/common"

type Document struct {
	ID        uint   `json:"id" gorm:"id"`
	Key       string `json:"key" gorm:"key"`
	Value     string `json:"value" gorm:"value"`
	CreatedAt int64  `json:"created_at" gorm:"created_at"`
}

func Set(d *Document) error {
	return common.DB.Table("document").Create(d).Error
}

func FindOne(key string) (*Document, error) {
	var d Document
	if err := common.DB.Table("document").First(&d, "`key` = ?", key).Error; err != nil {
		return nil, err
	}
	return &d, nil
}
