package models

import (
	"time"

	"github.com/annsbakehouse/golibraries/library"
	"gorm.io/gorm"
)

// InVid 4 alpha 3 numeric
type OrderInvoiceModel struct {
	ID                     string     `gorm:"column:id;primary_key" json:"id"`                           //
	InvoiceDate            NullString `gorm:"column:invoice_date" json:"invoice_date"`                   //
	InVid                  NullString `gorm:"column:invid" json:"invid"`                                 //
	CustomerRetailID       NullString `gorm:"column:customer_retail_id" json:"customer_retail_id"`       //
	CostomerCorporateID    NullString `gorm:"column:customer_corporate_id" json:"customer_corporate_id"` //
	CustomerInfo           NullString `gorm:"column:customer_info" json:"customer_info"`
	ReferralID             NullString `gorm:"column:referral_id" json:"referral_id"`                             //
	CustomerCorporatePicID NullString `gorm:"column:customer_corporate_pic_id" json:"customer_corporate_pic_id"` //
	ShortCode              NullString `gorm:"column:short_code" json:"short_code"`                               //
	StatusINV              int        `gorm:"column:status_inv" json:"status_inv"`                               //
	PlatformID             NullString `gorm:"column:platform_id" json:"platform_id"`                             //
	CustomersLevelID       NullString `gorm:"column:customers_level_id" json:"customers_level_id"`               //
	ContactMethodID        NullString `gorm:"column:contact_method_id" json:"contact_method_id"`                 //
	SumGrandTotal          float64    `gorm:"column:sum_grand_total" json:"sum_grand_total"`                     //
	Note                   NullString `gorm:"column:note" json:"note"`
	DueDate                NullString `gorm:"column:due_date" json:"due_date"`
	FirstDeliveryDate      NullString `gorm:"column:first_delivery_date" json:"first_delivery_date"`
	FirstDeliveryTimeStart NullString `gorm:"column:first_delivery_time_start" json:"first_delivery_time_start"`
	FirstDeliveryTimeEnd   NullString `gorm:"column:first_delivery_time_end" json:"first_delivery_time_end"`
	INC                    int        `gorm:"->;column:inc" json:"inc"`                         //
	CreatedBy              string     `gorm:"column:created_by;default:null" json:"created_by"` //
	CreatedAt              *time.Time `gorm:"column:created_at" json:"created_at"`              //
	UpdatedBy              string     `gorm:"column:updated_by;default:null" json:"updated_by"` //
	UpdatedAt              *time.Time `gorm:"column:updated_at" json:"updated_at"`              //
}

type OrderInvoiceModelPreload struct {
	ID                     string                             `gorm:"column:id;primary_key" json:"id"`                           //
	InvoiceDate            NullString                         `gorm:"column:invoice_date" json:"invoice_date"`                   //
	InVid                  NullString                         `gorm:"column:invid" json:"invid"`                                 //
	CustomerRetailID       NullString                         `gorm:"column:customer_retail_id" json:"customer_retail_id"`       //
	CostomerCorporateID    NullString                         `gorm:"column:customer_corporate_id" json:"customer_corporate_id"` //
	CustomerInfo           NullString                         `gorm:"column:customer_info" json:"customer_info"`
	ReferralID             NullString                         `gorm:"column:referral_id" json:"referral_id"`                             //
	CustomerCorporatePicID NullString                         `gorm:"column:customer_corporate_pic_id" json:"customer_corporate_pic_id"` //
	ShortCode              NullString                         `gorm:"column:short_code" json:"short_code"`                               //
	StatusINV              NullString                         `gorm:"column:status_inv" json:"status_inv"`                               //
	PlatformID             NullString                         `gorm:"column:platform_id" json:"platform_id"`                             //
	CustomersLevelID       NullString                         `gorm:"column:customers_level_id" json:"customers_level_id"`               //
	ContactMethodID        NullString                         `gorm:"column:contact_method_id" json:"contact_method_id"`                 //
	SumGrandTotal          float64                            `gorm:"column:sum_grand_total" json:"sum_grand_total"`                     //
	Note                   NullString                         `gorm:"column:note" json:"note"`
	DueDate                NullString                         `gorm:"column:due_date" json:"due_date"`
	FirstDeliveryDate      NullString                         `gorm:"column:first_delivery_date" json:"first_delivery_date"`
	FirstDeliveryTimeStart NullString                         `gorm:"column:first_delivery_time_start" json:"first_delivery_time_start"`
	FirstDeliveryTimeEnd   NullString                         `gorm:"column:first_delivery_time_end" json:"first_delivery_time_end"`
	INC                    int                                `gorm:"->;column:inc" json:"inc"`
	InvoiceTracking        []OrderInvoiceTrackingModelPreload `gorm:"->;foreignKey:OrderInvoiceID;references:ID" json:"invoice_tracking"`
	OrderDataInfo          []OrderDataModelPreload            `gorm:"->;foreignKey:OrderInvoiceID;references:ID" json:"order_data_info"`
	OrderPaymentInfo       []OrderPaymentModelPreload         `gorm:"->;foreignKey:OrderInvoiceID;references:ID" json:"order_payment_info"`
}

// TableName sets the insert table name for this struct type
func (c *OrderInvoiceModel) TableName() string {
	return "order_invoice"
}
func (c *OrderInvoiceModelPreload) TableName() string {
	return "order_invoice"
}

func (c *OrderInvoiceModel) BeforeCreate(tx *gorm.DB) (err error) {
	c.ShortCode = NullStringInput(library.RandStringFromDb(8, tx, "order_invoice", "short_code"))
	return
}
func (c *OrderInvoiceModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}
func (c *OrderInvoiceModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderInvoiceModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}

type RetailInvoiceOrderAdd struct {
	InvoiceDate      string                     `json:"invoice_date" binding:"required"`
	CustomerRetailID string                     `json:"customer_retail_id" binding:"required"`
	CustomersLevelID string                     `json:"customers_level_id" binding:"required"`
	SumGrandTotal    float64                    `json:"sum_grand_total"`
	InvoiceTracking  []OrderInvoiceTrackingSave `json:"invoice_tracking" binding:"dive"`
}
