package models

import (
	"time"
	"gorm.io/gorm"
)

type ProductTypeModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`     //
	Name      NullString `gorm:"column:name" json:"name"`             //
}

// TableName sets the insert table name for this struct type
func (p *ProductTypeModel) TableName() string {
	return "product_type"
}
