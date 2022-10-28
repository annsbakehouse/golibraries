package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategoryPageProductModel struct {
	ID                    string     `gorm:"column:id;primary_key" json:"id"`
	ProductCategoryPageID string     `gorm:"column:mkt_product_category_page_id" json:"mkt_product_category_page_id"`
	ProductID             string     `gorm:"column:product_id" json:"product_id"`
	ProductModelID        string     `gorm:"type:product_model_id" json:"product_model_id"`
	ProductModelDataID    string     `gorm:"type:product_model_data_id" json:"product_model_data_id"`
	Ordering              int        `gorm:"type:ordering" json:"ordering"`
	CreatedBy             NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt             time.Time  `gorm:"column:created_at" json:"crated_at"`
}

type ProductCategoryPageProductModelPreload struct {
	ID                    string                   `gorm:"column:id;primary_key" json:"id"`
	ProductCategoryPageID string                   `gorm:"column:mkt_product_category_page_id" json:"mkt_product_category_page_id"`
	ProductID             string                   `gorm:"column:product_id" json:"product_id"`
	ProductModelID        string                   `gorm:"type:product_model_id" json:"product_model_id"`
	ProductModelDataID    string                   `gorm:"type:product_model_data_id" json:"product_model_data_id"`
	Ordering              int                      `gorm:"type:ordering" json:"ordering"`
	ProductInfo           ProductModelPreload      `gorm:"foreignKey:ProductID;references:ID" json:"product_info"`
	ModelInfo             ProductModelModelPreload `gorm:"foreignKey:ProductModelID;references:ID" json:"product_model_info"`
}

// TableName sets the insert table name for this struct type
func (p *ProductCategoryPageProductModel) TableName() string {
	return "mkt_product_category_page_product"
}
func (p *ProductCategoryPageProductModelPreload) TableName() string {
	return "mkt_product_category_page_product"
}

func (p *ProductCategoryPageProductModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductCategoryPageProductModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductCategoryPageProductModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductCategoryPageProductModel) BeforeDelete(tx *gorm.DB) (err error) {
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
type ProductCategoryPageProductSave struct {
	ProductID          string `json:"product_id" binding:"required"`
	ProductModelID     string `json:"product_model_id" binding:"required"`
	ProductModelDataID string `json:"product_model_data_id" binding:"required"`
	Ordering           int    `json:"ordering"`
}

// type ProductCollectionPageUpdate struct {
// 	ID                                       string                                             `json:"id" binding:"required"`
// 	Name                                     string                                             `json:"name" binding:"required"`
// 	ProductCollectionPageLanguage            []ProductCollectionLanguage                        `json:"product_collection_page_language" binding:"required,dive"`
// 	ProductCollectionPageFilterTagCategory   []ProductCollectionPageFilterTagCategoryModelInput `json:"product_collection_page_filter_tag_category" binding:"dive"`
// 	ProductCollectionPageProductCategoryPage []ProductCollectionPageProductCategoryPageInput    `json:"product_collection_page_product_category_page" binding:"dive"`
// }
// type ProductCategoryPageProductModelInfo struct {
// 	ID string `json:"id" binding:"required"`
// }
