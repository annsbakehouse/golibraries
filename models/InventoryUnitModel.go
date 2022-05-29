package models

import (
	"time"

	"gorm.io/gorm"
)

type InventoryUnitModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	Name      string     `gorm:"column:name;default:null" json:"name"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type InventoryUnitModelPreload struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	Name      string     `gorm:"column:name;default:null" json:"name"`
	CreatedBy NullString `gorm:"column:created_by" json:"-"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedBy NullString `gorm:"column:updated_by" json:"-"`
	// DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (p *InventoryUnitModel) TableName() string {
	return "inventory_unit"
}
func (p *InventoryUnitModelPreload) TableName() string {
	return "inventory_unit"
}

func (p *InventoryUnitModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *InventoryUnitModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryUnitModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryUnitModel) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println("Before Delete")
	// var inventoryUnitModel InventoryUnitModel
	// tx.Model(&inventoryUnitModel).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return
}

//strcture input
type InventoryUnitInput struct {
	Name string `json:"name" binding:"required"`
}
type InventoryUnitUpdate struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
type InventoryUnitInfo struct {
	ID string `json:"id" binding:"required"`
}

//strcture output
type InventoryUnitData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
