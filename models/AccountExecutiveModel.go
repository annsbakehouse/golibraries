package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountExecutiveModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`     //
	AeId      string     `gorm:"column:aeid" json:"aeid"`             //
	Name      string     `gorm:"column:name" json:"name"`             //
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"` //
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"` //
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"` //
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"` //
}

type AccountExecutiveModelPreload struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`     //
	AeId      string     `gorm:"column:aeid" json:"aeid"`             //
	Name      string     `gorm:"column:name" json:"name"`             //
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"` //
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"` //
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"` //
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"` //
}

func (c *AccountExecutiveModel) TableName() string {
	return "account_executive"
}

func (c *AccountExecutiveModelPreload) TableName() string {
	return "account_executive"
}

func (c *AccountExecutiveModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *AccountExecutiveModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *AccountExecutiveModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type AccountExecutiveInfoForm struct {
	ID string `json:"id" binding:"required"`
}

type AccountExecutiveSaveForm struct {
	AeId string `json:"aeid" binding:"required"` //
	Name string `json:"name" binding:"required"` //
}

type AccountExecutiveUpdateForm struct {
	Id   string `json:"id" binding:"required"`   //
	AeId string `json:"aeid" binding:"required"` //
	Name string `json:"name" binding:"required"` //
}
