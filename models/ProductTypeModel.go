package golibraries/models

import (
	//"database/sql"

	"time"

	"gorm.io/gorm"
)

type ProductTypeModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`     //
	Name      NullString `gorm:"column:name" json:"name"`             //
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"` //
	CreatedAt NullString `gorm:"column:created_at" json:"created_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"` //
	UpdatedAt NullString `gorm:"column:updated_at" json:"updated_at"`
	// DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"` //
	// DeletedAt NullString `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (p *ProductTypeModel) TableName() string {
	return "product_type"
}

func (p *ProductTypeModel) BeforeCreate(tx *gorm.DB) (err error) {
	tt := time.Now()
	tFormat := tt.Format("2006-01-02 15:04:05")
	p.CreatedAt = NullStringInput(tFormat)
	// p.UpdatedAt = NullStringInput(tFormat)
	return
}

func (p *ProductTypeModel) BeforeUpdate(tx *gorm.DB) (err error) {
	tt := time.Now()
	tFormat := tt.Format("2006-01-02 15:04:05")
	p.UpdatedAt = NullStringInput(tFormat)
	return
}

// func (b *ProductTypeModel) BeforeDelete(tx *gorm.DB) (err error) {
// 	tt := time.Now()
// 	tFormat := tt.Format("2006-01-02 15:04:05")
// 	d.DeletedAt = NullStringInput(tFormat)
// 	return
// }
