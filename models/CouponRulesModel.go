package models

import (
	"time"

	"gorm.io/gorm"
)

type CouponRulesModel struct {
	ID         string     `gorm:"column:id;primary_key" json:"id"`
	RulesName  string     `gorm:"column:rules_name" json:"rules_name"`
	RulesQuery NullString `gorm:"column:rules_query" json:"rules_query"`
	Created_By string     `gorm:"column:created_by" json:"created_by"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated_By string     `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type CouponRulesModelPreload struct {
	ID         string               `gorm:"column:id;primary_key" json:"id"`
	RulesName  string               `gorm:"column:rules_name" json:"rules_name"`
	RulesQuery NullString           `gorm:"column:rules_query" json:"rules_query"`
	Coupons    []CouponModelPreload `gorm:"->;foreignKey:CouponRulesID;references:ID" json:"coupons"`
}

// TableName sets the insert table name for this struct type
func (c *CouponRulesModel) TableName() string {
	return "mkt_coupon_rules"
}

func (c *CouponRulesModelPreload) TableName() string {
	return "mkt_coupon_rules"
}
func (p *CouponRulesModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *CouponRulesModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *CouponRulesModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *CouponRulesModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type CouponRulesModelSave struct {
	RulesName  string `json:"rules_name" binding:"required"`
	RulesQuery string `json:"rules_query" binding:"required"`
}
type CouponRulesModelUpdate struct {
	ID         string `json:"id" binding:"required"`
	RulesName  string `json:"rules_name" binding:"required"`
	RulesQuery string `json:"rules_query" binding:"required"`
}

type CouponRulesModelInfo struct {
	ID string `json:"id" binding:"required"`
}
