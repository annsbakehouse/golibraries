package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderPaymentModel struct {
	ID             string     `gorm:"column:id;primary_key" json:"id"`                 //
	OrderInvoiceID NullString `gorm:"column:order_invoice_id" json:"order_invoice_id"` //
	PaymentTypeID  NullString `gorm:"column:payment_type_id" json:"payment_type_id"`   //
	PaymentType    NullString `gorm:"column:payment_type" json:"payment_type"`         //
	PaymentDate    NullString `gorm:"column:payment_date" json:"payment_date"`         //
	Value          float64    `gorm:"column:value" json:"value"`                       //
	Status         int        `gorm:"column:status" json:"status"`                     //
	Note           NullString `gorm:"column:note" json:"note"`                         //
	PaidOn         NullString `gorm:"column:paid_on" json:"paid_on"`                   //
	MidtransID     NullString `gorm:"column:midtrans_id" json:"midtrans_id"`           //
	MidtransInfo   NullString `gorm:"column:midtrans_info" json:"midtrans_info"`       //
	MidtransStatus int        `gorm:"column:midtrans_status" json:"midtrans_status"`   //
	MaxTimePay     *time.Time `gorm:"column:max_time_pay" json:"max_time_pay"`         //
	CreatedBy      string     `gorm:"column:created_by" json:"created_by"`             //
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`             //
	UpdatedBy      string     `gorm:"column:updated_by" json:"updated_by"`             //
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`             //
}

type OrderPaymentModelPreload struct {
	ID             string     `gorm:"column:id;primary_key" json:"id"`                 //
	OrderInvoiceID NullString `gorm:"column:order_invoice_id" json:"order_invoice_id"` //
	PaymentTypeID  NullString `gorm:"column:payment_type_id" json:"payment_type_id"`   //
	PaymentType    NullString `gorm:"column:payment_type" json:"payment_type"`         //
	PaymentDate    NullString `gorm:"column:payment_date" json:"payment_date"`         //
	Value          float64    `gorm:"column:value" json:"value"`                       //
	Status         int        `gorm:"column:status" json:"status"`                     //
	Note           NullString `gorm:"column:note" json:"note"`                         //
	PaidOn         NullString `gorm:"column:paid_on" json:"paid_on"`                   //
	MidtransID     NullString `gorm:"column:midtrans_id" json:"midtrans_id"`           //
	MidtransInfo   NullString `gorm:"column:midtrans_info" json:"midtrans_info"`       //
	MidtransStatus int        `gorm:"column:midtrans_status" json:"midtrans_status"`   //
	MaxTimePay     *time.Time `gorm:"column:max_time_pay" json:"max_time_pay"`         //
}

// TableName sets the insert table name for this struct type
func (c *OrderPaymentModel) TableName() string {
	return "order_payment"
}
func (c *OrderPaymentModelPreload) TableName() string {
	return "order_payment"
}

func (c *OrderPaymentModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderPaymentModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderPaymentModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}
