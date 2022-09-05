package models

import (
	"time"

	"gorm.io/gorm"
)

type HiltLogModel struct {
	IP         string    `gorm:"column:ip" json:"ip"`
	DeviceID   string    `gorm:"column:device_id" json:"device_id"`
	CustTime   time.Time `gorm:"column:cust_time" json:"cust_time"`
	EndPoint   string    `gorm:"column:end_point" json:"end_point"`
	DeviceInfo string    `gorm:"column:device_info" json:"device_info"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	// DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type HiltLogModelInput struct {
	IP         string    `json:"ip" binding:"required"`
	DeviceID   string    `json:"device_id" binding:"required"`
	CustTime   time.Time `json:"cust_time" binding:"required"`
	EndPoint   string    `json:"end_point" binding:"required"`
	DeviceInfo string    `json:"device_info"`
	CreatedAt  time.Time `json:"created_at"`
	// DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (p *HiltLogModel) TableName() string {
	return "hilt_log"
}

func (p *HiltLogModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *HiltLogModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *HiltLogModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *HiltLogModel) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println("Before Delete")
	// var inventoryUnitModel InventoryUnitModel
	// tx.Model(&inventoryUnitModel).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return
}
