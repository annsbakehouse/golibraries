package models

import (
	"gorm.io/gorm"
)

type OrderDeliveryLogModel struct {
	ID                    string     `gorm:"column:id;primary_key" json:"id"`                                 //
	OrderDataID           NullString `gorm:"column:order_data_id" json:"order_data_id"`                       //
	OrderProductPackageID NullString `gorm:"column:order_product_package_id" json:"order_product_package_id"` //
	CourierTypeID         NullString `gorm:"column:courier_type_id" json:"courier_type_id"`                   //
	CourierSubtypeID      NullString `gorm:"column:courier_subtype_id" json:"courier_subtype_id"`             //
	EmployeeID            NullString `gorm:"column:employee_id" json:"employee_id"`                           //
	CourierName           NullString `gorm:"column:courier_name" json:"courier_name"`                         //
	StatusBefore          int        `gorm:"column:status_before" json:"status_before"`                       //
	Status                int        `gorm:"column:status" json:"status"`                                     //
	OrderDeliveryID       NullString `gorm:"column:order_delivery_id" json:"order_delivery_id"`               //
	StatusDelivery        int        `gorm:"column:status_delivery" json:"status_delivery"`                   //
	DeliveryLogInfo       NullString `gorm:"column:delivery_log_info" json:"delivery_log_info"`               //
}

type OrderDeliveryLogModelPreload struct {
	ID                    string     `gorm:"column:id;primary_key" json:"id"`                                 //
	OrderDataID           NullString `gorm:"column:order_data_id" json:"order_data_id"`                       //
	OrderProductPackageID NullString `gorm:"column:order_product_package_id" json:"order_product_package_id"` //
	CourierTypeID         NullString `gorm:"column:courier_type_id" json:"courier_type_id"`                   //
	CourierSubtypeID      NullString `gorm:"column:courier_subtype_id" json:"courier_subtype_id"`             //
	EmployeeID            NullString `gorm:"column:employee_id" json:"employee_id"`                           //
	CourierName           NullString `gorm:"column:courier_name" json:"courier_name"`                         //
	StatusBefore          int        `gorm:"column:status_before" json:"status_before"`                       //
	Status                int        `gorm:"column:status" json:"status"`                                     //
	OrderDeliveryID       NullString `gorm:"column:order_delivery_id" json:"order_delivery_id"`               //
	StatusDelivery        int        `gorm:"column:status_delivery" json:"status_delivery"`                   //
	DeliveryLogInfo       NullString `gorm:"column:delivery_log_info" json:"delivery_log_info"`               //
}

// TableName sets the insert table name for this struct type
func (c *OrderDeliveryLogModel) TableName() string {
	return "order_delivery_log"
}
func (c *OrderDeliveryLogModelPreload) TableName() string {
	return "order_delivery_log"
}

func (c *OrderDeliveryLogModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderDeliveryLogModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderDeliveryLogModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}
