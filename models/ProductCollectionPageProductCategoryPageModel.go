package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductCollectionPageProductCategoryPageModel struct {
	ID                      string     `gorm:"column:id;primary_key" json:"id"`
	ProductCollectionPageID string     `gorm:"column:mkt_product_collection_page_id" json:"mkt_product_collection_page_id"`
	ProductCategoryPageID   string     `gorm:"column:mkt_product_category_page_id" json:"mkt_product_category_page_id"`
	Ordering                int        `gorm:"column:ordering" json:"ordering"`
	CreatedBy               NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt               time.Time  `gorm:"column:created_at" json:"crated_at"`
}

type ProductCollectionPageProductCategoryPageModelPreload struct {
	ID                      string                          `gorm:"column:id;primary_key" json:"id"`
	ProductCollectionPageID string                          `gorm:"column:mkt_product_collection_page_id" json:"mkt_product_collection_page_id"`
	ProductCategoryPageID   string                          `gorm:"column:mkt_product_category_page_id" json:"mkt_product_category_page_id"`
	Ordering                int                             `gorm:"column:ordering" json:"ordering"`
	ProductCategoryPageInfo ProductCategoryPageModelPreload `gorm:"foreignKey:ProductCategoryPageID;references:ID" json:"product_category_page_info"`
}

// TableName sets the insert table name for this struct type
func (p *ProductCollectionPageProductCategoryPageModel) TableName() string {
	return "mkt_product_collection_page_product_category_page"
}
func (p *ProductCollectionPageProductCategoryPageModelPreload) TableName() string {
	return "mkt_product_collection_page_product_category_page"
}

func (p *ProductCollectionPageProductCategoryPageModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductCollectionPageProductCategoryPageModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductCollectionPageProductCategoryPageModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductCollectionPageProductCategoryPageModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type ProductCollectionPageProductCategoryPageInput struct {
	ProductCategoryPageID string `json:"mkt_product_category_page_id"`
	Ordering              int    `json:"ordering"`
}
