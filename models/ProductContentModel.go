package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductContentModelModel struct {
	ID          string     `gorm:"column:id;primary_key" json:"id"`
	Name        NullString `gorm:"column:name;default:null" json:"name"`
	ContentType int        `gorm:"column:content_type" json:"content_type"`
	Inc         int        `gorm:"column:inc" json:"inc"`
	CreatedBy   NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy   NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy   NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductContentModelPreload struct {
	ID          string     `gorm:"column:id;primary_key" json:"id"`
	Name        NullString `gorm:"column:name;default:null" json:"name"`
	ContentType int        `gorm:"column:content_type" json:"content_type"`
	// DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (p *ProductContentModelModel) TableName() string {
	return "mkt_product_content"
}
func (p *ProductContentModelPreload) TableName() string {
	return "mkt_product_content"
}

func (p *ProductContentModelModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductContentModelModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductContentModelModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductContentModelModel) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println(p.ID)
	// var model ProductContentParameterTitleModel
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return
}

//strcture input
type ProductContentModelInput struct {
	Name          string `json:"name" binding:"required"`
	OrderSequence int64  `json:"order_sequence"`
}
type ProductContentModelUpdate struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
type ProductContentModelInfo struct {
	ID string `json:"id" binding:"required"`
}
