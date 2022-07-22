package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderPackageLogModel struct {
	ID                 string     `gorm:"column:id;primary_key" json:"id"`                         //
	StatusBefore       int        `gorm:"column:status_before" json:"status_before"`               //
	Status             int        `gorm:"column:status" json:"status"`                             //
	ReceivedBy         string     `gorm:"column:received_by" json:"received_by"`                   //
	ReceivedEmployeeBy string     `gorm:"column:received_employee_by" json:"received_employee_by"` //
	PackageLogInfo     string     `gorm:"column:package_log_info" json:"package_log_info"`         //
	OrderPackageID     string     `gorm:"column:order_package_id" json:"order_package_id"`         //
	CreatedBy          string     `gorm:"column:created_by" json:"created_by"`                     //
	CreatedAt          *time.Time `gorm:"column:created_at" json:"created_at"`                     //
}

type OrderPackageLogModelPreload struct {
	ID                 string `gorm:"column:id;primary_key" json:"id"`                         //
	StatusBefore       int    `gorm:"column:status_before" json:"status_before"`               //
	Status             int    `gorm:"column:status" json:"status"`                             //
	ReceivedBy         string `gorm:"column:received_by" json:"received_by"`                   //
	ReceivedEmployeeBy string `gorm:"column:received_employee_by" json:"received_employee_by"` //
	PackageLogInfo     string `gorm:"column:package_log_info" json:"package_log_info"`         //
	OrderPackageID     string `gorm:"column:order_package_id" json:"order_package_id"`         //
}

// TableName sets the insert table name for this struct type
func (c *OrderPackageLogModel) TableName() string {
	return "order_package_log"
}
func (c *OrderPackageLogModelPreload) TableName() string {
	return "order_package_log"
}

func (c *OrderPackageLogModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderPackageLogModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderPackageLogModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}
