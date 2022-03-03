package models

import (
	"gorm.io/gorm"
	"time"
	"gorm.io/plugin/soft_delete"
)

type InventoryMaterialPackagingModel struct {
	ID           string     `gorm:"column:id;primary_key" json:"id"`
	PackagingName1         NullString     `gorm:"column:packaging_name_1;default:null" json:"packaging_name_1"`
	PackagingName2         NullString     `gorm:"column:packaging_name_2;default:null" json:"packaging_name_2"`
	InternalID	 		   NullString		`gorm:"column:internal_id" json:"internal_id"`
	ContentConverter	 	int		`gorm:"column:content_converter" json:"content_converter"`
	InventoryUnitId	 	 string		 	`gorm:"column:inventory_unit_id" json:"inventory_unit_id"`
	Active	 	int		`gorm:"column:active" json:"active"`
	CreatedBy    NullString  	`gorm:"column:created_by" json:"created_by"`
	CreatedAt    time.Time 		`gorm:"column:created_at" json:"crated_at"`
	UpdatedBy    NullString 	`gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy	 NullString `gorm:"column:updated_by" json:"deleted_by"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
	Unit  InventoryUnitModelPreload `gorm:"foreignKey:InventoryUnitId;references:ID" json:"unit"`
}

type InventoryMaterialPackagingPreloadModel struct {
	ID           string     `gorm:"column:id;primary_key" json:"id"`
	PackagingName1         string     `gorm:"column:packaging_name_1;default:null" json:"packaging_name_1"`
	PackagingName2         string     `gorm:"column:packaging_name_2;default:null" json:"packaging_name_2"`
	InternalID	 		   string		`gorm:"column:internal_id" json:"internal_id"`
	ContentConverter	 	int		`gorm:"column:content_converter" json:"content_converter"`
	InventoryUnitId	 	 string		 	`gorm:"column:inventory_unit_id" json:"inventory_unit_id"`
	Active	 	int		`gorm:"column:active" json:"active"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
	Unit  InventoryUnitModelPreload `gorm:"foreignKey:InventoryUnitId;references:ID" json:"unit"`
}



// TableName sets the insert table name for this struct type
func (p *InventoryMaterialPackagingModel) TableName() string {
	return "inventory_material_packaging"
}
func (p *InventoryMaterialPackagingPreloadModel) TableName() string {
	return "inventory_material_packaging"
}

func (p *InventoryMaterialPackagingModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *InventoryMaterialPackagingModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryMaterialPackagingModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryMaterialPackagingModel	) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println("Before Delete")
	_,con,_ := DbConnect()
	var model InventoryMaterialPackagingModel
	con.Model(&model).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return
}

//strcture input
type InventoryMaterialPackagingModelInput struct {
	PackagingName1         string     	`json:"packaging_name_1" binding:"required"`
	PackagingName2         string     	`json:"packaging_name_2"`
	InternalID	 		   string		`json:"internal_id" binding:"required"`
	ContentConverter	   int			`json:"content_converter" binding:"required"`
	InventoryUnitId	 	   string		`json:"inventory_unit_id" binding:"required"`
	Active	 			   int			`json:"active"`
}
type InventoryMaterialPackagingModelUpdate struct {
	ID	string	`json:"id" binding:"required"`
	PackagingName1         string     	`json:"packaging_name_1" binding:"required"`
	PackagingName2         string     	`json:"packaging_name_2"`
	InternalID	 		   string		`json:"internal_id" binding:"required"`
	ContentConverter	   int			`json:"content_converter" binding:"required"`
	InventoryUnitId	 	   string		`json:"inventory_unit_id" binding:"required"`
	Active	 			   int			`json:"active"`
}
type InventoryMaterialPackagingModelInfo struct {
	ID	string	`json:"id" binding:"required"`
}


