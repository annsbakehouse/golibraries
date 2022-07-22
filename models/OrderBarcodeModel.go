package models

import (
	"gorm.io/gorm"
)

type OrderBarcodeModel struct {
	ID  string `gorm:"column:id;primary_key" json:"id"` //
	INC string `gorm:"->;column:inc" json:"inc"`        //
}

// TableName sets the insert table name for this struct type
func (c *OrderBarcodeModel) TableName() string {
	return "order_barcode"
}

func (c *OrderBarcodeModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderBarcodeModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderBarcodeModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}
