package models

import (
	"time"

	"gorm.io/gorm"
)

type CouponProductModel struct {
	ID                 string     `gorm:"column:id;primary_key" json:"id"`
	ProductID          string     `gorm:"column:product_id" json:"product_id"`
	ProductModelID     NullString `gorm:"column:product_model_id" json:"product_model_id"`
	ProductModelDataID NullString `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	MktCouponID        string     `gorm:"column:mkt_coupon_id" json:"mkt_coupon_id"`
	Created_By         string     `gorm:"column:created_by" json:"created_by"`
	CreatedAt          time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated_By         string     `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type CouponProductModelPreload struct {
	ID                   string                   `gorm:"column:id;primary_key" json:"id"`
	ProductID            string                   `gorm:"column:product_id" json:"product_id"`
	ProductModelID       string                   `gorm:"column:product_model_id" json:"product_model_id"`
	ProductModelDataID   string                   `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	MktCouponID          string                   `gorm:"column:mkt_coupon_id" json:"mkt_coupon_id"`
	ProductInfo          ProductModelPreload      `gorm:"->;foreignKey:ID;references:ProductID" json:"product_info"`
	ProductModelInfo     ProductModelModelPreload `gorm:"->;foreignKey:ID;references:ProductModelID" json:"product_model_info"`
	ProductModelDataInfo ProductModelDataPreload  `gorm:"->;foreignKey:ID;references:ProductModelDataID" json:"product_model_data_info"`
}

// TableName sets the insert table name for this struct type
func (c *CouponProductModel) TableName() string {
	return "mkt_coupon_product"
}

func (c *CouponProductModelPreload) TableName() string {
	return "mkt_coupon_product"
}
func (p *CouponProductModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *CouponProductModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *CouponProductModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *CouponProductModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type CouponProductModelSave struct {
	ProductID          string `json:"product_id" binding:"required"`
	ProductModelID     string `json:"product_model_id"`
	ProductModelDataID string `json:"product_model_data_id"`
}
type CouponProductModelUpdate struct {
	ID                 string `json:"id"`
	ProductID          string `json:"product_id" binding:"required"`
	ProductModelID     string `json:"product_model_id"`
	ProductModelDataID string `json:"product_model_data_id"`
}

type CouponProductModelInfo struct {
	ID string `json:"id" binding:"required"`
}
