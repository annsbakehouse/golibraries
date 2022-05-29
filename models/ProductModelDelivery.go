package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductModelDeliveryModel struct {
	ID                 string     `gorm:"column:id;primary_key" json:"id"`
	ProductModelDataID string     `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	CourierTypeID      string     `gorm:"column:courier_type_id" json:"courier_type_id"`
	CourierSubtypeID   string     `gorm:"column:courier_subtype_id" json:"courier_subtype_id"`
	PackagingFee       int        `gorm:"column:packaging_fee" json:"packaging_fee"`
	Weight             int        `gorm:"column:weight" json:"weight"`
	IsSpRequired       int        `gorm:"column:is_sp_required" json:"is_sp_required"`
	CreatedBy          NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt          time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy          NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy          NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt          soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductModelDeliveryPreload struct {
	ID                 string                     `gorm:"column:id;primary_key" json:"id"`
	ProductModelDataID string                     `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	CourierTypeID      string                     `gorm:"column:courier_type_id" json:"courier_type_id"`
	CourierSubtypeID   string                     `gorm:"column:courier_subtype_id" json:"courier_subtype_id"`
	PackagingFee       int                        `gorm:"column:packaging_fee" json:"packaging_fee"`
	Weight             int                        `gorm:"column:weight" json:"weight"`
	IsSpRequired       int                        `gorm:"column:is_sp_required" json:"is_sp_required"`
	CourierSubTypeInfo CourierSubTypeModelPreload `gorm:"foreignKey:ID;references:CourierSubtypeID" json:"courier_subtype_info"`
	CourierTypeInfo    CourierTypeModelPreload    `gorm:"foreignKey:ID;references:CourierTypeID" json:"courier_type_info"`
	// DeletedAt          soft_delete.DeletedAt      `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (p *ProductModelDeliveryModel) TableName() string {
	return "product_delivery_setting"
}
func (p *ProductModelDeliveryPreload) TableName() string {
	return "product_delivery_setting"
}

func (p *ProductModelDeliveryModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductModelDeliveryModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductModelDeliveryModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductModelDeliveryModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model ProductModelDeliveryModel
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by", ActiveUser)
	return
}

//strcture input
type ProductModelDeliveryInput struct {
	ProductModelDataID string                     `json:"product_model_data_id" binding:"required"`
	DeliveryData       []ProductModelDeliveryForm `json:"delivery_data" binding:"required,dive"`
}
type ProductModelDeliveryForm struct {
	ID               string `json:"id"`
	CourierTypeID    string `json:"courier_type_id" binding:"required"`
	CourierSubtypeID string `json:"courier_subtype_id" binding:"required"`
	PackagingFee     int    `json:"packaging_fee"`
	Weight           int    `json:"weight"`
	IsSpRequired     int    `json:"is_sp_required"`
}
type ProductModelDeliveryInfo struct {
	ID string `json:"id" binding:"required"`
}
