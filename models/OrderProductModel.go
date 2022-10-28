package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderProductModel struct {
	ID                 string     `gorm:"column:id;primary_key" json:"id"`                           //
	OrderInvoiceID     string     `gorm:"column:order_invoice_id" json:"order_invoice_id"`           //
	OrderDataID        string     `gorm:"column:order_data_id" json:"order_data_id"`                 //
	ProductTypeID      string     `gorm:"column:product_type_id" json:"product_type_id"`             //
	ProductModelDataID string     `gorm:"column:product_model_data_id" json:"product_model_data_id"` //
	ProductID          string     `gorm:"column:product_id" json:"product_id"`                       //
	ProductModelID     string     `gorm:"column:product_model_id" json:"product_model_id"`           //
	Name               string     `gorm:"column:name" json:"name"`                                   //
	Text               NullString `gorm:"column:text" json:"text"`                                   //
	Price              float64    `gorm:"column:price" json:"price"`                                 //
	PriceRetail        float64    `gorm:"column:price_retail" json:"price_retail"`                   //
	Discount           float64    `gorm:"column:discount" json:"discount"`                           //
	Note               NullString `gorm:"column:note" json:"note"`                                   //
	INC                int        `gorm:"column:inc" json:"inc"`                                     //
	W2                 NullString `gorm:"column:w2" json:"w2"`                                       //
	W3                 NullString `gorm:"column:w3" json:"w3"`                                       //
	W4                 NullString `gorm:"column:w4" json:"w4"`                                       //
	W2Info             NullString `gorm:"column:w2_info" json:"w2_info"`                             //
	W3Info             NullString `gorm:"column:w3_info" json:"w3_info"`                             //
	W4Info             NullString `gorm:"column:w4_info" json:"w4_info"`                             //
	ProductInfo        NullString `gorm:"column:product_info" json:"product_info"`                   //
	ModelInfo          NullString `gorm:"column:model_info" json:"model_info"`                       //
	Qty                int        `gorm:"column:qty" json:"qty"`                                     //
	CourierSubtypeID   NullString `gorm:"column:courier_subtype_id" json:"courier_subtype_id"`       //
	CreatedBy          string     `gorm:"column:created_by" json:"created_by"`                       //
	CreatedAt          *time.Time `gorm:"column:created_at" json:"created_at"`                       //
	UpdatedBy          string     `gorm:"column:updated_by" json:"updated_by"`                       //
	UpdatedAt          *time.Time `gorm:"column:updated_at" json:"updated_at"`                       //
}

type OrderProductModelPreload struct {
	ID                     string                           `gorm:"column:id;primary_key" json:"id"`                           //
	OrderInvoiceID         string                           `gorm:"column:order_invoice_id" json:"order_invoice_id"`           //
	OrderDataID            string                           `gorm:"column:order_data_id" json:"order_data_id"`                 //
	ProductTypeID          string                           `gorm:"column:product_type_id" json:"product_type_id"`             //
	ProductModelDataID     string                           `gorm:"column:product_model_data_id" json:"product_model_data_id"` //
	ProductID              string                           `gorm:"column:product_id" json:"product_id"`                       //
	ProductModelID         string                           `gorm:"column:product_model_id" json:"product_model_id"`           //
	Name                   string                           `gorm:"column:name" json:"name"`                                   //
	Text                   NullString                       `gorm:"column:text" json:"text"`                                   //
	Price                  float64                          `gorm:"column:price" json:"price"`                                 //
	PriceRetail            float64                          `gorm:"column:price_retail" json:"price_retail"`                   //
	Discount               float64                          `gorm:"column:discount" json:"discount"`                           //
	Note                   NullString                       `gorm:"column:note" json:"note"`                                   //
	INC                    int                              `gorm:"column:inc" json:"inc"`                                     //
	W2                     NullString                       `gorm:"column:w2" json:"w2"`                                       //
	W3                     NullString                       `gorm:"column:w3" json:"w3"`                                       //
	W4                     NullString                       `gorm:"column:w4" json:"w4"`                                       //
	W2Info                 NullString                       `gorm:"column:w2_info" json:"w2_info"`                             //
	W3Info                 NullString                       `gorm:"column:w3_info" json:"w3_info"`                             //
	W4Info                 NullString                       `gorm:"column:w4_info" json:"w4_info"`                             //
	ProductInfo            NullString                       `gorm:"column:product_info" json:"product_info"`                   //
	ModelInfo              NullString                       `gorm:"column:model_info" json:"model_info"`                       //
	Qty                    int                              `gorm:"column:qty" json:"qty"`                                     //
	CourierSubtypeID       NullString                       `gorm:"column:courier_subtype_id" json:"courier_subtype_id"`       //
	OrderProductNestedInfo []OrderNestedProductModelPreload `gorm:"->;foreignKey:OrderProductID;references:ID" json:"order_product_nested_info"`
}

// TableName sets the insert table name for this struct type
func (c *OrderProductModel) TableName() string {
	return "order_product"
}
func (c *OrderProductModelPreload) TableName() string {
	return "order_product"
}

func (c *OrderProductModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderProductModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *OrderProductModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}
