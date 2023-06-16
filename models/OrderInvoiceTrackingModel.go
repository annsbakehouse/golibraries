package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderInvoiceTrackingModel struct {
	ID                  string     `gorm:"column:id;primary_key" json:"id"`                           //
	OrderInvoiceID      string     `gorm:"column:order_invoice_id" json:"order_invoice_id"`           //
	TrackingTitleID     string     `gorm:"column:tracking_title_id" json:"tracking_title_id"`         //
	TrackingSelectionID string     `gorm:"column:tracking_selection_id" json:"tracking_selection_id"` //
	CreatedBy           string     `gorm:"column:created_by" json:"created_by"`                       //
	CreatedAt           *time.Time `gorm:"column:created_at" json:"created_at"`                       //
	UpdatedBy           string     `gorm:"column:updated_by" json:"updated_by"`                       //
	UpdatedAt           *time.Time `gorm:"column:updated_at" json:"updated_at"`                       //
}

type OrderInvoiceTrackingModelPreload struct {
	ID                    string                             `gorm:"column:id;primary_key" json:"id"`                           //
	OrderInvoiceID        string                             `gorm:"column:order_invoice_id" json:"order_invoice_id"`           //
	TrackingTitleID       string                             `gorm:"column:tracking_title_id" json:"tracking_title_id"`         //
	TrackingSelectionID   string                             `gorm:"column:tracking_selection_id" json:"tracking_selection_id"` //
	TrackingTitleInfo     OrderTrackingTitleModelPreload     `gorm:"->;foreignKey:ID;references:TrackingTitleID" json:"tracking_title_info"`
	TrackingSelectionInfo OrderTrackingSelectionModelPreload `gorm:"->;foreignKey:ID;references:TrackingSelectionID" json:"tracking_selection_info"`
}

// TableName sets the insert table name for this struct type
func (c *OrderInvoiceTrackingModel) TableName() string {
	return "order_invoice_tracking"
}
func (c *OrderInvoiceTrackingModelPreload) TableName() string {
	return "order_invoice_tracking"
}

func (c *OrderInvoiceTrackingModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}
func (c *OrderInvoiceTrackingModel) AfterCreate(tx *gorm.DB) (err error) {
	// var m []OrderInvoiceTrackingModel
	// tx.Model(&m).Where("id=?", c.ID).Find(&m)
	// ArangoDbInsert("order_invoice_tracking", m)
	return
}

func (c *OrderInvoiceTrackingModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (c *OrderInvoiceTrackingModel) AfterUpdate(tx *gorm.DB) (err error) {
	// var m []OrderInvoiceTrackingModel
	// tx.Raw("SELECT * FROM order_invoice_tracking ? ", tx.Statement.Clauses["WHERE"]).Find(&m)
	// for _, v := range m {
	// 	ArangoDbUpdate("order_invoice_tracking", fmt.Sprintf("u.id=='%v'", v.ID), v)
	// }
	return
}

func (c *OrderInvoiceTrackingModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}

type OrderInvoiceTrackingSave struct {
	TrackingTitleID     string `json:"tracking_title_id" binding:"required"`     //
	TrackingSelectionID string `json:"tracking_selection_id" binding:"required"` //
}

type OrderInvoiceTrackingUpdate struct {
	ID                  string `json:"id"`                                       //
	OrderInvoiceID      string `json:"order_invoice_id" binding:"required"`      //
	TrackingTitleID     string `json:"tracking_title_id" binding:"required"`     //
	TrackingSelectionID string `json:"tracking_selection_id" binding:"required"` //
}
