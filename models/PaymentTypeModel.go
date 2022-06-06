package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentTypeModel struct {
	ID             string     `gorm:"column:id;primary_key" json:"id"`
	CoaID          string     `gorm:"column:coa_id" json:"coa_id"`
	Name           NullString `gorm:"column:name" json:"name"`
	ThirdParty     int        `gorm:"column:third_party" json:"third_party"`
	ThirdPartyName NullString `gorm:"column:third_party_name" json:"third_party_name"`
	SystemTagging  NullString `gorm:"column:system_tagging" json:"system_tagging"`
	Ordering       NullInt64  `gorm:"column:ordering" json:"ordering"`
	Active         int        `gorm:"column:active" json:"active"`
	CreatedBy      NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt      time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy      NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt      time.Time  `gorm:"column:updated_at" json:"updated_at"`
}
type PaymentTypeModelPreload struct {
	ID             string         `gorm:"column:id;primary_key" json:"id"`
	CoaID          string         `gorm:"column:coa_id" json:"coa_id"`
	Name           NullString     `gorm:"column:name" json:"name"`
	ThirdParty     int            `gorm:"column:third_party" json:"third_party"`
	ThirdPartyName NullString     `gorm:"column:third_party_name" json:"third_party_name"`
	SystemTagging  NullString     `gorm:"column:system_tagging" json:"system_tagging"`
	Ordering       NullInt64      `gorm:"column:ordering" json:"ordering"`
	Active         int            `gorm:"column:active" json:"active"`
	CoaInfo        CoaModelNested `gorm:"foreignKey:CoaID;references:ID" json:"coa_info"`
}

// TableName sets the insert table name for this struct type
func (p *PaymentTypeModel) TableName() string {
	return "payment_type"
}
func (p *PaymentTypeModelPreload) TableName() string {
	return "payment_type"
}

func (p *PaymentTypeModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *PaymentTypeModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

type PaymentTypeInputSave struct {
	CoaID          string `json:"coa_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	ThirdParty     int    `json:"third_party"`
	ThirdPartyName string `json:"third_party_name"`
	SystemTagging  string `json:"system_tagging"`
	Ordering       int    `json:"ordering"`
	Active         int    `json:"active"`
}

type PaymentTypeInputUpdate struct {
	ID             string `json:"id" binding:"required"`
	CoaID          string `json:"coa_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	ThirdParty     int    `json:"third_party"`
	ThirdPartyName string `json:"third_party_name"`
	SystemTagging  string `json:"system_tagging"`
	Ordering       int    `json:"ordering"`
	Active         int    `json:"active"`
}

type PaymentTypeInputInfo struct {
	ID string `json:"id" binding:"required"`
}
