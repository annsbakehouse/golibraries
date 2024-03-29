package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderDeliveryModel struct {
	ID                string     `gorm:"column:id;primary_key" json:"id"`                                    //
	OrderInvoiceID    string     `gorm:"column:order_invoice_id" json:"order_invoice_id"`                    //
	OrderDataID       string     `gorm:"column:order_data_id" json:"order_data_id"`                          //
	CourierTypeID     string     `gorm:"column:courier_type_id;default:null;" json:"courier_type_id"`        //
	CustomerAddressID string     `gorm:"column:customer_address_id;default:NULL" json:"customer_address_id"` //
	CourierName       string     `gorm:"column:courier_name;default:null" json:"courier_name"`               //
	DeliveryDate      NullString `gorm:"column:delivery_date" json:"delivery_date"`                          //
	DeliveryStart     NullString `gorm:"column:delivery_start" json:"delivery_start"`                        //
	DeliveryEnd       NullString `gorm:"column:delivery_end" json:"delivery_end"`                            //
	SenderName        NullString `gorm:"column:sender_name" json:"sender_name"`                              //
	SenderPhone       NullString `gorm:"column:sender_phone" json:"sender_phone"`                            //
	StatusDelivery    int        `gorm:"column:status_delivery" json:"status_delivery"`                      //
	OrderDeliveryInfo string     `gorm:"column:order_delivery_info;default:null" json:"order_delivery_info"`
	DeliveryTimeID    string     `gorm:"delivery_time_id" json:"delivery_time_id"`
	WarehouseID       NullString `gorm:"warehouse_id" json:"warehouse_id"`
	AssignedCourier   NullString `gorm:"assigned_courier" json:"assigned_courier"`
	CourierRemark     NullString `gorm:"courier_remark" json:"courier_remark"`
	DeliveryTime      NullString `gorm:"delivery_time" json:"delivery_time"`
	DeliveryNote      NullString `gorm:"delivery_note" json:"delivery_note"`
	DeliveryDateReal  NullString `gorm:"delivery_date_real" json:"delivery_date_real"`
	Images            NullString `gorm:"images" json:"images"`
	CreatedBy         string     `gorm:"column:created_by;default:null;" json:"created_by"` //
	CreatedAt         *time.Time `gorm:"column:created_at" json:"created_at"`               //
	UpdatedBy         string     `gorm:"column:updated_by;default:null;" json:"updated_by"` //
	UpdatedAt         *time.Time `gorm:"column:updated_at" json:"updated_at"`               //
}

type OrderDeliveryModelPreload struct {
	ID                string     `gorm:"column:id;primary_key" json:"id"`                       //
	OrderInvoiceID    string     `gorm:"column:order_invoice_id" json:"order_invoice_id"`       //
	OrderDataID       string     `gorm:"column:order_data_id" json:"order_data_id"`             //
	CourierTypeID     string     `gorm:"column:courier_type_id" json:"courier_type_id"`         //
	CustomerAddressID string     `gorm:"column:customer_address_id" json:"customer_address_id"` //
	CourierName       string     `gorm:"column:courier_name" json:"courier_name"`               //
	DeliveryDate      NullString `gorm:"column:delivery_date" json:"delivery_date"`             //
	DeliveryStart     NullString `gorm:"column:delivery_start" json:"delivery_start"`           //
	DeliveryEnd       NullString `gorm:"column:delivery_end" json:"delivery_end"`               //
	SenderName        NullString `gorm:"column:sender_name" json:"sender_name"`                 //
	SenderPhone       NullString `gorm:"column:sender_phone" json:"sender_phone"`               //
	StatusDelivery    int        `gorm:"column:status_delivery" json:"status_delivery"`         //
	OrderDeliveryInfo string     `gorm:"order_delivery_info" json:"order_delivery_info"`
	DeliveryTimeID    string     `gorm:"delivery_time_id" json:"delivery_time_id"`
	WarehouseID       NullString `gorm:"warehouse_id" json:"warehouse_id"`
	AssignedCourier   NullString `gorm:"assigned_courier" json:"assigned_courier"`
	CourierRemark     NullString `gorm:"courier_remark" json:"courier_remark"`
	DeliveryTime      NullString `gorm:"delivery_time" json:"delivery_time"`
	DeliveryNote      NullString `gorm:"delivery_note" json:"delivery_note"`
	DeliveryDateReal  NullString `gorm:"delivery_date_real" json:"delivery_date_real"`
	Images            NullString `gorm:"images" json:"images"`
}

// TableName sets the insert table name for this struct type
func (c *OrderDeliveryModel) TableName() string {
	return "order_delivery"
}
func (c *OrderDeliveryModelPreload) TableName() string {
	return "order_delivery"
}

func (c *OrderDeliveryModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}
func (c *OrderDeliveryModel) AfterCreate(tx *gorm.DB) (err error) {
	// var m []OrderDeliveryModel
	// tx.Model(&m).Where("id=?", c.ID).Find(&m)
	// ArangoDbInsert("order_delivery", m)
	return
}
func (c *OrderDeliveryModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (c *OrderDeliveryModel) AfterUpdate(tx *gorm.DB) (err error) {
	// var m []OrderDeliveryModel
	// tx.Raw("SELECT * FROM order_delivery ? ", tx.Statement.Clauses["WHERE"]).Find(&m)
	// for _, v := range m {
	// 	ArangoDbUpdate("order_delivery", fmt.Sprintf("u.id=='%v'", v.ID), v)
	// }
	return
}

func (c *OrderDeliveryModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}
