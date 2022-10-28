package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductCollectionPageFilterTagCategoryModel struct {
	ID                         string     `gorm:"column:id;primary_key" json:"id"`
	ProductCollectionPageID    string     `gorm:"column:mkt_product_collection_page_id" json:"mkt_product_collection_page_id"`
	ProductFilterTagCategoryID string     `gorm:"column:mkt_product_filter_tag_category_id" json:"mkt_product_filter_tag_category_id"`
	Ordering                   int        `gorm:"column:ordering" json:"ordering"`
	CreatedBy                  NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt                  time.Time  `gorm:"column:created_at" json:"crated_at"`
}

type ProductCollectionPageFilterTagCategoryModelPreload struct {
	ID                           string                               `gorm:"column:id;primary_key" json:"id"`
	ProductCollectionPageID      string                               `gorm:"column:mkt_product_collection_page_id" json:"mkt_product_collection_page_id"`
	ProductFilterTagCategoryID   string                               `gorm:"column:mkt_product_filter_tag_category_id" json:"mkt_product_filter_tag_category_id"`
	Ordering                     int                                  `gorm:"column:ordering" json:"ordering"`
	ProductFilterTagCategoryInfo ProductFilterTagCategoryModelPreload `gorm:"foreignKey:ProductFilterTagCategoryID;references:ID" json:"product_filter_tag_category_info"`
}

// TableName sets the insert table name for this struct type
func (p *ProductCollectionPageFilterTagCategoryModel) TableName() string {
	return "mkt_product_collection_page_filter_tag_category"
}
func (p *ProductCollectionPageFilterTagCategoryModelPreload) TableName() string {
	return "mkt_product_collection_page_filter_tag_category"
}

func (p *ProductCollectionPageFilterTagCategoryModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductCollectionPageFilterTagCategoryModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductCollectionPageFilterTagCategoryModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductCollectionPageFilterTagCategoryModel) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println(p.ID)
	var langModel LanguageTableModel
	tx.Model(&langModel).Where("table_target=?", "product_filter_tag").Where("column_name=?", "name").Where("table_target_id=?", p.ID).Delete(&langModel)
	return
}

type ProductCollectionPageFilterTagCategoryModelInput struct {
	ProductFilterTagCategoryID string `json:"mkt_product_filter_tag_category_id" binding:"required"`
	Ordering                   int    `json:"ordering"`
}
