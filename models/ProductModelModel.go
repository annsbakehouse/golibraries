package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductModelModel struct {
	ID            string     `gorm:"column:id;primary_key" json:"id"`
	Name          NullString `gorm:"column:name;default:null" json:"name"`
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

type ProductModelModelPreload struct {
	ID            string           `gorm:"column:id;primary_key" json:"id"`
	Name          NullString       `gorm:"column:name;default:null" json:"name"`
	ProductTypeId NullString       `gorm:"column:product_type_id" json:"product_type_id"`
	Active        int              `gorm:"column:active" json:"active"`
	ProductType   ProductTypeModel `gorm:"foreignKey:ID;references:ProductTypeId" json:"product_type"`
	// DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (p *ProductModelModel) TableName() string {
	return "product_model"
}
func (p *ProductModelModelPreload) TableName() string {
	return "product_model"
}

func (p *ProductModelModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductModelModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductModelModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductModelModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model ProductModelModel
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by", ActiveUser)
	return
}

//strcture input
type ProductModelModelInput struct {
	Name          string   `json:"name" binding:"required"`
	ProductTypeId string   `json:"product_type_id" binding:"required"`
	Active        int      `json:"active"`
	CountryID     []string `json:"country_id"`
	Word          []string `json:"word"`
}
type ProductModelModelUpdate struct {
	ID            string   `json:"id" binding:"required"`
	Name          string   `json:"name" binding:"required"`
	ProductTypeId string   `json:"product_type_id" binding:"required"`
	Active        int      `json:"active"`
	CountryID     []string `json:"country_id"`
	Word          []string `json:"word"`
}
type ProductModelModelInfo struct {
	ID string `json:"id" binding:"required"`
}
