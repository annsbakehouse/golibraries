package models

import (
	"gorm.io/gorm"
)

type InventoryRecipeIngredientModel struct {
	ID                string                      `gorm:"column:id;primary_key" json:"id"`
	InventoryRecipeId string                      `gorm:"column:inventory_recipe_id;default:null" json:"inventory_recipe_id"`
	RecipeId          string                      `gorm:"column:recipe_id;default:null" json:"recipe_id"`
	Amount            float32                     `gorm:"column:amount;default:null" json:"amount"`
	Inc               int                         `gorm:"column:inc;default:null" json:"inc"`
	Recipe            InventoryRecipeModelPreload `gorm:"foreignKey:ID;references:RecipeId"`
}

type InventoryRecipeIngredientModelPreload struct {
	ID                string                      `gorm:"column:id;primary_key" json:"id"`
	InventoryRecipeId string                      `gorm:"column:inventory_recipe_id;default:null" json:"inventory_recipe_id"`
	RecipeId          string                      `gorm:"column:recipe_id;default:null" json:"recipe_id"`
	Amount            float32                     `gorm:"column:amount;default:null" json:"amount"`
	Inc               int                         `gorm:"column:inc;default:null" json:"inc"`
	Recipe            InventoryRecipeModelPreload `gorm:"foreignKey:ID;references:RecipeId"`
}

// TableName sets the insert table name for this struct type
func (p *InventoryRecipeIngredientModel) TableName() string {
	return "inventory_recipe_ingredient"
}
func (p *InventoryRecipeIngredientModelPreload) TableName() string {
	return "inventory_recipe_ingredient"
}

func (p *InventoryRecipeIngredientModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *InventoryRecipeIngredientModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryRecipeIngredientModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryRecipeIngredientModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

//strcture input
type InventoryRecipeIngredientInput struct {
	Name string `json:"name" binding:"required"`
}
type InventoryRecipeIngredientUpdate struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
type InventoryRecipeIngredientInfo struct {
	ID string `json:"id" binding:"required"`
}

//strcture output
type InventoryRecipeIngredientData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
