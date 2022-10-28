package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategoryPageFilterTagCategoryModel struct {
	ID                         string     `gorm:"column:id;primary_key" json:"id"`
	ProductCategoryPageID      string     `gorm:"column:mkt_product_category_page_id" json:"mkt_product_category_page_id"`
	ProductFilterTagCategoryID string     `gorm:"column:mkt_product_filter_tag_category_id" json:"mkt_product_filter_tag_category_id"`
	Ordering                   int        `gorm:"column:ordering" json:"ordering"`
	CreatedBy                  NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt                  time.Time  `gorm:"column:created_at" json:"crated_at"`
}

type ProductCategoryPageFilterTagCategoryModelPreload struct {
	ID                           string                               `gorm:"column:id;primary_key" json:"id"`
	ProductCategoryPageID        string                               `gorm:"column:mkt_product_category_page_id" json:"mkt_product_category_page_id"`
	ProductFilterTagCategoryID   string                               `gorm:"column:mkt_product_filter_tag_category_id" json:"mkt_product_filter_tag_category_id"`
	Ordering                     int                                  `gorm:"column:ordering" json:"ordering"`
	ProductFilterTagCategoryInfo ProductFilterTagCategoryModelPreload `gorm:"foreignKey:ProductFilterTagCategoryID;references:ID" json:"product_filter_tag_category"`
}

// TableName sets the insert table name for this struct type
func (p *ProductCategoryPageFilterTagCategoryModel) TableName() string {
	return "mkt_product_category_page_filter_tag_category"
}
func (p *ProductCategoryPageFilterTagCategoryModelPreload) TableName() string {
	return "mkt_product_category_page_filter_tag_category"
}

func (p *ProductCategoryPageFilterTagCategoryModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductCategoryPageFilterTagCategoryModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductCategoryPageFilterTagCategoryModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductCategoryPageFilterTagCategoryModel) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println(p.ID)
	var langModel LanguageTableModel
	tx.Model(&langModel).Where("table_target=?", "product_filter_tag").Where("column_name=?", "name").Where("table_target_id=?", p.ID).Delete(&langModel)
	return
}

// type ProductCollectionLanguage struct {
// 	CountryID       string `json:"country_id" binding:"required"`
// 	Keyword         string `json:"keyword" binding:"required"`
// 	MetaTitle       string `json:"meta_title" binding:"required"`
// 	MetaDescription string `json:"meta_description" binding:"required"`
// 	SeoFooter       string `json:"seo_footer" binding:"required"`
// 	URLSlug         string `json:"url_slug" binding:"required"`
// }
type ProductCategoryPageFilterTagCategorySave struct {
	ProductFilterTagCategoryID string `json:"mkt_product_filter_tag_category_id" binding:"required"`
	Ordering                   int    `json:"ordering"`
}

// type ProductCollectionPageUpdate struct {
// 	ID                                       string                                             `json:"id" binding:"required"`
// 	Name                                     string                                             `json:"name" binding:"required"`
// 	ProductCollectionPageLanguage            []ProductCollectionLanguage                        `json:"product_collection_page_language" binding:"required,dive"`
// 	ProductCollectionPageFilterTagCategory   []ProductCollectionPageFilterTagCategoryModelInput `json:"product_collection_page_filter_tag_category" binding:"dive"`
// 	ProductCollectionPageProductCategoryPage []ProductCollectionPageProductCategoryPageInput    `json:"product_collection_page_product_category_page" binding:"dive"`
// }
// type ProductCategoryPageFilterTagCategoryModelInfo struct {
// 	ID string `json:"id" binding:"required"`
// }
