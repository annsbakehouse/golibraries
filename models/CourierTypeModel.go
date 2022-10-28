package models

import (
	"time"

	"gorm.io/gorm"
)

type CourierTypeModel struct {
	ID               string     `gorm:"column:id;primary_key" json:"id"`
	Name             NullString `gorm:"column:name;default:null;size:100" json:"name"`
	ShortName        NullString `gorm:"column:short_name;default:null;size:40" json:"short_name"`
	CummulativeRange int        `gorm:"column:cummulative_range;default:1" json:"cummulative_range"`
	TimeRange        int        `gorm:"column:time_range;default:1" json:"time_range"`
	Active           uint64     `gorm:"column:active;default:1" json:"active"`
	INC              int        `gorm:"column:inc;default:null" json:"inc"`
	CreatedBy        NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt        time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy        NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt        time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type CourierTypeModelPreload struct {
	ID                 string                        `gorm:"column:id;primary_key" json:"id"`
	Name               NullString                    `gorm:"column:name;default:null;size:100" json:"name"`
	ShortName          NullString                    `gorm:"column:short_name;default:null;size:40" json:"short_name"`
	CummulativeRange   int                           `gorm:"column:cummulative_range;default:1" json:"cummulative_range"`
	TimeRange          int                           `gorm:"column:time_range;default:1" json:"time_range"`
	Active             uint64                        `gorm:"column:active;default:1" json:"active"`
	INC                int                           `gorm:"column:inc;default:1" json:"inc"`
	CourierPlatform    []CourierPlatformModelPreload `gorm:"foreignKey:CourierTypeID;references:ID" json:"courier_platform"`
	CourierSubTypeInfo []CourierSubTypeModelPreload  `gorm:"foreignKey:CourierTypeID;references:ID" json:"courier_subtype_info"`
}

// TableName sets the insert table name for this struct type
func (c *CourierTypeModel) TableName() string {
	return "courier_type"
}
func (c *CourierTypeModelPreload) TableName() string {
	return "courier_type"
}

func (c *CourierTypeModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *CourierTypeModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

type CourierTypeModelInput struct {
	Name             string                      `json:"name" binding:"required"`
	ShortName        string                      `json:"short_name" binding:"required"`
	CummulativeRange int                         `json:"cummulative_range"`
	TimeRange        int                         `json:"time_range"`
	Active           uint64                      `json:"active"`
	PlatformID       []string                    `json:"platform"`
	CourierSubType   []CourierSubTypeModelUpdate `json:"courier_sub_type"`
}

type CourierTypeModelUpdate struct {
	ID               string                      `json:"id" binding:"required"`
	Name             string                      `json:"name" binding:"required"`
	ShortName        string                      `json:"short_name" binding:"required"`
	CummulativeRange int                         `json:"cummulative_range"`
	TimeRange        int                         `json:"time_range"`
	Active           uint64                      `json:"active"`
	PlatformID       []string                    `json:"platform"`
	CourierSubType   []CourierSubTypeModelUpdate `json:"courier_sub_type"`
}

type CourierTypeModelInfo struct {
	ID string `json:"id" binding:"required"`
}

type CourierTypeSelectBySubdistrict struct {
	SubdistrictID string `json:"subdistrict_id" binding:"required"`
}
