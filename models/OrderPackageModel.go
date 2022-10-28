package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderPackageModel struct {
	ID             string     `gorm:"column:id;primary_key" json:"id"`                 //
	ProductID      string     `gorm:"column:product_id" json:"product_id"`             //
	ProductTypeID  string     `gorm:"column:product_type_id" json:"product_type_id"`   //
	OrderBarcodeID string     `gorm:"column:order_barcode_id" json:"order_barcode_id"` //
	OrderDataID    string     `gorm:"column:order_data_id" json:"order_data_id"`       //
	OrderInvoiceID string     `gorm:"column:order_invoice_id" json:"order_invoice_id"` //
	Status         int        `gorm:"column:status" json:"status"`                     //
	CreatedBy      string     `gorm:"column:created_by" json:"created_by"`             //
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`             //
}

type OrderPackageModelPreload struct {
	ID             string     `gorm:"column:id;primary_key" json:"id"`                 //
	ProductID      string     `gorm:"column:product_id" json:"product_id"`             //
	ProductTypeID  string     `gorm:"column:product_type_id" json:"product_type_id"`   //
	OrderBarcodeID string     `gorm:"column:order_barcode_id" json:"order_barcode_id"` //
	OrderDataID    string     `gorm:"column:order_data_id" json:"order_data_id"`       //
	OrderInvoiceID string     `gorm:"column:order_invoice_id" json:"order_invoice_id"` //
	Status         int        `gorm:"column:status" json:"status"`                     //
	CreatedBy      string     `gorm:"column:created_by" json:"created_by"`             //
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`             //
}

// TableName sets the insert table name for this struct type
func (c *OrderPackageModel) TableName() string {
	return "order_package"
}
func (c *OrderPackageModelPreload) TableName() string {
	return "order_package"
}

func (c *OrderPackageModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderPackageModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderPackageModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}
