package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductPlatformModel struct {
	ID                 string     `gorm:"column:id;primary_key" json:"id"`
	PlatformID         string     `gorm:"column:platform_id" json:"platform_id"`
	ProductModelDataID string     `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	Active             int        `gorm:"column:active" json:"active"`
	INC                int        `gorm:"column:inc" json:"inc"`
	CreatedBy          NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt          time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy          NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy          NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt 			soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductPlatformModelPreload struct {
	ID                 string `gorm:"column:id;primary_key" json:"id"`
	PlatformID         string `gorm:"column:platform_id" json:"platform_id"`
	ProductModelDataID string `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	Active             int    `gorm:"column:active" json:"active"`
	INC                int    `gorm:"column:inc" json:"inc"`
	// DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
	Platform DataPlatformModelPreload `gorm:"foreignKey:ID;references:PlatformID" json:"platform_data"`
}

// TableName sets the insert table name for this struct type
func (p *ProductPlatformModel) TableName() string {
	return "product_platform"
}
func (p *ProductPlatformModelPreload) TableName() string {
	return "product_platform"
}

func (p *ProductPlatformModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductPlatformModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductPlatformModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductPlatformModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model ProductPlatformModel
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return
}

//strcture input
type ProductPlatformModelInput struct {
	PlatformID         string `json:"platform_id" binding:"required"`
	ProductModelDataID string `json:"product_model_data_id" binding:"required"`
	Active             int    `json:"active"`
}
type ProductPlatformModelUpdate struct {
	ID            string `json:"id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	PID           string `json:"pid" binding:"required"`
	ProductTypeId string `json:"product_type_id" binding:"required"`
	Active        int    `json:"active"`
}
type ProductPlatformModellInfo struct {
	ID string `json:"id" binding:"required"`
}
