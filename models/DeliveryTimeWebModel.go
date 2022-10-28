package models

import (
	"time"

	"gorm.io/gorm"
)

type DeliveryTimeWebModel struct {
	ID             string     `gorm:"column:id;primaryKey" json:"id"`                  //
	DeliveryTimeID string     `gorm:"column:delivery_time_id" json:"delivery_time_id"` //
	WarehouseID    string     `gorm:"column:warehouse_id" json:"warehouse_id"`         //
	CourierTypeID  string     `gorm:"column:courier_type_id" json:"courier_type_id"`   //
	Ordering       int        `gorm:"column:ordering" json:"ordering"`                 //
	CreatedBy      NullString `gorm:"column:created_by" json:"created_by"`             //
	CreatedAt      time.Time  `gorm:"column:created_at" json:"created_at"`             //
	UpdatedBy      NullString `gorm:"column:updated_by" json:"updated_by"`             //
	UpdatedAt      time.Time  `gorm:"column:updated_at" json:"updated_at"`             //
}

type DeliveryTimeWebModelPreload struct {
	ID               string                   `gorm:"column:id;primaryKey" json:"id"`                  //
	DeliveryTimeID   string                   `gorm:"column:delivery_time_id" json:"delivery_time_id"` //
	WarehouseID      string                   `gorm:"column:warehouse_id" json:"warehouse_id"`         //
	CourierTypeID    string                   `gorm:"column:courier_type_id" json:"courier_type_id"`   //
	Ordering         int                      `gorm:"column:ordering" json:"ordering"`                 //
	DeliveryTimeInfo DeliveryTimeModelPreload `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DeliveryTimeID;references:ID" json:"delivery_time_info"`
	WarehouseInfo    WarehousePreloadModel    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:WarehouseID;references:ID" json:"warehouse_info"`
	CourierTypeInfo  DeliveryTimeModelPreload `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CourierTypeID;references:ID" json:"courier_type_info"`
}

func (p *DeliveryTimeWebModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

// TableName sets the insert table name for this struct type
func (u *DeliveryTimeWebModel) TableName() string {
	return "delivery_time_customer"
}
func (u *DeliveryTimeWebModelPreload) TableName() string {
	return "delivery_time_customer"
}

type DeliveryTimeWebSaveForm struct {
	Ordering       int    `json:"ordering" binding:"required"`
	DeliveryTimeID string `json:"delivery_time_id" binding:"required"` //
	WarehouseID    string `json:"warehouse_id" binding:"required"`     //
	CourierTypeID  string `json:"courier_type_id" binding:"required"`  //
}
type DeliveryTimeWebUpdateForm struct {
	Ordering       int    `json:"ordering" binding:"required"`
	ID             string `json:"id" binding:"required"`               //
	DeliveryTimeID string `json:"delivery_time_id" binding:"required"` //
	WarehouseID    string `json:"warehouse_id" binding:"required"`     //
	CourierTypeID  string `json:"courier_type_id" binding:"required"`  //
}
type DeliveryTimeWebUpdateInfo struct {
	ID string `json:"id" binding:"required"`
}
