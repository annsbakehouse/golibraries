package models

import (
	"time"

	"gorm.io/gorm"
)

type ContactPersonModel struct {
	ID                  string     `gorm:"column:id;primary_key" json:"id"`
	CustomerRetailId    string     `gorm:"column:customer_retail_id;" json:"customer_retail_id"`
	CustomerCorporateId string     `gorm:"column:customer_corporate_id;" json:"customer_corporate_id"`
	DefaultSenderName   NullString `gorm:"column:default_sender_name" json:"default_sender_name"`
	PhoneCountryId      NullString `gorm:"column:phone_country_id" json:"phone_country_id"`
	PhoneCountryCode    NullString `gorm:"column:phone_country_code" json:"phone_country_code"`
	Phone               NullString `gorm:"column:phone" json:"phone"`
	DefaultSenderPhone  NullString `gorm:"column:default_sender_phone" json:"default_sender_phone"`
	Active              uint64     `gorm:"column:active" json:"active"`
	CreatedBy           NullString `gorm:"column:created_by" json:"created_by"` //
	CreatedAt           time.Time  `gorm:"column:created_at" json:"created_at"` //
	UpdatedBy           NullString `gorm:"column:updated_by" json:"updated_by"` //
	UpdatedAt           time.Time  `gorm:"column:updated_at" json:"updated_at"` //
	DeletedBy           NullString `gorm:"column:deleted_by" json:"deleted_by"` //
}

type ContactPersonFormSave struct {
	CustomerCorporateId string `json:"customer_corporate_id" binding:"required"`
	CustomerRetailID    string `json:"customer_retail_id" binding:"required"`
	DefaultSenderName   string `json:"default_sender_name"`
	// PhoneCountryId      string `json:"phone_country_id"`
	// PhoneCountryCode    string `json:"phone_country_code"`
	// Phone               string `json:"phone"`
	DefaultSenderPhone string `json:"default_sender_phone"`
}

type ContactPersonPreload struct {
	ID                  string     `gorm:"column:id;primary_key" json:"id"`
	CustomerRetailId    string     `gorm:"column:customer_retail_id" json:"customer_retail_id"`
	CustomerCorporateId string     `gorm:"column:customer_corporate_id" json:"customer_corporate_id"`
	DefaultSenderName   NullString `gorm:"column:default_sender_name" json:"default_sender_name"`
	// PhoneCountryID            NullString                        `gorm:"column:phone_country_id" json:"phone_country_id"`
	// PhoneCountryCode          NullString                        `gorm:"column:phone_country_code" json:"phone_country_code"`
	// Phone                     NullString                        `gorm:"column:phone" json:"phone"`
	DefaultSenderPhone        NullString                        `gorm:"column:default_sender_phone" json:"default_sender_phone"`
	Active                    uint64                            `gorm:"column:active" json:"active"`
	ContactCustomerRetailInfo CustomerRetailContactModelPreload `gorm:"foreignKey:CustomerRetailId;references:ID" json:"contact_customer_retail_info"`
}

type ContactPersonCustomPreload struct {
	ID                        string                                  `gorm:"column:id;primary_key" json:"id"`
	CustomerRetailId          string                                  `gorm:"column:customer_retail_id" json:"customer_retail_id"`
	CustomerCorporateId       string                                  `gorm:"column:customer_corporate_id" json:"customer_corporate_id"`
	DefaultSenderName         NullString                              `gorm:"column:default_sender_name" json:"default_sender_name"`
	DefaultSenderPhone        NullString                              `gorm:"column:default_sender_phone" json:"default_sender_phone"`
	Active                    uint64                                  `gorm:"column:active" json:"active"`
	ContactCustomerRetailInfo CustomerRetailContactCustomModelPreload `gorm:"foreignKey:CustomerRetailId;references:ID" json:"contact_customer_retail_info"`
}

func (c *ContactPersonModel) TableName() string {
	return "contact_person"
}

func (c *ContactPersonPreload) TableName() string {
	return "contact_person"
}

func (c *ContactPersonCustomPreload) TableName() string {
	return "contact_person"
}

func (c *ContactPersonModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *ContactPersonModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *ContactPersonModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}

type ContactPersonSetActiveForm struct {
	ID     string `json:"id" binding:"required"`
	Active uint64 `json:"active"`
}
