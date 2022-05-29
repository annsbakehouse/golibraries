package models

import (
	"time"

	"gorm.io/gorm"
)

type CourierPlatformModel struct {
	ID            string     `gorm:"column:id;primary_key" json:"id"`
	PlatformID    string     `gorm:"column:platform_id" json:"platform_id"`
	CourierTypeID string     `gorm:"column:courier_type_id" json:"courier_type_id"`
	INC           int        `gorm:"column:inc;default:1" json:"inc"`
	CreatedBy     NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt     time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy     NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt     time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedBy     NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt     soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type CourierPlatformModelPreload struct {
	ID            string                   `gorm:"column:id;primary_key" json:"id"`
	PlatformID    string                   `gorm:"column:platform_id" json:"platform_id"`
	CourierTypeID string                   `gorm:"column:courier_type_id" json:"courier_type_id"`
	INC           int                      `gorm:"column:inc;default:1" json:"inc"`
	DataPlatform  DataPlatformModelPreload `gorm:"foreignKey:ID;references:PlatformID" json:"data_platform"`
	// DeletedAt     soft_delete.DeletedAt    `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (c *CourierPlatformModel) TableName() string {
	return "courier_platform"
}
func (c *CourierPlatformModelPreload) TableName() string {
	return "courier_platform"
}

func (c *CourierPlatformModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *CourierPlatformModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *CourierPlatformModel) BeforeDelete(tx *gorm.DB) (err error) {
	//var model CourierPlatformModel
	//tx.Model(&model).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return
}
