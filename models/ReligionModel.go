package golibraries/models

import (
	//"database/sql"

	"time"

	"gorm.io/gorm"
)

type ReligionModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`     //
	Name      NullString `gorm:"column:name" json:"name"`             //
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"` //
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"` //
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	// DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"` //
	// DeletedAt NullString `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (r *ReligionModel) TableName() string {
	return "religion"
}

func (r *ReligionModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (r *ReligionModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

// func (b *ReligionModel) BeforeDelete(tx *gorm.DB) (err error) {
// 	tt := time.Now()
// 	tFormat := tt.Format("2006-01-02 15:04:05")
// 	d.DeletedAt = NullStringInput(tFormat)
// 	return
// }
