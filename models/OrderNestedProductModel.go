package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderNestedProductModel struct {
	ID                 string     `gorm:"column:id;primary_key" json:"id"`
	OrderInvoiceID     NullString `gorm:"column:order_invoice_id" json:"order_invoice_id"`
	OrderDataID        NullString `gorm:"column:order_data_id" json:"order_data_id"`
	OrderProductID     NullString `gorm:"column:order_product_id" json:"order_product_id"`
	ProductTypeID      NullString `gorm:"column:product_type_id" json:"product_type_id"`
	ProductID          NullString `gorm:"column:product_id" json:"product_id"`
	ProductModelID     NullString `gorm:"column:product_model_id" json:"product_model_id"`
	ProductModelDataID NullString `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	ProductName        NullString `gorm:"column:product_name" json:"product_name"`
	Price              float64    `gorm:"column:price" json:"price"`
	PriceRetail        float64    `gorm:"column:price_retail" json:"price_retail"`
	Discount           float64    `gorm:"column:discount" json:"discount"`
	Status             int        `gorm:"column:status;default:1" json:"status"`
	Qty                int        `gorm:"column:qty" json:"qty"`
	ProductInfo        NullString `gorm:"column:product_info" json:"product_info"`
	ModelInfo          NullString `gorm:"column:model_info" json:"model_info"`
	ProductContent     NullString `gorm:"column:product_content" json:"product_content"`
	CreatedBy          string     `gorm:"column:created_by;default:null;" json:"created_by"`
	CreatedAt          *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy          string     `gorm:"column:updated_by;default:null;" json:"updated_by"`
	UpdatedAt          *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type OrderNestedProductModelPreload struct {
	ID                 string            `gorm:"column:id;primary_key" json:"id"`
	OrderInvoiceID     NullString        `gorm:"column:order_invoice_id" json:"order_invoice_id"`
	OrderDataID        NullString        `gorm:"column:order_data_id" json:"order_data_id"`
	OrderProductID     NullString        `gorm:"column:order_product_id" json:"order_product_id"`
	ProductTypeID      NullString        `gorm:"column:product_type_id" json:"product_type_id"`
	ProductID          NullString        `gorm:"column:product_id" json:"product_id"`
	ProductModelID     NullString        `gorm:"column:product_model_id" json:"product_model_id"`
	ProductModelDataID NullString        `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	ProductName        NullString        `gorm:"column:product_name" json:"product_name"`
	Price              float64           `gorm:"column:price" json:"price"`
	PriceRetail        float64           `gorm:"column:price_retail" json:"price_retail"`
	Discount           float64           `gorm:"column:discount" json:"discount"`
	Status             int               `gorm:"column:status" json:"status"`
	Qty                int               `gorm:"column:qty" json:"qty"`
	ProductInfo        NullString        `gorm:"column:product_info" json:"product_info"`
	ModelInfo          NullString        `gorm:"column:model_info" json:"model_info"`
	ProductContent     NullString        `gorm:"column:product_content" json:"product_content"`
	CreatedAt          *time.Time        `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          *time.Time        `gorm:"column:updated_at" json:"updated_at"`
	BarcodeInfo        OrderBarcodeModel `gorm:"->;foreignKey:OrderNestedProductID;references:ID" json:"barcode_info"`
}

// TableName sets the insert table name for this struct type
func (c *OrderNestedProductModel) TableName() string {
	return "order_nested_product"
}
func (c *OrderNestedProductModelPreload) TableName() string {
	return "order_nested_product"
}

func (c *OrderNestedProductModel) AfterCreate(tx *gorm.DB) (err error) {
	// var m []OrderNestedProductModel
	// tx.Model(&m).Where("id=?", c.ID).Find(&m)
	// ArangoDbInsert("order_nested_product", m)
	return
}
func (c *OrderNestedProductModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (c *OrderNestedProductModel) AfterUpdate(tx *gorm.DB) (err error) {
	// var m []OrderNestedProductModel
	// tx.Raw("SELECT * FROM order_nested_product ? ", tx.Statement.Clauses["WHERE"]).Find(&m)
	// for _, v := range m {
	// 	ArangoDbUpdate("order_nested_product", fmt.Sprintf("u.id=='%v'", v.ID), v)
	// }
	return
}
func (c *OrderNestedProductModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
