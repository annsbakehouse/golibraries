package models

import (
	"time"

	"gorm.io/gorm"
)

type CouponModel struct {
	ID                string     `gorm:"column:id;primary_key" json:"id"`
	CouponName        NullString `gorm:"column:coupon_name;default:null" json:"coupon_name"`
	CouponCode        string     `gorm:"column:coupon_code;index:,unique;" json:"coupon_code"`
	CouponType        int        `gorm:"column:coupon_type;default:1" json:"coupon_type"`
	CouponValueType   int        `gorm:"column:coupon_value_type;default:1" json:"coupon_value_type"`
	CouponStart       string     `gorm:"column:coupon_start" json:"coupon_start"`
	CouponEnd         string     `gorm:"column:coupon_end" json:"coupon_end"`
	CouponValue       float64    `gorm:"column:coupon_value;default:0" json:"coupon_value"`
	CustomersRelation int        `gorm:"column:customers_relation;default:1" json:"customers_relation"`
	CouponRulesID     NullString `gorm:"column:coupon_rules_id" json:"coupon_rules_id"`
	Active            int        `gorm:"column:active" json:"active"`
	Created_By        string     `gorm:"column:created_by" json:"created_by"`
	CreatedAt         time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated_By        string     `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt         time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type CouponModelPreload struct {
	ID                string                        `gorm:"column:id;primary_key" json:"id"`
	CouponName        NullString                    `gorm:"column:coupon_name;default:null" json:"coupon_name"`
	CouponCode        string                        `gorm:"column:coupon_code;index:,unique;" json:"coupon_code"`
	CouponType        int                           `gorm:"column:coupon_type;default:1" json:"coupon_type"`
	CouponValueType   int                           `gorm:"column:coupon_value_type;default:1" json:"coupon_value_type"`
	CouponStart       string                        `gorm:"column:coupon_start" json:"coupon_start"`
	CouponEnd         string                        `gorm:"column:coupon_end" json:"coupon_end"`
	CouponValue       float64                       `gorm:"column:coupon_value;default:0" json:"coupon_value"`
	CustomersRelation int                           `gorm:"column:customers_relation;default:1" json:"customers_relation"`
	CouponRulesID     NullString                    `gorm:"column:coupon_rules_id" json:"coupon_rules_id"`
	Active            int                           `gorm:"column:active" json:"active"`
	CouponCustomers   []CouponCustomersModelPreload `gorm:"->;foreignKey:MktCouponID;references:ID" json:"coupon_customers"`
	CouponProduct     []CouponProductModelPreload   `gorm:"->;foreignKey:MktCouponID;references:ID" json:"coupon_product"`
	CouponRules       CouponRulesModelPreload       `gorm:"->;foreignKey:ID;references:CouponRulesID" json:"coupon_rules"`
	CouponUsage       []CouponUsageModelPreload     `gorm:"->;foreignKey:CouponID;references:ID" json:"coupon_usage"`
}

// TableName sets the insert table name for this struct type
func (c *CouponModel) TableName() string {
	return "mkt_coupon"
}

func (c *CouponModelPreload) TableName() string {
	return "mkt_coupon"
}
func (p *CouponModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *CouponModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *CouponModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *CouponModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type CouponModelSave struct {
	CouponName        string                     `json:"coupon_name" binding:"required"`
	CouponCode        string                     `json:"coupon_code"`
	CouponType        int                        `json:"coupon_type" binding:"required"`
	CouponStart       string                     `json:"coupon_start" binding:"required" time_format:"2006-01-02 15:04:05"`
	CouponEnd         string                     `json:"coupon_end" binding:"required" time_format:"2006-01-02 15:04:05"`
	CouponValue       float64                    `json:"coupon_value" binding:"required"`
	CouponValueType   int                        `json:"coupon_value_type" binding:"required"`
	CustomersRelation int                        `json:"customers_relation" binding:"required"`
	CouponRulesID     string                     `json:"coupon_rules_id"`
	IsMultiple        bool                       `json:"is_multiple"`
	IncNum            int                        `json:"inc_num"`
	Prefix            string                     `json:"prefix"`
	Active            int                        `json:"active"`
	CouponProduct     []CouponProductModelSave   `json:"coupon_product"`
	CouponCustomers   []CouponCustomersModelSave `json:"coupon_customers"`
}
type CouponModelUpdate struct {
	ID                string                       `json:"id" binding:"required"`
	CouponName        string                       `json:"coupon_name" binding:"required"`
	CouponCode        string                       `json:"coupon_code"`
	CouponType        int                          `json:"coupon_type" binding:"required"`
	CouponStart       string                       `json:"coupon_start" binding:"required" time_format:"2006-01-02 15:04:05"`
	CouponEnd         string                       `json:"coupon_end" binding:"required" time_format:"2006-01-02 15:04:05"`
	CouponValue       float64                      `json:"coupon_value" binding:"required"`
	CouponValueType   int                          `json:"coupon_value_type" binding:"required"`
	CustomersRelation int                          `json:"customers_relation" binding:"required"`
	CouponRulesID     string                       `json:"coupon_rules_id"`
	Active            int                          `json:"active"`
	CouponProduct     []CouponProductModelUpdate   `json:"coupon_product"`
	CouponCustomers   []CouponCustomersModelUpdate `json:"coupon_customers"`
}

type CouponModelInfo struct {
	ID string `json:"id" binding:"required"`
}
