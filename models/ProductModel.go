package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductModel struct {
	ID            string     `gorm:"column:id;primary_key" json:"id"`
	Name          NullString `gorm:"column:name;default:null" json:"name"`
	PID           NullString `gorm:"column:pid;default:null" json:"pid"`
	ProductTypeId NullString `gorm:"column:product_type_id" json:"product_type_id"`
	Active        int        `gorm:"column:active" json:"active"`
	CreatedBy     NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt     time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy     NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt     time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy     NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
	ProductType ProductTypeModel `gorm:"foreignKey:ID;references:ProductTypeId" json:"product_type"`
}

type ProductModelPreload struct {
	ID            string           `gorm:"column:id;primary_key" json:"id"`
	Name          NullString       `gorm:"column:name;default:null" json:"name"`
	PID           NullString       `gorm:"column:pid;default:null" json:"pid"`
	ProductTypeId NullString       `gorm:"column:product_type_id" json:"product_type_id"`
	Active        int              `gorm:"column:active" json:"active"`
	ProductType   ProductTypeModel `gorm:"foreignKey:ID;references:ProductTypeId" json:"product_type"`
	// DeletedAt     soft_delete.DeletedAt      `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
	ModelData                  []ProductModelDataPreload            `gorm:"foreignKey:ProductId;references:ID" json:"model_data"`
	ProductImage               []ProductImageModelPreload           `gorm:"foreignKey:ProductID;references:ID" json:"product_image"`
	ProductPlarformDisplay     []ProductPlatformDisplayModelPreload `gorm:"foreignKey:ProductID;references:ID" json:"product_platform"`
	ProductDisplayName         []LanguageTableModelPreload          `gorm:"foreignKey:TableTargetID;references:ID" json:"product_display_name"`
	ContentDisplay             []LanguageTableModelPreload          `gorm:"foreignKey:TableTargetID;references:ID" json:"content_display"`
	MktProductContentAccordion []MKTProductAccordionModelPreload    `gorm:"foreignKey:ProductID;references:ID" json:"product_content_accordion"`
}

// TableName sets the insert table name for this struct type
func (p *ProductModel) TableName() string {
	return "product"
}
func (p *ProductModelPreload) TableName() string {
	return "product"
}

func (p *ProductModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model ProductModel
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by", ActiveUser)
	return
}
func (p *ProductModel) AfterDelete(tx *gorm.DB) (err error) {
	// var model ProductModel
	idProduct := p.ID
	var langModel LanguageTableModel
	tx.Model(&langModel).Where("table_target=?", "product").Where("table_target_id=?", idProduct).Delete(&langModel)
	tx.Model(&langModel).Where("table_target=?", "product_image").Where("table_target_id=?", idProduct).Delete(&langModel)
	//delete product display name lang

	return
}

//strcture input
type ProductModelInput struct {
	Name          string `json:"name" binding:"required"`
	PID           string `json:"pid" binding:"required"`
	ProductTypeId string `json:"product_type_id" binding:"required"`
	Active        int    `json:"active"`
}
type ProductModelUpdate struct {
	ID            string `json:"id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	PID           string `json:"pid" binding:"required"`
	ProductTypeId string `json:"product_type_id" binding:"required"`
	Active        int    `json:"active"`
}
type ProductModelInfo struct {
	ID string `json:"id" binding:"required"`
}
