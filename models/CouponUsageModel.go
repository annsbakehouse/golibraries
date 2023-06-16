package models

import (
	"time"

	"gorm.io/gorm"
)

type CouponUsageModel struct {
	ID                 string    `gorm:"column:id;primary_key" json:"id"`
	OrderInvoiceID     string    `gorm:"column:order_invoice_id" json:"order_invoice_id"`
	OrderDataID        string    `gorm:"column:order_data_id" json:"order_data_id"`
	CustomerID         string    `gorm:"column:customers_id" json:"customers_id"`
	CouponID           string    `gorm:"column:mkt_coupon_id" json:"mkt_coupon_id"`
	OrderProductID     string    `gorm:"column:order_product_id" json:"order_product_id"`
	ProductModelID     string    `gorm:"column:order_product_model_id" json:"order_product_model_id"`
	ProductModelDataID string    `gorm:"column:order_product_model_data_id" json:"order_product_model_data_id"`
	CouponValue        float64   `gorm:"column:coupon_value" json:"coupon_value"`
	ProductID          string    `gorm:"column:product_id" json:"product_id"`
	Created_By         string    `gorm:"column:created_by" json:"created_by"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
	Updated_By         string    `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type CouponUsageModelPreload struct {
	ID                 string       `gorm:"column:id;primary_key" json:"id"`
	OrderInvoiceID     string       `gorm:"column:order_invoice_id" json:"order_invoice_id"`
	OrderDataID        string       `gorm:"column:order_data_id" json:"order_data_id"`
	CustomerID         string       `gorm:"column:customers_id" json:"customers_id"`
	CouponID           string       `gorm:"column:mkt_coupon_id" json:"mkt_coupon_id"`
	OrderProductID     string       `gorm:"column:order_product_id" json:"order_product_id"`
	CouponValue        float64      `gorm:"column:coupon_value" json:"coupon_value"`
	ProductID          string       `gorm:"column:product_id" json:"product_id"`
	ProductModelID     string       `gorm:"column:order_product_model_id" json:"order_product_model_id"`
	ProductModelDataID string       `gorm:"column:order_product_model_data_id" json:"order_product_model_data_id"`
	OrderInvoice       InvoiceModel `gorm:"->;foreignKey:Id;references:OrderInvoiceID" json:"invoice_info"`
}

// TableName sets the insert table name for this struct type
func (c *CouponUsageModel) TableName() string {
	return "mkt_coupon_usage"
}

func (c *CouponUsageModelPreload) TableName() string {
	return "mkt_coupon_usage"
}
func (p *CouponUsageModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *CouponUsageModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *CouponUsageModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *CouponUsageModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
