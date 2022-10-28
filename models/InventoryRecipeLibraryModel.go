package models

import (
	"gorm.io/gorm"
)

type InventoryRecipeLibraryModel struct {
	ID           						string     	`gorm:"column:id;primary_key" json:"id"`
	InventoryLibraryItem     			string	`gorm:"column:inventory_library_item;default:null" json:"nventory_library_item"`
	Inventory_library_item_child    	NullString	`gorm:"column:inventory_library_item_child;default:null" json:"inventory_library_item_child"`
	InventoryRecipeId   				NullString	`gorm:"column:inventory_recipe_id;default:null" json:"inventory_recipe_id"`
	Inc									int `gorm:"column:inc;default:null" json:"inc"`
	Library []InventoryLibraryDataModel 	`gorm:"foreignKey:ID;references:Inventory_library_item_child"`
}

type InventoryRecipeLibraryModelPreload struct {
	ID           					string     	`gorm:"column:id;primary_key" json:"id"`
	InventoryLibraryItem     		NullString		`gorm:"column:inventory_library_item;default:null" json:"nventory_library_item"`
	Inventory_library_item_child    NullString	`gorm:"column:inventory_library_item_child;default:null" json:"inventory_library_item_child"`
	InventoryRecipeId    			NullString	`gorm:"column:inventory_recipe_id;default:null" json:"inventory_recipe_id"`
	Inc								int `gorm:"column:inc;default:null" json:"inc"`
	LibraryParent 					InventoryLibraryDataModel `gorm:"foreignKey:ID;references:InventoryLibraryItem" json:"parent"`
	LibraryChild 					InventoryLibraryDataModel `gorm:"foreignKey:ID;references:Inventory_library_item_child" json:"child"`
}

// TableName sets the insert table name for this struct type
func (p *InventoryRecipeLibraryModel) TableName() string {
	return "inventory_recipe_library"
}
func (p *InventoryRecipeLibraryModelPreload) TableName() string {
	return "inventory_recipe_library"
}

func (p *InventoryRecipeLibraryModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *InventoryRecipeLibraryModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryRecipeLibraryModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryRecipeLibraryModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

//strcture input
type InventoryRecipeLibraryInput struct {
	Name	string	`json:"name" binding:"required"`
}
type InventoryRecipeLibraryUpdate struct {
	ID	string	`json:"id" binding:"required"`
	Name	string	`json:"name" binding:"required"`
}
type InventoryRecipeLibraryInfo struct {
	ID	string	`json:"id" binding:"required"`
}
//strcture output
type InventoryRecipeLibraryData struct {
	ID	string	`json:"id"`
	Name	string	`json:"name"`
}
