package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductPricingModel struct {
	ID                 string     `gorm:"column:id;primary_key" json:"id"`
	CustomerLevelID    string     `gorm:"column:customer_level_id" json:"customer_level_id"`
	ProductModelDataID string     `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	Active             int        `gorm:"column:active" json:"active"`
	INC                int        `gorm:"column:inc" json:"inc"`
	Price              int        `gorm:"column:price" json:"price"`
	CreatedBy          NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt          time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy          NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy          NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt          soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductPricingModelPreload struct {
	ID                 string `gorm:"column:id;primary_key" json:"id"`
	CustomerLevelID    string `gorm:"column:customer_level_id" json:"customer_level_id"`
	ProductModelDataID string `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	Active             int    `gorm:"column:active" json:"active"`
	INC                int    `gorm:"column:inc" json:"inc"`
	Price              int    `gorm:"column:price" json:"price"`
	// DeletedAt          soft_delete.DeletedAt    `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
	CustomerLevel CustomerLevelModelInfo `gorm:"foreignKey:ID;references:CustomerLevelID" json:"customer_level"`
}

// TableName sets the insert table name for this struct type
func (p *ProductPricingModel) TableName() string {
	return "product_price"
}
func (p *ProductPricingModelPreload) TableName() string {
	return "product_price"
}

func (p *ProductPricingModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductPricingModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductPricingModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductPricingModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

//strcture input
type ProductPricingModelInput struct {
	CustomerLevelID    map[string][]string `json:"customer_level_id"`
	ProductModelDataID []string            `json:"product_model_data_id" binding:"required"`
	Price              map[string][]int    `json:"price"`
}
type ProductPricingModelUpdate struct {
	ID                 string `json:"id" binding:"required"`
	CustomerLevelID    string `json:"customer_level_id" binding:"required"`
	ProductModelDataID string `json:"product_model_data_id" binding:"required"`
	Active             int    `json:"active"`
	Price              int    `json:"price"`
}
type ProductPricingModellInfo struct {
	ID string `json:"id" binding:"required"`
}
