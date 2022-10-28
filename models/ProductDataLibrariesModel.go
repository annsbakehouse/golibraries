package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductDataLibrariesModel struct {
	ID                         string     `gorm:"column:id;primary_key" json:"id"`
	ProductLibraryPropertiesID NullString `gorm:"column:product_library_properties_id" json:"product_library_properties_id"`
	CustomerLevelID            NullString `gorm:"column:customer_level_id" json:"customer_level_id"`
	ProductModelDataID         NullString `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	ProductID                  NullString `gorm:"column:product_id" json:"product_id"`
	IsRequired                 int        `gorm:"column:is_required" json:"is_required"`
	LinkTo                     NullInt64  `gorm:"column:link_to" json:"link_to"`
	ProductLibraryID           NullString `gorm:"column:product_library_id" json:"product_library_id"`
	ProductModelID             NullString `gorm:"column:product_model_id" json:"product_model_id"`
	CreatedBy                  NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt                  time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy                  NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt                  time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
}

type ProductDataLibrariesModelPreload struct {
	ID                           string                             `gorm:"column:id;primary_key" json:"id"`
	ProductLibraryPropertiesID   NullString                         `gorm:"column:product_library_properties_id" json:"product_library_properties_id"`
	CustomerLevelID              NullString                         `gorm:"column:customer_level_id" json:"customer_level_id"`
	ProductModelDataID           NullString                         `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	ProductID                    NullString                         `gorm:"column:product_id" json:"product_id"`
	IsRequired                   int                                `gorm:"column:is_required" json:"is_required"`
	LinkTo                       int                                `gorm:"column:link_to" json:"link_to"`
	ProductLibraryID             NullString                         `gorm:"column:product_library_id" json:"product_library_id"`
	ProductModelID               NullString                         `gorm:"column:product_model_id" json:"product_model_id"`
	ProductLibraryPropertiesInfo ProductLibraryPropertyModelPreload `gorm:"foreignKey:ProductLibraryPropertiesID;references:ID" json:"product_model_library_properties"`
	ProductModelInfo             ProductModelModelPreload           `gorm:"foreignKey:ProductModelID;references:ID" json:"product_model_info"`
	ProductInfo                  ProductModelPreload                `gorm:"foreignKey:ProductID;references:ID" json:"product_info"`
}

// TableName sets the insert table name for this struct type
func (p *ProductDataLibrariesModel) TableName() string {
	return "product_data_libraries"
}
func (p *ProductDataLibrariesModelPreload) TableName() string {
	return "product_data_libraries"
}

func (p *ProductDataLibrariesModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductDataLibrariesModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductDataLibrariesModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

//strcture input
type ProductDataLibrariesModelInput struct {
	ProductModelDataID         string `json:"product_model_data_id"`
	ProductLibraryID           string `json:"product_library_id" binding:"required"`
	ProductLibraryPropertiesID string `json:"product_library_properties_id" binding:"required"`
	CustomerLevelID            string `json:"customer_level_id" binding:"required"`
	ProductID                  string `json:"product_id"`
	ProductModelID             string `json:"product_model_id"`
	IsRequired                 int    `json:"is_required"`
	LinkTo                     int64  `json:"link_to" binding:"required"`
}
type ProductDataLibrariesModelUpdate struct {
	ID                         string `json:"id" binding:"required"`
	ProductModelDataID         string `json:"product_model_data_id"`
	ProductLibraryID           string `json:"product_library_id" binding:"required"`
	ProductLibraryPropertiesID string `json:"product_library_properties_id" binding:"required"`
	CustomerLevelID            string `json:"customer_level_id" binding:"required"`
	ProductID                  string `json:"product_id"`
	ProductModelID             string `json:"product_model_id"`
	IsRequired                 int    `json:"is_required"`
	LinkTo                     int64  `json:"link_to" binding:"required"`
}

type ProductDataLibrariesModelInfo struct {
	ID string `json:"id" binding:"required"`
}
type ProductDataLibrariesModellist struct {
	ProductModelDataID string `json:"product_model_data_id" binding:"required"`
}
