package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerCorporateContact struct {
	ID                  string     `gorm:"column:id;primary_key" json:"id"`
	CustomerRetailId    string     `gorm:"column:customer_retail_id;" json:"customer_retail_id"`
	CustomerCorporateId string     `gorm:"column:customer_corporate_id;" json:"customer_corporate_id"`
	Level               int        `gorm:"column:level" json:"level"`
	CreatedBy           NullString `gorm:"column:created_by" json:"created_by"` //
	CreatedAt           time.Time  `gorm:"column:created_at" json:"created_at"` //
}

func (c *CustomerCorporateContact) TableName() string {
	return "customer_corporate_contact"
}

func (c *CustomerCorporateContact) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *CustomerCorporateContact) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *CustomerCorporateContact) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}
