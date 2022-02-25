package golibraries/models

import (
	"gorm.io/gorm"
	"time"
)

type DataPlatformModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`      //
	Name      string     `gorm:"column:name;default:null" json:"name"` //
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`  //
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"` //
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	// DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"` //
	// DeletedAt NullString `gorm:"column:deleted_at" json:"deleted_at"`
}

// info dataplatform
type DataPlatformInfoForm struct {
	ID string `json:"id" binding:"required"`
}
type DataPlatformDetailData struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type DataPlatformSaveForm struct {
	Name string `json:"name" binding:"required"`
}
type DataPlatformUpdateForm struct {
	Name string `json:"name" binding:"required"`
	ID   string `json:"id" binding:"required"`
}
// TableName sets the insert table name for this struct type
func (d *DataPlatformModel) TableName() string {
	return "platform"
}

func (d *DataPlatformModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (d *DataPlatformModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

// func (d *DataPlatformModel) BeforeDelete(tx *gorm.DB) (err error) {
// 	tt := time.Now()
// 	tFormat := tt.Format("2006-01-02 15:04:05")
// 	d.DeletedAt = NullStringInput(tFormat)
// 	return
// }
