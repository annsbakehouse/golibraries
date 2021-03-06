package models

import (
	"time"

	"gorm.io/gorm"
	// "fmt"
	// library "github.com/annsbakehouse/golibraries/library"
)

type InventoryLibraryModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	Name      string     `gorm:"column:name;default:null" json:"name"`
	IsParent  int        `gorm:"column:is_parent" json:"is_parent"`
	ParentID  NullString `gorm:"column:parent_id" json:"parent_id"`
	Active    int        `gorm:"column:active" json:"active"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"_"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"`
	//DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
	Child []InventoryLibraryModel `gorm:"foreignKey:ParentID;references:ID;constraint:OnDelete:CASCADE" json:"child"`
}

type InventoryLibraryDataModelPreload struct {
	ID       string     `gorm:"column:id;primary_key" json:"id"`
	Name     string     `gorm:"column:name;default:null" json:"name"`
	IsParent int        `gorm:"column:is_parent" json:"is_parent"`
	ParentID NullString `gorm:"column:parent_id" json:"parent_id"`
	Active   int        `gorm:"column:active" json:"active"`
	//DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at" json:"-"`
}

type InventoryLibraryDataModel struct {
	ID       string     `gorm:"column:id;primary_key" json:"id"`
	Name     string     `gorm:"column:name;default:null" json:"name"`
	IsParent int        `gorm:"column:is_parent" json:"is_parent"`
	ParentID NullString `gorm:"column:parent_id" json:"parent_id"`
	Active   int        `gorm:"column:active" json:"active"`
	//DeletedAt soft_delete.DeletedAt              `gorm:"column:deleted_at" json:"-"`
	Child  []InventoryLibraryDataModelPreload `gorm:"foreignKey:ParentID;references:ID;Constraint:OnDelete:CASCADE" json:"child"`
	Parent InventoryLibraryDataModelPreload   `gorm:"foreignKey:ID;references:ParentID" json:"parent"`
}

// TableName sets the insert table name for this struct type
func (p *InventoryLibraryModel) TableName() string {
	return "inventory_library"
}
func (p *InventoryLibraryDataModel) TableName() string {
	return "inventory_library"
}
func (p *InventoryLibraryDataModelPreload) TableName() string {
	return "inventory_library"
}

func (p *InventoryLibraryModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *InventoryLibraryModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryLibraryModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryLibraryModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model InventoryLibraryModel
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by", ActiveUser)
	// tx.Model(&model).Where("parent_id=?", p.ID).Update("deleted_by", ActiveUser)
	return
}
func (p *InventoryLibraryModel) AfterDelete(tx *gorm.DB) (err error) {
	return
}

//strcture input
type InventoryLibraryInput struct {
	Name     string `json:"name" binding:"required"`
	IsParent int    `json:"is_parent"`
	ParentID string `json:"parent_id"`
	Active   int    `json:"active"`
}
type InventoryLibraryUpdate struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	IsParent int    `json:"is_parent"`
	ParentID string `json:"parent_id"`
	Active   int    `json:"active"`
}
type InventoryLibraryInfo struct {
	ID string `json:"id" binding:"required"`
}

type InventoryLibraryData struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	IsParent int        `json:"is_parent"`
	ParentID NullString `json:"parent_id"`
	Active   int        `json:"active"`
	//DeletedAt soft_delete.DeletedAt       `gorm:`
	Child []InventoryLibraryDataModel `gorm:"foreignKey:ParentID;references:ID" json:"child"`
}
