package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductLibraryCategoryModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	Name      NullString `gorm:"column:name;default:null" json:"name"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductLibraryCategoryModelPreload struct {
	ID   string     `gorm:"column:id;primary_key" json:"id"`
	Name NullString `gorm:"column:name;default:null" json:"name"`
	// DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (p *ProductLibraryCategoryModel) TableName() string {
	return "product_library_category"
}
func (p *ProductLibraryCategoryModelPreload) TableName() string {
	return "product_library_category"
}

func (p *ProductLibraryCategoryModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductLibraryCategoryModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductLibraryCategoryModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductLibraryCategoryModel) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println("Before Delete")
	// var model ProductLibraryCategoryModel
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by", ActiveUser)
	return
}

//strcture input
type ProductLibraryCategoryModelInput struct {
	Name string `json:"name" binding:"required"`
}
type ProductLibraryCategoryModelUpdate struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
type ProductLibraryCategoryModelInfo struct {
	ID string `json:"id" binding:"required"`
}
