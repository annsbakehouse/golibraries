package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductImageUsageModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	Name      NullString `gorm:"column:name;default:null" json:"name"`
	PUID      string     `gorm:"column:piuid;unique" json:"piuid"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductImageUsageModelPreload struct {
	ID   string     `gorm:"column:id;primary_key" json:"id"`
	Name NullString `gorm:"column:name;default:null" json:"name"`
	PUID string     `gorm:"column:piuid" json:"piuid"`
	// DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (p *ProductImageUsageModel) TableName() string {
	return "product_image_usage"
}
func (p *ProductImageUsageModelPreload) TableName() string {
	return "product_image_usage"
}

func (p *ProductImageUsageModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductImageUsageModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductImageUsageModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductImageUsageModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model ProductImageUsageModel
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by", ActiveUser)
	return
}
func (p *ProductImageUsageModel) AfterDelete(tx *gorm.DB) (err error) {

	return
}

//strcture input
type ProductImageUsageModelInput struct {
	Name string `json:"name" binding:"required"`
	PUID string `json:"piuid" binding:"required"`
}
type ProductImageUsageModelUpdate struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	PUID string `json:"piuid"`
}
type ProductImageUsageModelInfo struct {
	ID string `json:"id" binding:"required"`
}
