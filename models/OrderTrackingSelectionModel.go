package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderTrackingSelectionModel struct {
	ID                   string     `gorm:"column:id;primary_key" json:"id"`
	Name                 string     `gorm:"column:name;default:null" json:"name"`
	OrderTrackingTitleID string     `gorm:"column:order_tracking_title_id" json:"order_tracking_title_id"`
	CreatedBy            NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt            time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy            NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt            time.Time  `gorm:"column:updated_at" json:"updated_at"`
}
type OrderTrackingSelectionModelPreload struct {
	ID                     string                           `gorm:"column:id;primary_key" json:"id"`
	Name                   string                           `gorm:"column:name;default:null" json:"name"`
	OrderTrackingTitleID   string                           `gorm:"column:order_tracking_title_id" json:"order_tracking_title_id"`
	OrderTrackingTitleInfo []OrderTrackingTitleModelPreload `gorm:"foreignKey:ID;references:OrderTrackingTitleID" json:"order_tracking_title_info"`
}

// TableName sets the insert table name for this struct type
func (c *OrderTrackingSelectionModel) TableName() string {
	return "order_tracking_selection"
}
func (c *OrderTrackingSelectionModelPreload) TableName() string {
	return "order_tracking_selection"
}

func (c *OrderTrackingSelectionModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderTrackingSelectionModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

type OrderTrackingSelectionInput struct {
	Name                 string `json:"name" binding:"required"`
	OrderTrackingTitleID string `json:"order_tracking_title_id" binding:"required"`
}

type OrderTrackingSelectionUpdate struct {
	ID                   string `json:"id" binding:"required"`
	Name                 string `json:"name" binding:"required"`
	OrderTrackingTitleID string `json:"order_tracking_title_id" binding:"required"`
}
type OrderTrackingSelectionInfo struct {
	ID string `json:"id" binding:"required"`
}
