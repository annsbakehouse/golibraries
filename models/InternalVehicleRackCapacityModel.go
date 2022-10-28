package models

import (
	"time"

	"gorm.io/gorm"
)

type InternalVehicleRackCapacityModel struct {
	ID                string     `gorm:"column:id;primary_key" json:"id"`
	InternalVehicleID string     `gorm:"column:internal_vehicle_id;default:null" json:"internal_vehicle_id"`
	Capacity          int        `gorm:"column:capacity;default:null" json:"capacity"`
	INC               int        `gorm:"column:inc;default:1" json:"inc"`
	CreatedBy         NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt         time.Time  `gorm:"column:created_at" json:"created_at"`
	DeletedBy         NullString `gorm:"column:deleted_by" json:"deleted_by"`
	//DeletedAt  				soft_delete.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

type InternalVehicleRackCapacityModelPreload struct {
	ID                string `gorm:"column:id;primary_key" json:"id"`
	InternalVehicleID string `gorm:"column:internal_vehicle_id;default:null" json:"internal_vehicle_id"`
	Capacity          int    `gorm:"column:capacity;default:null" json:"capacity"`
	INC               int    `gorm:"column:inc;default:1" json:"inc"`
	//DeletedAt         soft_delete.DeletedAt `gorm:"column:deleted_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (c *InternalVehicleRackCapacityModel) TableName() string {
	return "internal_vehicle_rack_capacity"
}
func (c *InternalVehicleRackCapacityModelPreload) TableName() string {
	return "internal_vehicle_rack_capacity"
}

func (c *InternalVehicleRackCapacityModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *InternalVehicleRackCapacityModel) BeforeDelete(tx *gorm.DB) (err error) {
	var model InternalVehicleRackCapacityModel
	tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}

type InternalVehicleRackCapacityInput struct {
	ID       string `json:"id"`
	Capacity int    `json:"capacity"`
}
type InternalVehicleRackCapacityUpdate struct {
	ID       string `json:"id" binding:"required"`
	Capacity int    `json:"capacity" binding:"required"`
}
