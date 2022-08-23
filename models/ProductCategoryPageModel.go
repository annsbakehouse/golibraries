package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategoryPageModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	Name      NullString `gorm:"column:name" json:"name"`
	Active    int        `gorm:"column:active" json:"active"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	// DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductCategoryPageModelPreload struct {
	ID                                   string                                             `gorm:"column:id;primary_key" json:"id"`
	Name                                 NullString                                         `gorm:"column:name" json:"name"`
	Active                               int                                                `gorm:"column:active" json:"active"`
	ProductCategoryPageFilterTagCategory []ProductCategoryPageFilterTagCategoryModelPreload `gorm:"foreignKey:ProductCategoryPageID;references:ID" json:"product_category_page_filter_tag_category_info"`
	ProductCategoryPageImage             []ProductCategoryPageImageModelPreload             `gorm:"foreignKey:ProductCategoryPageID;references:ID" json:"product_category_page_image"`
	ProductCategoryPageProduct           []ProductCategoryPageProductModelPreload           `gorm:"foreignKey:ProductCategoryPageID;references:ID" json:"product_category_page_product"`
	ProductCategoryPageLang              []LanguageTableModelPreload                        `gorm:"foreignKey:TableTargetID;references:ID" json:"product_category_page_lang"`
}

// TableName sets the insert table name for this struct type
func (p *ProductCategoryPageModel) TableName() string {
	return "mkt_product_category_page"
}
func (p *ProductCategoryPageModelPreload) TableName() string {
	return "mkt_product_category_page"
}

func (p *ProductCategoryPageModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductCategoryPageModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductCategoryPageModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductCategoryPageModel) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println(p.ID)
	var langModel LanguageTableModel
	tx.Model(&langModel).Where("table_target=?", "product_category_page").Where("table_target_id=?", p.ID).Delete(&langModel)
	return
}

type ProductCategoryLanguage struct {
	CountryID           string `json:"country_id" binding:"required"`
	CategoryTitle       string `json:"category_title" binding:"required"`
	Keyword             string `json:"keyword" binding:"required"`
	MetaTitle           string `json:"meta_title" binding:"required"`
	MetaDescription     string `json:"meta_description" binding:"required"`
	SeoFooter           string `json:"seo_footer" binding:"required"`
	URLSlug             string `json:"url_slug" binding:"required"`
	CategoryDescription string `json:"category_description" binding:"required"`
	ImageAlt            string `json:"image_alt" binding:"required"`
}
type ProductCategoryPageSave struct {
	Name                                 string                                     `json:"name" binding:"required"`
	Active                               int                                        `json:"active"`
	ProductCategoryPageImage             []ProductCategoryPageImageSave             `json:"product_category_page_image" binding:"dive"`
	ProductCategoryPageFilterTagCategory []ProductCategoryPageFilterTagCategorySave `json:"product_category_page_filter_tag_category" binding:"dive"`
	ProductCategoryPageProduct           []ProductCategoryPageProductSave           `json:"product_category_page_product" binding:"dive"`
	ProductCategoryLang                  []ProductCategoryLanguage                  `json:"product_categorylanguage" binding:"dive"`
}

type ProductCategoryPageUpdate struct {
	ID                                   string                                     `json:"id" binding:"required"`
	Name                                 string                                     `json:"name" binding:"required"`
	Active                               int                                        `json:"active"`
	ProductCategoryPageImage             []ProductCategoryPageImageUpdate           `json:"product_category_page_image" binding:"dive"`
	ProductCategoryPageFilterTagCategory []ProductCategoryPageFilterTagCategorySave `json:"product_category_page_filter_tag_category" binding:"dive"`
	ProductCategoryPageProduct           []ProductCategoryPageProductSave           `json:"product_category_page_product" binding:"dive"`
	ProductCategoryLang                  []ProductCategoryLanguage                  `json:"product_categorylanguage" binding:"dive"`
}

// type ProductCollectionPageUpdate struct {
// 	ID                                       string                                             `json:"id" binding:"required"`
// 	Name                                     string                                             `json:"name" binding:"required"`
// 	ProductCollectionPageLanguage            []ProductCollectionLanguage                        `json:"product_collection_page_language" binding:"required,dive"`
// 	ProductCollectionPageFilterTagCategory   []ProductCollectionPageFilterTagCategoryModelInput `json:"product_collection_page_filter_tag_category" binding:"dive"`
// 	ProductCollectionPageProductCategoryPage []ProductCollectionPageProductCategoryPageInput    `json:"product_collection_page_product_category_page" binding:"dive"`
// }
type ProductCategoryPageModelInfo struct {
	ID string `json:"id" binding:"required"`
}

type ProductCategoryPageModelImageInfo struct {
	ID   string `json:"id" binding:"required"`
	Type int    `json:"type" binding:"required"`
}
