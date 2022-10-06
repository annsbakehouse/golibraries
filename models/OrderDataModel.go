package models

import (
	"time"

	"github.com/annsbakehouse/golibraries/library"
	"gorm.io/gorm"
)

type OrderDataModel struct {
	ID                 string     `gorm:"column:id;primary_key" json:"id"`                           //
	OrderInvoiceID     string     `gorm:"column:order_invoice_id" json:"order_invoice_id"`           //
	OrderID            string     `gorm:"column:orderid" json:"orderid"`                             //
	StatusOrder        int        `gorm:"column:status_order" json:"status_order"`                   //
	ShortCode          string     `gorm:"column:short_code" json:"short_code"`                       //
	Notes              NullString `gorm:"column:notes" json:"notes"`                                 //
	OnHold             int        `gorm:"column:on_hold" json:"on_hold"`                             //
	TotalPriceDiscount float64    `gorm:"column:total_price_discount" json:"total_price_discount"`   //
	TotalPrice         float64    `gorm:"column:total_price" json:"total_price"`                     //
	TotalRetailPriceDm float64    `gorm:"column:total_retail_price_dm" json:"total_retail_price_dm"` //
	DeliveryFee        float64    `gorm:"column:delivery_fee" json:"delivery_fee"`                   //
	DeliveryDiscount   float64    `gorm:"column:delivery_discount" json:"delivery_discount"`         //
	PromoValue         float64    `gorm:"column:promo_value" json:"promo_value"`                     //
	PromoDelivery      float64    `gorm:"column:promo_delivery" json:"promo_delivery"`               //
	GrandTotal         float64    `gorm:"column:grand_total" json:"grand_total"`                     //
	PromoCode          NullString `gorm:"column:promo_code" json:"promo_code"`                       //
	PromoName          NullString `gorm:"column:promo_name" json:"promo_name"`                       //
	PromoInfo          NullString `gorm:"column:promo_info" json:"promo_info"`                       //
	CreatedBy          string     `gorm:"column:created_by;default:NULL" json:"created_by"`          //
	CreatedAt          *time.Time `gorm:"column:created_at" json:"created_at"`                       //
	UpdatedBy          string     `gorm:"column:updated_by;default:NULL" json:"updated_by"`          //
	UpdatedAt          *time.Time `gorm:"column:updated_at" json:"updated_at"`                       //
}

type OrderDataModelPreload struct {
	ID                 string     `gorm:"column:id;primary_key" json:"id"`                           //
	OrderInvoiceID     string     `gorm:"column:order_invoice_id" json:"order_invoice_id"`           //
	OrderID            string     `gorm:"column:orderid" json:"orderid"`                             //
	StatusOrder        int        `gorm:"column:status_order" json:"status_order"`                   //
	ShortCode          string     `gorm:"column:short_code" json:"short_code"`                       //
	Notes              NullString `gorm:"column:notes" json:"notes"`                                 //
	OnHold             int        `gorm:"column:on_hold" json:"on_hold"`                             //
	TotalPriceDiscount float64    `gorm:"column:total_price_discount" json:"total_price_discount"`   //
	TotalPrice         float64    `gorm:"column:total_price" json:"total_price"`                     //
	TotalRetailPriceDm float64    `gorm:"column:total_retail_price_dm" json:"total_retail_price_dm"` //
	DeliveryFee        float64    `gorm:"column:delivery_fee" json:"delivery_fee"`                   //
	DeliveryDiscount   float64    `gorm:"column:delivery_discount" json:"delivery_discount"`         //
	PromoValue         float64    `gorm:"column:promo_value" json:"promo_value"`                     //
	PromoDelivery      float64    `gorm:"column:promo_delivery" json:"promo_delivery"`               //
	GrandTotal         float64    `gorm:"column:grand_total" json:"grand_total"`                     //
	PromoCode          NullString `gorm:"column:promo_code" json:"promo_code"`                       //
	PromoName          NullString `gorm:"column:promo_name" json:"promo_name"`                       //
	PromoInfo          NullString `gorm:"column:promo_info" json:"promo_info"`                       //
}

// TableName sets the insert table name for this struct type
func (c *OrderDataModel) TableName() string {
	return "order_data"
}
func (c *OrderDataModelPreload) TableName() string {
	return "order_data"
}

func (c *OrderDataModel) BeforeCreate(tx *gorm.DB) (err error) {
	c.ShortCode = library.RandStringFromDb(8, tx, "order_data", "short_code")
	return
}

func (c *OrderDataModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderDataModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}
