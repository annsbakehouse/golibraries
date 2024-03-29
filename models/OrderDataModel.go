package models

import (
	"time"

	"github.com/annsbakehouse/golibraries/library"
	"gorm.io/gorm"
)

type OrderDataModel struct {
	ID                  string     `gorm:"column:id;primary_key" json:"id"`                           //
	OrderInvoiceID      string     `gorm:"column:order_invoice_id" json:"order_invoice_id"`           //
	OrderID             string     `gorm:"column:orderid" json:"orderid"`                             //
	StatusOrder         int        `gorm:"column:status_order" json:"status_order"`                   //
	ShortCode           string     `gorm:"column:short_code" json:"short_code"`                       //
	Notes               NullString `gorm:"column:notes" json:"notes"`                                 //
	OnHold              int        `gorm:"column:on_hold" json:"on_hold"`                             //
	TotalPriceDiscount  float64    `gorm:"column:total_price_discount" json:"total_price_discount"`   //
	TotalPrice          float64    `gorm:"column:total_price" json:"total_price"`                     //
	TotalRetailPriceDm  float64    `gorm:"column:total_retail_price_dm" json:"total_retail_price_dm"` //
	DeliveryFee         float64    `gorm:"column:delivery_fee" json:"delivery_fee"`                   //
	DeliveryDiscount    float64    `gorm:"column:delivery_discount" json:"delivery_discount"`         //
	PromoValue          float64    `gorm:"column:promo_value" json:"promo_value"`                     //
	PromoDelivery       float64    `gorm:"column:promo_delivery" json:"promo_delivery"`               //
	GrandTotal          float64    `gorm:"column:grand_total" json:"grand_total"`                     //
	PromoCode           NullString `gorm:"column:promo_code" json:"promo_code"`                       //
	PromoName           NullString `gorm:"column:promo_name" json:"promo_name"`                       //
	PromoInfo           NullString `gorm:"column:promo_info" json:"promo_info"`                       //
	PrintDeliveryStatus int        `gorm:"column:print_delivery_status;default:0" json:"print_delivery_status"`
	PrintPickupStatus   int        `gorm:"column:print_pickup_status;default:0" json:"print_pickup_status"`
	InternalNote        NullString `gorm:"column:internal_note" json:"internal_note"`        //
	CreatedBy           string     `gorm:"column:created_by;default:NULL" json:"created_by"` //
	CreatedAt           *time.Time `gorm:"column:created_at" json:"created_at"`              //
	UpdatedBy           string     `gorm:"column:updated_by;default:NULL" json:"updated_by"` //
	UpdatedAt           *time.Time `gorm:"column:updated_at" json:"updated_at"`              //
}

type OrderDataModelPreload struct {
	ID                  string                     `gorm:"column:id;primary_key" json:"id"`                           //
	OrderInvoiceID      string                     `gorm:"column:order_invoice_id" json:"order_invoice_id"`           //
	OrderID             string                     `gorm:"column:orderid" json:"orderid"`                             //
	StatusOrder         int                        `gorm:"column:status_order" json:"status_order"`                   //
	ShortCode           string                     `gorm:"column:short_code" json:"short_code"`                       //
	Notes               NullString                 `gorm:"column:notes" json:"notes"`                                 //
	OnHold              int                        `gorm:"column:on_hold" json:"on_hold"`                             //
	TotalPriceDiscount  float64                    `gorm:"column:total_price_discount" json:"total_price_discount"`   //
	TotalPrice          float64                    `gorm:"column:total_price" json:"total_price"`                     //
	TotalRetailPriceDm  float64                    `gorm:"column:total_retail_price_dm" json:"total_retail_price_dm"` //
	DeliveryFee         float64                    `gorm:"column:delivery_fee" json:"delivery_fee"`                   //
	DeliveryDiscount    float64                    `gorm:"column:delivery_discount" json:"delivery_discount"`         //
	PromoValue          float64                    `gorm:"column:promo_value" json:"promo_value"`                     //
	PromoDelivery       float64                    `gorm:"column:promo_delivery" json:"promo_delivery"`               //
	GrandTotal          float64                    `gorm:"column:grand_total" json:"grand_total"`                     //
	PromoCode           NullString                 `gorm:"column:promo_code" json:"promo_code"`                       //
	PromoName           NullString                 `gorm:"column:promo_name" json:"promo_name"`                       //
	PromoInfo           NullString                 `gorm:"column:promo_info" json:"promo_info"`                       //
	OrderProductInfo    []OrderProductModelPreload `gorm:"->;foreignKey:OrderDataID;references:ID" json:"order_product_info"`
	OrderDeliveryInfo   OrderDeliveryModelPreload  `gorm:"->;foreignKey:OrderDataID;references:ID" json:"order_delivery_info"`
	OrderInvoiceInfo    OrderInvoiceModelPreload   `gorm:"->;foreignKey:ID;references:OrderInvoiceID" json:"order_invoice_info"`
	PrintDeliveryStatus int                        `gorm:"column:print_delivery_status;default:0" json:"print_delivery_status"`
	PrintPickupStatus   int                        `gorm:"column:print_pickup_status;default:0" json:"print_pickup_status"`
	InternalNote        NullString                 `gorm:"column:internal_note" json:"internal_note"` //
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
func (c *OrderDataModel) AfterCreate(tx *gorm.DB) (err error) {
	// var m []OrderDataModel
	// tx.Model(&m).Where("id=?", c.ID).Find(&m)
	// ArangoDbInsert("order_data", m)
	return
}

func (c *OrderDataModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (c *OrderDataModel) AfterUpdate(tx *gorm.DB) (err error) {
	// var m []OrderDataModel
	// tx.Raw("SELECT * FROM order_data ? ", tx.Statement.Clauses["WHERE"]).Find(&m)
	// for _, v := range m {
	// 	ArangoDbUpdate("order_data", fmt.Sprintf("u.id=='%v'", v.ID), v)
	// }
	return
}
func (c *OrderDataModel) AfterDelete(tx *gorm.DB) (err error) {
	return
}

func (c *OrderDataModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}
