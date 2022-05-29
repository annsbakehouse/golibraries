package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductModelWarehouseModel struct {
	ID                 string     `gorm:"column:id;primary_key" json:"id"`
	ProductModelDataID string     `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	WarehouseType      int        `gorm:"column:warehouse_type" json:"warehouse_type"`
	WarehouseFrom      string     `gorm:"column:warehouse_from" json:"warehouse_from"`
	WarehouseTo        string     `gorm:"column:warehouse_to" json:"warehouse_to"`
	ProductionTime     float64    `gorm:"column:production_time" json:"production_time"`
	CreatedBy          NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt          time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy          NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy          NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt          soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductModelWarehouseModelPreload struct {
	ID                 string  `gorm:"column:id;primary_key" json:"id"`
	ProductModelDataID string  `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	WarehouseType      int     `gorm:"column:warehouse_type" json:"warehouse_type"`
	WarehouseFrom      string  `gorm:"column:warehouse_from" json:"warehouse_from"`
	WarehouseTo        string  `gorm:"column:warehouse_to" json:"warehouse_to"`
	ProductionTime     float64 `gorm:"column:production_time" json:"production_time"`
	// DeletedAt          soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (p *ProductModelWarehouseModel) TableName() string {
	return "product_warehouse"
}
func (p *ProductModelWarehouseModelPreload) TableName() string {
	return "product_warehouse"
}

func (p *ProductModelWarehouseModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductModelWarehouseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductModelWarehouseModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductModelWarehouseModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model ProductModelWarehouseModel
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by", ActiveUser)
	return
}

//strcture input
type ProductModelWarehouseFromTo struct {
	ID             string  `json:"id"`
	From           string  `json:"from" binding:"required"`
	To             string  `json:"to" binding:"required"`
	ProductionTime float64 `json:"production_time"`
}
type ProductModelWarehouseInput struct {
	ProductModelDataID string                        `json:"product_model_data_id" binding:"required"`
	W4W2               []ProductModelWarehouseFromTo `json:"w4w2" binding:"required,dive"`
	W3W2               []ProductModelWarehouseFromTo `json:"w3w2" binding:"required,dive"`
}
type ProductModelWarehouseInfo struct {
	ID string `json:"id" binding:"required"`
}
