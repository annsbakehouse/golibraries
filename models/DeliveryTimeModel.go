package models

import (
	"time"

	"gorm.io/gorm"
)

type DeliveryTimeModel struct {
	ID         string     `gorm:"column:id;primaryKey" json:"id"`      //
	Time_start string     `gorm:"column:time_start" json:"time_start"` //
	Time_end   string     `gorm:"column:time_end" json:"time_end"`     //
	Ordering   int64      `gorm:"column:ordering" json:"ordering"`     //
	CreatedBy  NullString `gorm:"column:created_by" json:"created_by"` //
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"` //
	UpdatedBy  NullString `gorm:"column:updated_by" json:"updated_by"` //
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updated_at"` //
}

type DeliveryTimeModelPreload struct {
	ID         string `gorm:"column:id;primaryKey" json:"id"`      //
	Time_start string `gorm:"column:time_start" json:"time_start"` //
	Time_end   string `gorm:"column:time_end" json:"time_end"`     //
	Ordering   int64  `gorm:"column:ordering" json:"ordering"`     //
}

func (p *DeliveryTimeModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

// TableName sets the insert table name for this struct type
func (u *DeliveryTimeModel) TableName() string {
	return "delivery_time"
}
func (u *DeliveryTimeModelPreload) TableName() string {
	return "delivery_time"
}

type DeliveryTimeSaveForm struct {
	Ordering   int64  `json:"ordering" binding:"required"`
	Time_start string `json:"time_start" binding:"required"`
	Time_end   string `json:"time_end" binding:"required"`
}
type DeliveryTimeUpdateForm struct {
	Ordering   int64  `json:"ordering" binding:"required"`
	Time_start string `json:"time_start" binding:"required"`
	Time_end   string `json:"time_end" binding:"required"`
	ID         string `json:"id" binding:"required"`
}
type DeliveryTimeUpdateInfo struct {
	ID string `json:"id" binding:"required"`
}
