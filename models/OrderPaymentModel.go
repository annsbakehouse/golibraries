package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderPaymentModel struct {
	ID             string     `gorm:"column:id;primary_key" json:"id"`                   //
	OrderInvoiceID NullString `gorm:"column:order_invoice_id" json:"order_invoice_id"`   //
	PaymentTypeID  NullString `gorm:"column:payment_type_id" json:"payment_type_id"`     //
	PaymentType    NullString `gorm:"column:payment_type" json:"payment_type"`           //
	PaymentDate    NullString `gorm:"column:payment_date" json:"payment_date"`           //
	Value          float64    `gorm:"column:value" json:"value"`                         //
	Status         int        `gorm:"column:status" json:"status"`                       //
	Note           NullString `gorm:"column:note" json:"note"`                           //
	PaidOn         NullString `gorm:"column:paid_on" json:"paid_on"`                     //
	PaymentID      NullString `gorm:"column:payment_id" json:"payment_id"`               //
	PaymentInfo    NullString `gorm:"column:payment_info" json:"payment_info"`           //
	PaymentStatus  int        `gorm:"column:payment_status" json:"payment_status"`       //
	MaxTimePay     NullString `gorm:"column:max_time_pay" json:"max_time_pay"`           //
	ApproveStatus  int        `gorm:"column:approve_status" json:"approve_status"`       //
	CreatedBy      string     `gorm:"column:created_by;default:null;" json:"created_by"` //
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`               //
	UpdatedBy      string     `gorm:"column:updated_by;default:null;" json:"updated_by"` //
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`               //
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
	PaymentID      NullString `gorm:"column:payment_id" json:"payment_id"`             //
	PaymentInfo    NullString `gorm:"column:payment_info" json:"payment_info"`         //
	PaymentStatus  int        `gorm:"column:payment_status" json:"payment_status"`     //
	MaxTimePay     *time.Time `gorm:"column:max_time_pay" json:"max_time_pay"`         //
	ApproveStatus  int        `gorm:"column:approve_status" json:"approve_status"`     //
}

type OrderPaymentForm struct {
	PaymentTypeId string  `json:"payment_type_id" binding:"required"`
	PaymentDate   string  `json:"payment_date" binding:"required"`
	Value         float64 `json:"value" binding:"omitempty"`
	Status        int     `json:"status;default:0" binding:"omitempty"`
	Note          string  `json:"note" binding:"required"`
	PaymentStatus int     `json:"payment_status;default:0" binding:"omitempty"`
}

type OrderPaymentEditForm struct {
	Id            string  `json:"id" binding:"required"`
	PaymentTypeId string  `json:"payment_type_id" binding:"required"`
	PaymentDate   string  `json:"payment_date" binding:"required"`
	Value         float64 `json:"value" binding:"omitempty"`
	Status        int     `json:"status;default:0" binding:"omitempty"`
	Note          string  `json:"note" binding:"required"`
	PaymentStatus int     `json:"payment_status;default:0" binding:"omitempty"`
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
func (c *OrderPaymentModel) AfterCreate(tx *gorm.DB) (err error) {
	// var m []OrderPaymentModel
	// tx.Model(&m).Where("id=?", c.ID).Find(&m)
	// ArangoDbInsert("order_payment", m)
	return
}
func (c *OrderPaymentModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (c *OrderPaymentModel) AfterUpdate(tx *gorm.DB) (err error) {
	// var m []OrderPaymentModel
	// tx.Raw("SELECT * FROM order_payment ? ", tx.Statement.Clauses["WHERE"]).Find(&m)
	// for _, v := range m {
	// 	ArangoDbUpdate("order_payment", fmt.Sprintf("u.id=='%v'", v.ID), v)
	// }
	return
}
func (c *OrderPaymentModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}
