package models

import (
	//"database/sql"

	"time"

	"gorm.io/gorm"
)

type PromocodeModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`     //
	Name      NullString `gorm:"column:name" json:"name"`             //
	Code      string     `gorm:"column:code" json:"code"`             //
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"` //
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"` //
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	// DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"` //
	// DeletedAt NullString `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (r *PromocodeModel) TableName() string {
	return "religion"
}

func (r *PromocodeModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (r *PromocodeModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
