package models

import (
	"time"

	"gorm.io/gorm"
)

type InternalVehicleModel struct {
	ID           string     `gorm:"column:id;primary_key" json:"id"`
	Type         NullString `gorm:"column:type;default:null" json:"type"`
	FlatCapacity int        `gorm:"column:flat_capacity;default:null" json:"flat_capacity"`
	TimeRange    NullString `gorm:"column:time_range;default:null" json:"time_range"`
	Active       uint64     `gorm:"column:active;default:1" json:"active"`
	IVID         string     `gorm:"column:ivid" json:"ivid"`
	CreatedBy    NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy    NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type InternalVehicleModelPreload struct {
	ID           string                                    `gorm:"column:id;primary_key" json:"id"`
	Type         NullString                                `gorm:"column:type;default:null" json:"type"`
	FlatCapacity int                                       `gorm:"column:flat_capacity;default:null" json:"flat_capacity"`
	TimeRange    NullString                                `gorm:"column:time_range;default:null" json:"time_range"`
	Active       uint64                                    `gorm:"column:active;default:1" json:"active"`
	IVID         string                                    `gorm:"column:ivid" json:"ivid"`
	RackCapacity []InternalVehicleRackCapacityModelPreload `gorm:"foreignKey:InternalVehicleID;references:ID" json:"rack_capacity"`
}

// TableName sets the insert table name for this struct type
func (c *InternalVehicleModel) TableName() string {
	return "internal_vehicle"
}
func (c *InternalVehicleModelPreload) TableName() string {
	return "internal_vehicle"
}

func (c *InternalVehicleModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *InternalVehicleModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

type InternalVehicleModelInput struct {
	Type         string                             `json:"type" binding:"required"`
	IVID         string                             `son:"ivid" binding:"required"`
	Active       uint64                             `json:"active"`
	FlatCapacity int                                `json:"flat_capacity"`
	RackCapacity []InternalVehicleRackCapacityInput `json:"rack_capacity"`
}

type InternalVehicleModelUpdate struct {
	ID           string                             `json:"id" binding:"required"`
	IVID         string                             `son:"ivid" binding:"required"`
	Type         string                             `json:"type" binding:"required"`
	Active       uint64                             `json:"active"`
	FlatCapacity int                                `json:"flat_capacity"`
	RackCapacity []InternalVehicleRackCapacityInput `json:"rack_capacity"`
}

type InternalVehicleModelInfo struct {
	ID string `json:"id" binding:"required"`
}
