package models

import (
	"gorm.io/gorm"
)

type DepartmentWarehouseModel struct {
	ID        	string     				`gorm:"column:id;primary_key" json:"id"`
	WarehouseId 	NullString  				`gorm:"column:warehouse_id" json:"warehouse_id"`
	DepartementId    NullString     		`gorm:"column:departement_id" json:"departement_id"`
	Inc			int						`gorm:"column:inc" json:"inc"`
	StockOrder	int						`gorm:"column:stock_order" json:"stock_order"`
	Warehouse 	WarehousePreloadModel 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID;references:WarehouseId" json:"warehouse"`
	Department DepartmentPreloadModel 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID;references:DepartementId" json:"department"`
}

// TableName sets the insert table name for this struct type
func (d *DepartmentWarehouseModel) TableName() string {
	return "departement_warehouse"
}

func (d *DepartmentWarehouseModel) BeforeCreate(tx *gorm.DB) (err error) {
	
	return
}

func (d *DepartmentWarehouseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	
	return
}
func (d *DepartmentWarehouseModel) AfterUpdate(tx *gorm.DB) (err error) {
	
	return
}