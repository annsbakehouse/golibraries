package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductFilterTagModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	Name      NullString `gorm:"column:name;default:null" json:"name"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	// DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductFilterTagModelPreload struct {
	ID                       string                                `gorm:"column:id;primary_key" json:"id"`
	Name                     NullString                            `gorm:"column:name;default:null" json:"name"`
	ProductFilterProduct     []ProductFilterTagProductModelPreload `gorm:"foreignKey:MktProductFilterTagID;references:ID" json:"filter_tag_product"`
	ProductFilterTagLanguage []LanguageTableModelPreload           `gorm:"foreignKey:TableTargetID;references:ID" json:"filter_tag_language"`
	// DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (p *ProductFilterTagModel) TableName() string {
	return "mkt_product_filter_tag"
}
func (p *ProductFilterTagModelPreload) TableName() string {
	return "mkt_product_filter_tag"
}

func (p *ProductFilterTagModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductFilterTagModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductFilterTagModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductFilterTagModel) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println(p.ID)
	var langModel LanguageTableModel
	tx.Model(&langModel).Where("table_target=?", "product_filter_tag").Where("column_name=?", "name").Where("table_target_id=?", p.ID).Delete(&langModel)
	return
}

//strcture input
type ProductFilterTagModeLangInput struct {
	CountryID string `json:"country_id"  binding:"required"`
	Name      string `json:"name"  binding:"required"`
}
type ProductFilterTagModelSave struct {
	Name            string                          `json:"name" binding:"required"`
	Product         []string                        `json:"product"`
	LanguageContent []ProductFilterTagModeLangInput `json:"product_filter_tag_language" binding:"dive"`
}
type ProductFilterTagModelUpdate struct {
	ID              string                          `json:"id"  binding:"required"`
	Name            string                          `json:"name" binding:"required"`
	Product         []string                        `json:"product"`
	LanguageContent []ProductFilterTagModeLangInput `json:"product_filter_tag_language" binding:"dive"`
}
type ProductFilterTagModelInfo struct {
	ID string `json:"id" binding:"required"`
}
