package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductPlatformDisplayModel struct {
	ID                 string     `gorm:"column:id;primary_key" json:"id"`
	ProductID          string     `gorm:"column:product_id" json:"product_id"`
	ProductModelID     NullString `gorm:"column:product_model_id" json:"product_model_id"`
	ProductModelDataID NullString `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	PlatformID         string     `gorm:"column:platform_id" json:"platform_id"`
	DisplayFrom        NullString `gorm:"column:display_from" json:"display_from"`
	DisplayUntil       NullString `gorm:"column:display_until" json:"display_until"`
	DeliveryFrom       NullString `gorm:"column:delivery_from" json:"delivery_from"`
	DeliveryUntil      NullString `gorm:"column:delivery_until" json:"delivery_until"`
	DisableFrom        NullString `gorm:"column:disable_from" json:"disable_from"`
	DisableUntil       NullString `gorm:"column:disable_until" json:"disable_until"`
	CreatedBy          NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt          time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy          NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" json:"updated_at"`
	// DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt NullString `gorm:"column:deleted_at" json:"deleted_at"`
}

type ProductPlatformDisplayModelPreload struct {
	ID                 string                   `gorm:"column:id;primary_key" json:"id"`
	ProductID          string                   `gorm:"column:product_id" json:"product_id"`
	ProductModelID     NullString               `gorm:"column:product_model_id" json:"product_model_id"`
	ProductModelDataID NullString               `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	PlatformID         string                   `gorm:"column:platform_id" json:"platform_id"`
	DisplayFrom        NullString               `gorm:"column:display_from" json:"display_from"`
	DisplayUntil       NullString               `gorm:"column:display_until" json:"display_until"`
	DeliveryFrom       NullString               `gorm:"column:delivery_from" json:"delivery_from"`
	DeliveryUntil      NullString               `gorm:"column:delivery_until" json:"delivery_until"`
	DisableFrom        NullString               `gorm:"column:disable_from" json:"disable_from"`
	DisableUntil       NullString               `gorm:"column:disable_until" json:"disable_until"`
	PlatformInfo       DataPlatformModelPreload `gorm:"foreignKey:ID;references:PlatformID" json:"platform_info"`
}

// TableName sets the insert table name for this struct type
func (d *ProductPlatformDisplayModel) TableName() string {
	return "mkt_product_platform_display"
}

func (d *ProductPlatformDisplayModelPreload) TableName() string {
	return "mkt_product_platform_display"
}

func (d *ProductPlatformDisplayModel) BeforeCreate(tx *gorm.DB) (err error) {

	return
}

func (d *ProductPlatformDisplayModel) BeforeUpdate(tx *gorm.DB) (err error) {

	return
}

type ProductPlatformDisplaySave struct {
	ProductID       string                            `json:"product_id" binding:"required"`
	ProductPlatform []ProductPlatformDisplayModelSave `json:"product_platform" binding:"required,dive"`
}
type ProductPlatformDisplayInfo struct {
	ProductID string `json:"product_id" binding:"required"`
}
type ProductModelPlatformDisplayInfo struct {
	ProductModelDataID string `json:"product_model_data_id" binding:"required"`
}
type ProductPlatformDisplayModelSave struct {
	ID                 string `json:"id"`
	ProductModelID     string `json:"product_model_id"`
	ProductModelDataID string `json:"product_model_data_id"`
	PlatformID         string `json:"platform_id" binding:"required"`
	DisplayFrom        string `json:"display_from"`
	DisplayUntil       string `json:"display_until"`
	DeliveryFrom       string `json:"delivery_from"`
	DeliveryUntil      string `json:"delivery_until"`
	DisableFrom        string `json:"disable_from"`
	DisableUntil       string `json:"disable_until"`
}
type ProductModelPlatformDisplaySave struct {
	ProductID          string                                 `json:"product_id" binding:"required"`
	ProductModelID     string                                 `json:"product_model_id"  binding:"required"`
	ProductModelDataID string                                 `json:"product_model_data_id" binding:"required"`
	ProductPlatform    []ProductModelPlatformDisplayModelSave `json:"product_platform" binding:"required,dive"`
}
type ProductModelPlatformDisplayModelSave struct {
	ID            string `json:"id"`
	PlatformID    string `json:"platform_id" binding:"required"`
	DisplayFrom   string `json:"display_from"`
	DisplayUntil  string `json:"display_until"`
	DeliveryFrom  string `json:"delivery_from"`
	DeliveryUntil string `json:"delivery_until"`
	DisableFrom   string `json:"disable_from"`
	DisableUntil  string `json:"disable_until"`
}

// func (d *DistrictModel) BeforeDelete(tx *gorm.DB) (err error) {
// 	tt := time.Now()
// 	tFormat := tt.Format("2006-01-02 15:04:05")
// 	d.DeletedAt = NullStringInput(tFormat)
// 	return
// }
