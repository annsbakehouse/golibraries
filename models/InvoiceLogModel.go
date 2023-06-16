package models

import "time"

type InvoiceLogModel struct {
	Id               string     `gorm:"column:id;primary_key" json:"id"`
	OrderInvoiceId   string     `gorm:"column:order_invoice_id" json:"order_invoice_id"`
	Date             NullString `gorm:"column:date" json:"date"`
	OrderInvoiceInfo NullString `gorm:"column:order_invoice_info" json:"order_invoice_info"`
	CreatedAt        *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        *time.Time `gorm:"column:updated_at" json:"updated_at"`
	CreatedBy        string     `gorm:"column:created_by" json:"created_by"`
	UpdatedBy        NullString `gorm:"column:updated_by" json:"updated_by"`
	Title            NullString `gorm:"column:title" json:"title"`
	Description      NullString `gorm:"column:description" json:"description"`
	UserInfo         NullString `gorm:"column:user_info" json:"user_info"`
}

type InvoiceLogModelPreload struct {
	Id               string                   `gorm:"column:id;primary_key" json:"id"`
	OrderInvoiceId   string                   `gorm:"column:order_invoice_id" json:"order_invoice_id"`
	Date             NullString               `gorm:"column:date" json:"date"`
	OrderInvoiceInfo OrderInvoiceModelPreload `gorm:"foreignKey:ID;references:OrderInvoiceId" json:"order_invoice_info"`
	CreatedBy        string                   `gorm:"column:created_by" json:"created_by"`
	UserInfo         UserModel                `gorm:"foreignKey:CreatedBy;references:ID" json:"user_info"`
	CreatedAt        NullString               `gorm:"column:created_at" json:"created_at"`
	Title            NullString               `gorm:"column:title" json:"title"`
	Description      NullString               `gorm:"column:description" json:"description"`
}

func (c *InvoiceLogModel) TableName() string {
	return "order_invoice_log"
}

func (c *InvoiceLogModelPreload) TableName() string {
	return "order_invoice_log"
}
