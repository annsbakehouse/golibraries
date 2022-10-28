package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductFilterTagCategoryModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	Name      NullString `gorm:"column:name;default:null" json:"name"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	// DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductFilterTagCategoryModelPreload struct {
	ID                                string                                          `gorm:"column:id;primary_key" json:"id"`
	Name                              NullString                                      `gorm:"column:name;default:null" json:"name"`
	ProductFilterTagCategoryFilterTag []ProductFilterTagCategoryFilterTagModelPreload `gorm:"foreignKey:MKTProductFilterTagCategoryID;references:ID" json:"product_filter_tag_category_filter_tag"`
	ProductFilterTagCategoryLanguage  []LanguageTableModelPreload                     `gorm:"foreignKey:TableTargetID;references:ID" json:"filter_tag_category_language"`
	// DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (p *ProductFilterTagCategoryModel) TableName() string {
	return "mkt_product_filter_tag_category"
}
func (p *ProductFilterTagCategoryModelPreload) TableName() string {
	return "mkt_product_filter_tag_category"
}

func (p *ProductFilterTagCategoryModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductFilterTagCategoryModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductFilterTagCategoryModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductFilterTagCategoryModel) BeforeDelete(tx *gorm.DB) (err error) {
	var langModel LanguageTableModel
	tx.Model(&langModel).Where("table_target=?", "product_filter_tag_category").Where("column_name=?", "name").Where("table_target_id=?", p.ID).Delete(&langModel)
	return
}

//strcture input
type ProductFilterTagCategoryModelLangInput struct {
	CountryID string `json:"country_id"  binding:"required"`
	Name      string `json:"name"  binding:"required"`
}
type ProductFilterTagCategoryModelSave struct {
	Name               string                                   `json:"name" binding:"required"`
	ProductFilterTagID []string                                 `json:"product_filter_tag_id"`
	LanguageContent    []ProductFilterTagCategoryModelLangInput `json:"product_filter_tag_category_language" binding:"required,dive"`
}
type ProductFilterTagCategoryModelUpdate struct {
	ID                 string                                   `json:"id"  binding:"required"`
	Name               string                                   `json:"name" binding:"required"`
	ProductFilterTagID []string                                 `json:"product_filter_tag_id"`
	LanguageContent    []ProductFilterTagCategoryModelLangInput `json:"product_filter_tag_category_language" binding:"required,dive"`
}
type ProductFilterTagCategoryModelInfo struct {
	ID string `json:"id" binding:"required"`
}
