package models

import (
	"gorm.io/gorm"
)

type InventoryRecipeMaterialPackagingModel struct {
	ID                           string                                 `gorm:"column:id;primary_key" json:"id"`
	InventoryMaterialPackagingId string                                 `gorm:"column:inventory_material_packaging_id;default:null" json:"inventory_material_packaging_id"`
	InventoryRecipeId            NullString                             `gorm:"column:inventory_recipe_id;default:null" json:"inventory_recipe_id"`
	Inc                          int                                    `gorm:"column:inc;default:null" json:"inc"`
	MaterialPackaging            InventoryMaterialPackagingPreloadModel `gorm:"foreignKey:ID;references:InventoryMaterialPackagingId"`
}

type InventoryRecipeMaterialPackagingModelPreload struct {
	ID                           string                                 `gorm:"column:id;primary_key" json:"id"`
	InventoryMaterialPackagingId string                                 `gorm:"column:inventory_material_packaging_id;default:null" json:"inventory_material_packaging_id"`
	InventoryRecipeId            NullString                             `gorm:"column:inventory_recipe_id;default:null" json:"inventory_recipe_id"`
	Inc                          int                                    `gorm:"column:inc;default:null" json:"inc"`
	MaterialPackaging            InventoryMaterialPackagingPreloadModel `gorm:"foreignKey:ID;references:InventoryMaterialPackagingId" json:"material_packaging"`
}

// TableName sets the insert table name for this struct type
func (p *InventoryRecipeMaterialPackagingModel) TableName() string {
	return "inventory_recipe_material_packaging"
}
func (p *InventoryRecipeMaterialPackagingModelPreload) TableName() string {
	return "inventory_recipe_material_packaging"
}

func (p *InventoryRecipeMaterialPackagingModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *InventoryRecipeMaterialPackagingModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryRecipeMaterialPackagingModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryRecipeMaterialPackagingModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

//strcture input
type InventoryRecipeMaterialPackagingInput struct {
	Name string `json:"name" binding:"required"`
}
type InventoryRecipeMaterialPackagingUpdate struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
type InventoryRecipeMaterialPackagingInfo struct {
	ID string `json:"id" binding:"required"`
}

//strcture output
type InventoryRecipeMaterialPackagingData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
