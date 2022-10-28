package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderBarcodeModel struct {
	ID                   string     `gorm:"column:id;primary_key" json:"id"`
	OrderBarcode         int        `gorm:"->;column:order_barcode" json:"order_barcode"`
	OrderNestedProductID string     `gorm:"column:order_nested_product_id" json:"order_nested_product_id"`
	CreatedBy            NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt            *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy            NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt            *time.Time `gorm:"column:updated_at" json:"updated_at"`
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
