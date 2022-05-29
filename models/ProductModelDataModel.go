package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductModelDataModel struct {
	ID                string     `gorm:"column:id;primary_key" json:"id"`
	ProductTypeId     string     `gorm:"column:product_type_id" json:"product_type_id"`
	ProductID         string     `gorm:"column:product_id" json:"product_id"`
	ProductModelID    string     `gorm:"column:product_model_id" json:"product_model_id"`
	InventoryRecipeID string     `gorm:"column:inventory_recipe_id" json:"inventory_recipe_id"`
	MOQ               int        `gorm:"column:moq" json:"moq"`
	Mid               string     `gorm:"column:mid;uniqueIndex" json:"mid"`
	MOQINC            int        `gorm:"column:moq_inc" json:"moq_inc"`
	Active            int        `gorm:"column:active" json:"active"`
	INC               int        `gorm:"column:inc" json:"inc"`
	CreatedBy         NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt         time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy         NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt         time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy         NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt         soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductModelDataPreload struct {
	ID                string `gorm:"column:id;primary_key" json:"id"`
	ProductTypeId     string `gorm:"column:product_type_id" json:"product_type_id"`
	ProductId         string `gorm:"column:product_id" json:"product_id"`
	ProductModelId    string `gorm:"column:product_model_id" json:"product_model_id"`
	InventoryRecipeId string `gorm:"column:inventory_recipe_id" json:"inventory_recipe_id"`
	MOQ               int    `gorm:"column:moq" json:"moq"`
	Mid               string `gorm:"column:mid;uniqueIndex" json:"mid"`
	MoqInc            int    `gorm:"column:moq_inc" json:"moq_inc"`
	Active            int    `gorm:"column:active" json:"active"`
	INC               int    `gorm:"column:inc" json:"inc"`
	// DeletedAt              soft_delete.DeletedAt               `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
	ModelData              ProductModelModelPreload            `gorm:"foreignKey:ID;references:ProductModelId" json:"model_data"`
	PlatformData           []ProductPlatformModelPreload       `gorm:"foreignKey:ProductModelDataID;references:ID" json:"platform"`
	Pricing                []ProductPricingModelPreload        `gorm:"foreignKey:ProductModelDataID;references:ID" json:"pricing"`
	ProductModelImage      []ProductModelImageModelPreload     `gorm:"foreignKey:ProductModelDataID;references:ID" json:"product_model_image"`
	InventoryRecipeInfo    InventoryRecipeModelPreload         `gorm:"foreignKey:InventoryRecipeId;references:ID" json:"inventory_recipe_info"`
	DeliveryInfo           ProductModelDeliveryPreload         `gorm:"foreignKey:ProductModelDataID;references:ID" json:"delivery_info"`
	ModelWarehouseInfoW4W2 []ProductModelWarehouseModelPreload `gorm:"foreignKey:ProductModelDataID;references:ID" json:"warehouse_info_w4w2"`
	ModelWarehouseInfoW3W2 []ProductModelWarehouseModelPreload `gorm:"foreignKey:ProductModelDataID;references:ID" json:"warehouse_info_w3w2"`
}

// TableName sets the insert table name for this struct type
func (p *ProductModelDataModel) TableName() string {
	return "product_model_data"
}
func (p *ProductModelDataPreload) TableName() string {
	return "product_model_data"
}

func (p *ProductModelDataModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductModelDataModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductModelDataModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductModelDataModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model ProductModelDataModel
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by", ActiveUser)
	return
}

//strcture input
type ProductModelDataModelInput struct {
	ProductTypeId     string   `json:"product_type_id" binding:"required"`
	ProductID         string   `json:"product_id" binding:"required"`
	ProductModelID    string   `json:"product_model_id" binding:"required"`
	Mid               string   `json:"mid" binding:"required"`
	InventoryRecipeID string   `json:"inventory_recipe_id" binding:"required"`
	MOQ               int      `json:"moq"`
	MOQINC            int      `json:"moq_inc"`
	Active            int      `json:"active"`
	Platform          []string `json:"platform"`
}
type ProductModelDataModelUpdate struct {
	ID                string   `json:"id" binding:"required"`
	InventoryRecipeID string   `json:"inventory_recipe_id" binding:"required"`
	MOQ               int      `json:"moq"`
	MOQINC            int      `json:"moq_inc"`
	Platform          []string `json:"platform"`
}
type ProductModelDataModelInfo struct {
	ID string `json:"id" binding:"required"`
}
