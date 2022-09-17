package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomPageImageModel struct {
	ID           string     `gorm:"column:id;primary_key" json:"id"`
	CustomPageID string     `gorm:"column:mkt_custom_page_id" json:"mkt_custom_page_id"`
	Image        string     `gorm:"column:image" json:"image"`
	Created_By   NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
}

type CustomPageImageModelPreload struct {
	ID           string    `gorm:"column:id;primary_key" json:"id"`
	CustomPageID string    `gorm:"column:mkt_custom_page_id" json:"mkt_custom_page_id"`
	Image        string    `gorm:"column:image" json:"image"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
}

type CustomPageImageModelFrontend struct {
	Display_Text string `gorm:"column:display_text;default:null" json:"display_text"`
}

// TableName sets the insert table name for this struct type
func (c *CustomPageImageModel) TableName() string {
	return "mkt_custom_page_image"
}

func (c *CustomPageImageModelPreload) TableName() string {
	return "mkt_custom_page_image"
}

func (p *CustomPageImageModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *CustomPageImageModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *CustomPageImageModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *CustomPageImageModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type CustomPageImageModelInput struct {
	ID    string `json:"id"`
	Image string `json:"image"`
}

type CustomPageImageModelDelete struct {
	Language_Id string `json:"id" binding:"required"`
}
