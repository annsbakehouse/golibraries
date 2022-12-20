package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductLibraryNameModel struct {
	ID          string     `gorm:"column:id;primary_key" json:"id"`
	Name        NullString `gorm:"column:name;default:null" json:"name"`
	SystemName  NullString `gorm:"column:system_name" json:"system_name"`
	Description NullString `gorm:"column:description" json:"description"`
	CreatedBy   NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy   NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	// DeletedBy   NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductLibraryNameModelPreload struct {
	ID          string     `gorm:"column:id;primary_key" json:"id"`
	Name        NullString `gorm:"column:name;default:null" json:"name"`
	SystemName  NullString `gorm:"column:system_name" json:"system_name"`
	Description NullString `gorm:"column:description" json:"description"`
	// DeletedAt                soft_delete.DeletedAt              `gorm:"column:deleted_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (p *ProductLibraryNameModel) TableName() string {
	return "product_library"
}
func (p *ProductLibraryNameModelPreload) TableName() string {
	return "product_library"
}

func (p *ProductLibraryNameModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductLibraryNameModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductLibraryNameModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductLibraryNameModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model ProductLibraryTitleModel
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by", ActiveUser)
	return
}

// strcture input
type ProductLibraryNameInput struct {
	Name        string `json:"name" binding:"required"`
	SystemName  string `json:"system_name" binding:"required"`
	Description string `json:"description"`
}
type ProductLibraryNameUpdate struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	SystemName  string `json:"system_name" binding:"required"`
	Description string `json:"description"`
}
type ProductLibraryNameInfo struct {
	ID string `json:"id" binding:"required"`
}
