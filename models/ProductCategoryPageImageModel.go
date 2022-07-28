package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategoryPageImageModel struct {
	ID                    string     `gorm:"column:id;primary_key" json:"id"`
	ProductCategoryPageID string     `gorm:"column:mkt_product_category_page_id" json:"mkt_product_category_page_id"`
	Image                 string     `gorm:"column:image" json:"image"`
	Type                  int        `gorm:"column:type" json:"type"`
	CreatedBy             NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt             time.Time  `gorm:"column:created_at" json:"crated_at"`
}

type ProductCategoryPageImageModelPreload struct {
	ID                    string `gorm:"column:id;primary_key" json:"id"`
	ProductCategoryPageID string `gorm:"column:mkt_product_category_page_id" json:"mkt_product_category_page_id"`
	Image                 string `gorm:"column:image" json:"image"`
	Type                  int    `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (p *ProductCategoryPageImageModel) TableName() string {
	return "mkt_product_category_page_image"
}
func (p *ProductCategoryPageImageModelPreload) TableName() string {
	return "mkt_product_category_page_image"
}

func (p *ProductCategoryPageImageModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductCategoryPageImageModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductCategoryPageImageModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductCategoryPageImageModel) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println(p.ID)
	var langModel LanguageTableModel
	tx.Model(&langModel).Where("table_target=?", "product_category_page").Where("table_target_id=?", p.ID).Delete(&langModel)
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
type ProductCategoryPageImageSave struct {
	Image string `json:"image" binding:"required"`
	Type  int    `json:"type" binding:"required"`
}
type ProductCategoryPageImageUpdate struct {
	Image string `json:"image"`
	Type  int    `json:"type" binding:"required"`
}

// type ProductCollectionPageUpdate struct {
// 	ID                                       string                                             `json:"id" binding:"required"`
// 	Name                                     string                                             `json:"name" binding:"required"`
// 	ProductCollectionPageLanguage            []ProductCollectionLanguage                        `json:"product_collection_page_language" binding:"required,dive"`
// 	ProductCollectionPageFilterTagCategory   []ProductCollectionPageFilterTagCategoryModelInput `json:"product_collection_page_filter_tag_category" binding:"dive"`
// 	ProductCollectionPageProductCategoryPage []ProductCollectionPageProductCategoryPageInput    `json:"product_collection_page_product_category_page" binding:"dive"`
// }
// type ProductCategoryPageImageModelInfo struct {
// 	ID string `json:"id" binding:"required"`
// }
