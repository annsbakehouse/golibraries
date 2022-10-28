package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductLibraryPropertyModel struct {
	ID               string     `gorm:"column:id;primary_key" json:"id"`
	Name             NullString `gorm:"column:name;default:null" json:"name"`
	ProductLibraryID string     `gorm:"column:product_library_id" json:"product_library_id"`
	CreatedBy        NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt        time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy        NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt        time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
}

type ProductLibraryPropertyModelPreload struct {
	ID                     string                         `gorm:"column:id;primary_key" json:"id"`
	Name                   NullString                     `gorm:"column:name;default:null" json:"name"`
	ProductLibraryID       string                         `gorm:"column:product_library_id" json:"product_library_id"`
	ProductLibraryNameInfo ProductLibraryNameModelPreload `gorm:"foreignKey:ID;references:ProductLibraryID" json:"product_library_name_info"`
	// DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (p *ProductLibraryPropertyModel) TableName() string {
	return "product_library_properties"
}
func (p *ProductLibraryPropertyModelPreload) TableName() string {
	return "product_library_properties"
}

func (p *ProductLibraryPropertyModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductLibraryPropertyModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductLibraryPropertyModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductLibraryPropertyModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model ProductLibraryPropertyModel
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by", ActiveUser)
	return
}

//strcture input
type ProductLibraryPropertyModelInput struct {
	Name             string `json:"name" binding:"required"`
	ProductLibraryID string `json:"product_library_id" binding:"required"`
}
type ProductLibraryPropertyModelUpdate struct {
	ID               string `json:"id" binding:"required"`
	Name             string `json:"name" binding:"required"`
	ProductLibraryID string `json:"product_library_id" binding:"required"`
}
type ProductLibraryPropertyModelInfo struct {
	ID string `json:"id" binding:"required"`
}
