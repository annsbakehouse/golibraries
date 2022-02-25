package golibraries/models

import (
	"gorm.io/gorm"
	"time"
)

type DepartmentModel struct {
	ID        string     `gorm:"column:id;primary_key;autoIncrement:false" json:"id"`
	Name      NullString `gorm:"column:name" json:"name"`
	Active    uint64     `gorm:"column:active" json:"active"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	CreatedOn time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedOn time.Time `gorm:"column:updated_at" json:"updated_at"`
	DepartmentWarehouse []DepartmentWarehouseModel `gorm:"foreignKey:DepartementId;references:ID" json:"departement_warehouse"`
	// DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedOn NullString `gorm:"column:deleted_at" json:"deleted_on"`
}

type DepartmentPreloadModel struct {
	ID        string     `gorm:"column:id;primary_key;autoIncrement:false" json:"id"`
	Name      NullString `gorm:"column:name" json:"name"`
	Active    uint64     `gorm:"column:active" json:"active"`
	DepartmentWarehouse []DepartmentWarehouseModel `gorm:"foreignKey:DepartementId;references:ID" json:"departement_warehouse"`
	// DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedOn NullString `gorm:"column:deleted_at" json:"deleted_on"`
}

// TableName sets the insert table name for this struct type
func (d *DepartmentModel) TableName() string {
	return "department"
}
func (d *DepartmentPreloadModel) TableName() string {
	return "department"
}

func (d *DepartmentModel) BeforeCreate(tx *gorm.DB) (err error) {
	
	return
}

func (d *DepartmentModel) BeforeUpdate(tx *gorm.DB) (err error) {
	
	return
}

// func (d *DepartmentModel) BeforeDelete(tx *gorm.DB) (err error) {
// 	tt := time.Now()
// 	tFormat := tt.Format("2006-01-02 15:04:05")
// 	d.DeletedOn = NullStringInput(tFormat)
// 	return
// }

type DepartmentInfoForm struct {
	ID string `json:"id" binding:"required"`
}
type DepartmentDetailData struct {
	ID        string     `json:"id"`
	Name      string 	`json:"name"`
	Active    uint64     `json:"active"`
}
type DepartmentSaveForm struct {
	Name   string `json:"name" binding:"required"`
	Active uint64 `json:"active"`
	Warehouse []string `json:"warehouse"`
}
type DepartmentUpdateForm struct {
	Name   string `json:"name" binding:"required"`
	ID     string `json:"id" binding:"required"`
	Active uint64 `json:"active"`
	Warehouse []string `json:"warehouse"`
}