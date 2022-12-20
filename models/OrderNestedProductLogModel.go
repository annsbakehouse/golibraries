package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderNestedProductLogModel struct {
	ID                     string     `gorm:"column:id;primary_key" json:"id"`                               //
	StatusBefore           int        `gorm:"column:status_before" json:"status_before"`                     //
	Status                 int        `gorm:"column:status" json:"status"`                                   //
	ReceivedBy             NullString `gorm:"column:received_by" json:"received_by"`                         //
	ReceivedEmployeeBy     NullString `gorm:"column:from_by" json:"from_by"`                                 //
	NestedProductOderInfo  NullString `gorm:"column:order_nested_info" json:"order_nested_info"`             //
	OrderNestedProductID   string     `gorm:"column:order_nested_product_id" json:"order_nested_product_id"` //
	ReceivedEmployeeByInfo NullString `gorm:"column:from_by_info" json:"from_by_info"`                       //
	Images                 NullString `gorm:"column:images;default:null;" json:"images"`
	ReceivedByInfo         NullString `gorm:"column:received_by_info" json:"received_by_info"` //
	AssignedCourier        NullString `gorm:"column:assigned_courier;default:null;" json:"assigned_courier"`
	CourierRemark          NullString `gorm:"column:courier_remark;default:null;" json:"courier_remark"`
	CreatedBy              NullString `gorm:"column:created_by" json:"created_by"` //
	CreatedAt              *time.Time `gorm:"column:created_at" json:"created_at"` //
}

type OrderNestedProductLogModelPreload struct {
	ID                     string                         `gorm:"column:id;primary_key" json:"id"`                         //
	StatusBefore           int                            `gorm:"column:status_before" json:"status_before"`               //
	Status                 int                            `gorm:"column:status" json:"status"`                             //
	ReceivedBy             NullString                     `gorm:"column:received_by" json:"received_by"`                   //
	ReceivedEmployeeBy     NullString                     `gorm:"column:received_employee_by" json:"received_employee_by"` //
	NestedProductOderInfo  NullString                     `gorm:"column:order_nested_info" json:"order_nested_info"`       //
	ReceivedEmployeeByInfo NullString                     `gorm:"column:from_by_info" json:"from_by_info"`                 //
	ReceivedByInfo         NullString                     `gorm:"column:received_by_info" json:"received_by_info"`         //
	Images                 NullString                     `gorm:"column:images;default:null;" json:"images"`
	OrderNestedProductID   string                         `gorm:"column:order_nested_product_id" json:"order_nested_product_id"` //
	OrderProductNestedInfo OrderNestedProductModelPreload `gorm:"->;foreignKey:ID;references:OrderProductNestedID" json:"order_product_nested_info"`
	AssignedCourier        NullString                     `gorm:"column:assigned_courier; default:null" json:"assigned_courier"`
	CourierRemark          NullString                     `gorm:"column:courier_remark;default:null;" json:"courier_remark"`
}

// TableName sets the insert table name for this struct type
func (c *OrderNestedProductLogModel) TableName() string {
	return "order_nested_product_log"
}
func (c *OrderNestedProductLogModelPreload) TableName() string {
	return "order_nested_product_log"
}

func (c *OrderNestedProductLogModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}
func (c *OrderNestedProductLogModel) AfterCreate(tx *gorm.DB) (err error) {
	// var m []OrderPackageLogModel
	// tx.Model(&m).Where("id=?", c.ID).Find(&m)
	// ArangoDbInsert("order_package_log", m)
	return
}

func (c *OrderNestedProductLogModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (c *OrderNestedProductLogModel) AfterUpdate(tx *gorm.DB) (err error) {
	// var m []OrderPackageLogModel
	// tx.Raw("SELECT * FROM order_package_log ? ", tx.Statement.Clauses["WHERE"]).Find(&m)
	// for _, v := range m {
	// 	ArangoDbUpdate("order_package_log", fmt.Sprintf("u.id=='%v'", v.ID), v)
	// }
	return
}
func (c *OrderNestedProductLogModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}
