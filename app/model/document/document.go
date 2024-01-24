package document

import (
	"AILN/app/common"
	"fmt"
)

type Document struct {
	ID       uint   `json:"id" gorm:"id"`
	Block    string `json:"block" form:"block" binding:"required"`
	Group    string `json:"group" form:"group" gorm:"column:group_name" binding:"required"`
	Title    string `json:"title" form:"title" binding:"required"`
	CreateAt int64  `json:"create_at" form:"create_at" binding:"required"`
	Content  string `json:"content" form:"content" binding:"required"`
}

func Set(d *Document) error {
	return common.DB.Table("document").Create(d).Error
}

func FindOneByBlockGroup(block string, group string) (*Document, error) {
	var d Document
	if err := common.DB.Table("document").First(&d, "block = ? AND group_name = ?", block, group).Error; err != nil {
		return nil, err
	}
	return &d, nil
}

func FindByBlockGroup(block string, group string) ([]*Document, error) {
	var d []*Document
	//这里传址？
	if err := common.DB.Table("document").Find(&d, "block = ? AND group_name = ?", block, group).Error; err != nil {
		return nil, err
	}
	return d, nil
}

func DeleteOne(id uint) error {
	result := common.DB.Table("document").Delete(&Document{ID: id})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("delete one document: error not found")
	}
	return nil
}

func FindOne(id uint) (*Document, error) {
	var d Document
	if err := common.DB.Table("document").First(&d, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &d, nil
}
