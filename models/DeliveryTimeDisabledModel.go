package models

import (
	"time"

	"gorm.io/gorm"
)

type DeliveryTimeDisabledModel struct {
	ID             string     `gorm:"column:id;primaryKey" json:"id"`                  //
	DeliveryTimeID NullString `gorm:"column:delivery_time_id" json:"delivery_time_id"` //
	WarehouseID    string     `gorm:"column:warehouse_id" json:"warehouse_id"`         //
	CourierTypeID  NullString `gorm:"column:courier_type_id" json:"courier_type_id"`   //
	DateStart      NullDate   `gorm:"column:date_start" json:"date_start"`             //
	DateEnd        NullDate   `gorm:"column:date_end" json:"date_end"`                 //
	DisabledType   int        `gorm:"column:disabled_type" json:"disabled_type"`       //
	CreatedBy      NullString `gorm:"column:created_by" json:"created_by"`             //
	CreatedAt      time.Time  `gorm:"column:created_at" json:"created_at"`             //
	UpdatedBy      NullString `gorm:"column:updated_by" json:"updated_by"`             //
	UpdatedAt      time.Time  `gorm:"column:updated_at" json:"updated_at"`             //
}

type DeliveryTimeDisabledModelPreload struct {
	ID               string                   `gorm:"column:id;primaryKey" json:"id"`                  //
	DeliveryTimeID   NullString               `gorm:"column:delivery_time_id" json:"delivery_time_id"` //
	WarehouseID      string                   `gorm:"column:warehouse_id" json:"warehouse_id"`         //
	CourierTypeID    NullString               `gorm:"column:courier_type_id" json:"courier_type_id"`   //
	DateStart        NullDate                 `gorm:"column:date_start" json:"date_start"`             //
	DateEnd          NullDate                 `gorm:"column:date_end" json:"date_end"`                 //
	DisabledType     int                      `gorm:"column:disabled_type" json:"disabled_type"`       //
	DeliveryTimeInfo DeliveryTimeModelPreload `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DeliveryTimeID;references:ID" json:"delivery_time_info"`
	WarehouseInfo    WarehousePreloadModel    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:WarehouseID;references:ID" json:"warehouse_info"`
	CourierTypeInfo  DeliveryTimeModelPreload `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CourierTypeID;references:ID" json:"courier_type_info"`
}

func (p *DeliveryTimeDisabledModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

// TableName sets the insert table name for this struct type
func (u *DeliveryTimeDisabledModel) TableName() string {
	return "delivery_time_disabled"
}
func (u *DeliveryTimeDisabledModelPreload) TableName() string {
	return "delivery_time_disabled"
}

type DeliveryTimeDisabledType2SaveForm struct {
	WarehouseID    string `json:"warehouse_id" binding:"required"`     //
	DateStart      string `json:"date_start" binding:"required"`       //
	DateEnd        string `json:"date_end" binding:"required"`         //
	DeliveryTimeID string `json:"delivery_time_id" binding:"required"` //
	CourierTypeID  string `json:"courier_type_id" binding:"required"`  //
}
type DeliveryTimeDisabledType2UpdateForm struct {
	ID             string `json:"id" binding:"required"`               //
	WarehouseID    string `json:"warehouse_id" binding:"required"`     //
	DateStart      string `json:"date_start" binding:"required"`       //
	DateEnd        string `json:"date_end" binding:"required"`         //
	DeliveryTimeID string `json:"delivery_time_id" binding:"required"` //
	CourierTypeID  string `json:"courier_type_id" binding:"required"`  //
}
type DeliveryTimeDisabledType1SaveForm struct {
	WarehouseID string `json:"warehouse_id" binding:"required"` //
	DateStart   string `json:"date_start" binding:"required"`   //
	DateEnd     string `json:"date_end" binding:"required"`     //
}
type DeliveryTimeDisabledType1UpdateForm struct {
	ID          string `json:"id" binding:"required"`           //
	WarehouseID string `json:"warehouse_id" binding:"required"` //
	DateStart   string `json:"date_start" binding:"required"`   //
	DateEnd     string `json:"date_end" binding:"required"`     //
}
type DeliveryTimeDisabledUpdateInfo struct {
	ID string `json:"id" binding:"required"`
}
