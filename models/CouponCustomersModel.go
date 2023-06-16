package models

import (
	"time"

	"gorm.io/gorm"
)

type CouponCustomersModel struct {
	ID               string     `gorm:"column:id;primary_key" json:"id"`
	MktCouponID      string     `gorm:"column:mkt_coupon_id" json:"mkt_coupon_id"`
	CustomersLevelID string     `gorm:"column:customers_level_id" json:"customers_level_id"`
	CustomersID      NullString `gorm:"column:customers_id" json:"customers_id"`
	Created_By       string     `gorm:"column:created_by" json:"created_by"`
	CreatedAt        time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated_By       string     `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt        time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type CouponCustomersModelPreload struct {
	ID               string                     `gorm:"column:id;primary_key" json:"id"`
	MktCouponID      string                     `gorm:"column:mkt_coupon_id" json:"mkt_coupon_id"`
	CustomersLevelID string                     `gorm:"column:customers_level_id" json:"customers_level_id"`
	CustomersID      string                     `gorm:"column:customers_id" json:"customers_id"`
	CustomersInfo    CustomerRetailModelPreload `gorm:"->;foreignKey:ID;references:CustomersID" json:"customers_info"`
}

// TableName sets the insert table name for this struct type
func (c *CouponCustomersModel) TableName() string {
	return "mkt_coupon_customers"
}

func (c *CouponCustomersModelPreload) TableName() string {
	return "mkt_coupon_customers"
}
func (p *CouponCustomersModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *CouponCustomersModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *CouponCustomersModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *CouponCustomersModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type CouponCustomersModelSave struct {
	CustomersLevelID string `json:"customers_level_id" binding:"required"`
	CustomersID      string `json:"customers_id"`
}
type CouponCustomersModelUpdate struct {
	ID               string `json:"id" binding:"required"`
	CustomersLevelID string `json:"customers_level_id" binding:"required"`
	CustomersID      string `json:"customers_id"`
}

type CouponCustomersModelInfo struct {
	ID string `json:"id" binding:"required"`
}
