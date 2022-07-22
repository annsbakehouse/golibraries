package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductFilterTagCategoryFilterTagModel struct {
	ID                            string     `gorm:"column:id;primary_key" json:"id"`
	MKTProductFilterTagID         string     `gorm:"column:mkt_product_filter_tag_id" json:"mkt_product_filter_tag_id"`
	MKTProductFilterTagCategoryID string     `gorm:"column:mkt_product_filter_tag_category_id" json:"mkt_product_filter_tag_category_id"`
	Ordering                      int        `gorm:"column:ordering" json:"ordering"`
	CreatedBy                     NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt                     time.Time  `gorm:"column:created_at" json:"crated_at"`
	// DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductFilterTagCategoryFilterTagModelPreload struct {
	ID                            string                         `gorm:"column:id;primary_key" json:"id"`
	MKTProductFilterTagID         string                         `gorm:"column:mkt_product_filter_tag_id" json:"mkt_product_filter_tag_id"`
	MKTProductFilterTagCategoryID string                         `gorm:"column:mkt_product_filter_tag_category_id" json:"mkt_product_filter_tag_category_id"`
	Ordering                      int                            `gorm:"column:ordering" json:"ordering"`
	ProductFilterTagIDInfo        []ProductFilterTagModelPreload `gorm:"foreignKey:ID;references:MKTProductFilterTagID" json:"product_filter_tag"`
}

// TableName sets the insert table name for this struct type
func (p *ProductFilterTagCategoryFilterTagModel) TableName() string {
	return "mkt_product_filter_tag_category_filter_tag"
}
func (p *ProductFilterTagCategoryFilterTagModelPreload) TableName() string {
	return "mkt_product_filter_tag_category_filter_tag"
}

func (p *ProductFilterTagCategoryFilterTagModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductFilterTagCategoryFilterTagModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductFilterTagCategoryFilterTagModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductFilterTagCategoryFilterTagModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

//strcture input
type ProductFilterTagCategoryFilterTagModelLangInput struct {
	CountryID string `json:"country_id"  binding:"required"`
	Name      string `json:"name"  binding:"required"`
}
type ProductFilterTagCategoryFilterTagModelSave struct {
	Name            string                          `json:"name" binding:"required"`
	Product         []string                        `json:"product"`
	LanguageContent []ProductFilterTagModeLangInput `json:"product_filter_tag_language" binding:"dive"`
}
type ProductFilterTagCategoryFilterTagModelUpdate struct {
	ID              string                          `json:"id"  binding:"required"`
	Name            string                          `json:"name" binding:"required"`
	Product         []string                        `json:"product"`
	LanguageContent []ProductFilterTagModeLangInput `json:"product_filter_tag_language" binding:"dive"`
}
type ProductFilterTagCategoryFilterTagModelInfo struct {
	ID string `json:"id" binding:"required"`
}
