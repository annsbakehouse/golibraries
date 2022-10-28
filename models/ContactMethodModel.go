package models

import (
	"time"

	"gorm.io/gorm"
)

type ContactMethodModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	Name      string     `gorm:"column:name;default:null" json:"name"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
}
type ContactMethodModelPreload struct {
	ID   string `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name;default:null" json:"name"`
}

// TableName sets the insert table name for this struct type
func (c *ContactMethodModel) TableName() string {
	return "contact_method"
}
func (c *ContactMethodModelPreload) TableName() string {
	return "contact_method"
}

func (c *ContactMethodModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *ContactMethodModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

type ContactMethodInput struct {
	Name string `json:"name" binding:"required"`
}

type ContactMethodUpdate struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
type ContactMethodInfo struct {
	ID string `json:"id" binding:"required"`
}
