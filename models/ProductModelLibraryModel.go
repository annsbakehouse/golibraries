package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductModelLibraryModel struct {
	ID                       string     `gorm:"column:id;primary_key" json:"id"`
	ProductLibraryPropertyID string     `gorm:"column:product_library_property_id" json:"product_library_property_id"`
	CustomerLevelID          NullString `gorm:"column:customer_level_id" json:"customer_level_id"`
	ProductModelID           NullString `gorm:"column:product_model_id" json:"product_model_id"`
	ProductID                NullString `gorm:"column:product_id" json:"product_id"`
	ProductLibraryTitleID    NullString `gorm:"column:product_library_title_id" json:"product_library_title_id"`
	ProductLibraryCategoryID NullString `gorm:"column:product_library_category_id" json:"product_library_category_id"`
	IsRequired               int        `gorm:"column:is_required" json:"is_required"`
	LinkTo                   int        `gorm:"column:link_to" json:"link_to"`
	CreatedBy                NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt                time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy                NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt                time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy                NullString `gorm:"column:deleted_by" json:"deleted_by"`
}

type ProductModelLibraryModelPreload struct {
	ID                 string  `gorm:"column:id;primary_key" json:"id"`
	ProductModelDataID string  `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	WarehouseType      int     `gorm:"column:warehouse_type" json:"warehouse_type"`
	WarehouseFrom      string  `gorm:"column:warehouse_from" json:"warehouse_from"`
	WarehouseTo        string  `gorm:"column:warehouse_to" json:"warehouse_to"`
	ProductionTime     float64 `gorm:"column:production_time" json:"production_time"`
}

// TableName sets the insert table name for this struct type
func (p *ProductModelLibraryModel) TableName() string {
	return "product_data_library"
}
func (p *ProductModelLibraryModelPreload) TableName() string {
	return "product_data_library"
}

func (p *ProductModelLibraryModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductModelLibraryModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductModelLibraryModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

//strcture input
type ProductModelLibraryInput struct {
	ProductModelID string                    `json:"product_model_id" binding:"required"`
	LibraryData    []ProductModelLibraryForm `json:"library_data" binding:"required,dive"`
}
type ProductModelLibraryForm struct {
	ID                       string `json:"id"`
	ProductLibraryPropertyID string `json:"product_library_property_id" binding:"required"`
	CustomerLevelID          string `json:"customer_level_id"`
	ProductID                string `json:"product_id"`
	ProductLibraryTitleID    string `json:"product_library_title_id"  binding:"required"`
	ProductLibraryCategoryID string `json:"product_library_category_id" binding:"required"`
	IsRequired               int    `json:"is_required" binding:"exists"`
	LinkTo                   int    `json:"link_to"`
}

type ProductModelLibraryInfo struct {
	ID string `json:"id" binding:"required"`
}
