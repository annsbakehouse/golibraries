package models

import (
	"time"

	"gorm.io/gorm"
)

type InventoryRecipeModel struct {
	ID           string     `gorm:"column:id;primary_key" json:"id"`
	Name1        NullString `gorm:"column:name_1;default:null" json:"name_1"`
	Name2        NullString `gorm:"column:name_2;default:null" json:"name_2"`
	RecipeId     NullString `gorm:"column:recipe_id;default:null" json:"recipe_id"`
	Category     int        `gorm:"column:category;default:null" json:"category"`
	UnitId       string     `gorm:"column:unit_id;default:null" json:"unit_id"`
	SyncStock    string     `gorm:"column:sync_stock;default:null" json:"sync_stock"`
	Active       int        `gorm:"column:active;default:null" json:"active"`
	Description1 NullString `gorm:"column:description_1;default:null" json:"description_1"`
	Description2 NullString `gorm:"column:description_2;default:null" json:"description_2"`
	Description3 NullString `gorm:"column:description_3;default:null" json:"description_3"`
	Image        NullString `gorm:"column:image;default:null" json:"image"`
	CreatedBy    NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy    NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedBy    NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt    soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type InventoryRecipeModelPreload struct {
	ID           string     `gorm:"column:id;primary_key" json:"id"`
	Name1        NullString `gorm:"column:name_1;default:null" json:"name_1"`
	Name2        NullString `gorm:"column:name_2;default:null" json:"name_2"`
	RecipeId     NullString `gorm:"column:recipe_id;default:null" json:"recipe_id"`
	Category     int        `gorm:"column:category;default:null" json:"category"`
	UnitId       string     `gorm:"column:unit_id;default:null" json:"unit_id"`
	SyncStock    string     `gorm:"column:sync_stock;default:null" json:"sync_stock"`
	Active       int        `gorm:"column:active;default:null" json:"active"`
	Description1 NullString `gorm:"column:description_1;default:null" json:"description_1"`
	Description2 NullString `gorm:"column:description_2;default:null" json:"description_2"`
	Description3 NullString `gorm:"column:description_3;default:null" json:"description_3"`
	Image        NullString `gorm:"column:image;default:null" json:"image"`
	// DeletedAt               soft_delete.DeletedAt                          `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
	RecipeIng               []InventoryRecipeIngredientModelPreload        `gorm:"foreignKey:InventoryRecipeId;references:ID" json:"recipe_ingredient"`
	RecipeLibrary           []InventoryRecipeLibraryModelPreload           `gorm:"foreignKey:InventoryRecipeId;references:ID" json:"recipe_library"`
	RecipeMaterialPackaging []InventoryRecipeMaterialPackagingModelPreload `gorm:"foreignKey:InventoryRecipeId;references:ID" json:"recipe_material_packaging"`
	Unit                    InventoryUnitModelPreload                      `gorm:"foreignKey:ID;references:UnitId" json:"unit"`
}

// TableName sets the insert table name for this struct type
func (p *InventoryRecipeModel) TableName() string {
	return "inventory_recipe"
}
func (p *InventoryRecipeModelPreload) TableName() string {
	return "inventory_recipe"
}

func (p *InventoryRecipeModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *InventoryRecipeModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryRecipeModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryRecipeModel) BeforeDelete(tx *gorm.DB) (err error) {
	var model InventoryRecipeModel
	tx.Model(&model).Where("id=?", p.ID).Update("deleted_by", ActiveUser)
	return
}

//strcture input
type InventoryRecipeInput struct {
	Name1        string `json:"name_1" binding:"required"`
	Name2        string `json:"name_2"`
	RecipeId     string `json:"recipe_id" binding:"required"`
	Category     int    `json:"category" binding:"required"`
	UnitId       string `json:"unit_id" binding:"required"`
	SyncStock    string `json:"sync_stock" binding:"required"`
	Active       int    `json:"active"`
	Description1 string `json:"description_1"`
	Description2 string `json:"description_2"`
	Description3 string `json:"description_3"`
	Image        string `json:"image"`

	//recipe ing
	RecipeID []string `json:"recipe"`
	Amount   []int    `json:"amount"`

	//library
	InventoryLibraryItem         []string `json:"inventory_library_item"`
	Inventory_library_item_child []string `json:"inventory_library_item_child"`

	//material packaging
	InventoryMaterialPackagingId []string `json:"inventory_material_packaging_id"`
}
type InventoryRecipeUpdate struct {
	ID           string `json:"id" binding:"required"`
	Name1        string `json:"name_1" binding:"required"`
	Name2        string `json:"name_2"`
	RecipeId     string `json:"recipe_id" binding:"required"`
	Category     int    `json:"category" binding:"required"`
	UnitId       string `json:"unit_id" binding:"required"`
	SyncStock    string `json:"sync_stock" binding:"required"`
	Active       int    `json:"active"`
	Description1 string `json:"description_1"`
	Description2 string `json:"description_2"`
	Description3 string `json:"description_3"`
	Image        string `json:"image"`

	//recipe ing
	RecipeID []string `json:"recipe"`
	Amount   []int    `json:"amount"`

	//library
	InventoryLibraryItem         []string `json:"inventory_library_item"`
	Inventory_library_item_child []string `json:"inventory_library_item_child"`

	//material packaging
	InventoryMaterialPackagingId []string `json:"inventory_material_packaging_id"`
}
type InventoryRecipeInfo struct {
	ID string `json:"id" binding:"required"`
}

//strcture output
type InventoryRecipeData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
