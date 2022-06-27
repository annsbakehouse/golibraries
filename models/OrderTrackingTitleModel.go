package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderTrackingTitleModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	Name      string     `gorm:"column:name;default:null" json:"name"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
}
type OrderTrackingTitleModelPreload struct {
	ID   string `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name;default:null" json:"name"`
}

// TableName sets the insert table name for this struct type
func (c *OrderTrackingTitleModel) TableName() string {
	return "order_tracking_title"
}
func (c *OrderTrackingTitleModelPreload) TableName() string {
	return "order_tracking_title"
}

func (c *OrderTrackingTitleModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderTrackingTitleModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

type OrderTrackingTitleInput struct {
	Name string `json:"name" binding:"required"`
}

type OrderTrackingTitleUpdate struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
type OrderTrackingTitleInfo struct {
	ID string `json:"id" binding:"required"`
}
