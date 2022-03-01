package models

import (
	"gorm.io/gorm"
	"time"
)

type InventoryUnitModel struct {
	gorm.Model
	ID           string     `gorm:"column:id;primary_key" json:"id"`
	Name         string     `gorm:"column:name;default:null" json:"name"`
	CreatedBy    NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy    NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    time.Time `gorm:"column:updated_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (p *InventoryUnitModel) TableName() string {
	return "inventory_unit"
}

func (p *InventoryUnitModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *InventoryUnitModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryUnitModel) BeforeDelete(tx *gorm.DB) (err error) {
	p.UpdatedBy=p.UpdatedBy
	return
}

//strcture input
type InventoryUnitInput struct {
	Name	string	`json:"name" binding:"required"`
}
type InventoryUnitUpdate struct {
	ID	string	`json:"id" binding:"required"`
	Name	string	`json:"name" binding:"required"`
}
type InventoryUnitInfo struct {
	ID	string	`json:"id" binding:"required"`
}
//strcture output
type InventoryUnitData struct {
	ID	string	`json:"id"`
	Name	string	`json:"name"`
}
