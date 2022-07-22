package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomPageModel struct {
	ID         string     `gorm:"column:id;primary_key" json:"id"`
	Name       string     `gorm:"column:name" json:"name"`
	Active     int        `gorm:"column:active" json:"active"`
	Created_By NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated_By NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type CustomPageModelPreload struct {
	ID     string                        `gorm:"column:id;primary_key" json:"id"`
	Name   string                        `gorm:"column:name" json:"name"`
	Active int                           `gorm:"column:active" json:"active"`
	Image  []CustomPageImageModelPreload `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CustomPageID;references:ID" json:"image"`
}

type CustomPageModelFrontend struct {
	Display_Text string `gorm:"column:display_text;default:null" json:"display_text"`
}

// TableName sets the insert table name for this struct type
func (c *CustomPageModel) TableName() string {
	return "mkt_custom_page"
}

func (c *CustomPageModelPreload) TableName() string {
	return "mkt_custom_page"
}

func (p *CustomPageModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *CustomPageModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *CustomPageModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *CustomPageModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type CustomPageModelContentColumnSave struct {
	ColumnName string `json:"column_name" binding:"required"`
	Content    string `json:"content"`
}

type CustomPageModelContentSave struct {
	CountryID       string `json:"country_id" binding:"required"`
	Word            string `json:"content"`
	Keyword         string `json:"keyword"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	SEOFooter       string `json:"seo_footer"`
	Slug            string `json:"slug"`
}
type CustomPageModelSave struct {
	Name    string                       `json:"name" binding:"required"`
	Active  int                          `json:"active"`
	Image   []CustomPageImageModelInput  `json:"image" binding:"dive"`
	Content []CustomPageModelContentSave `json:"content" binding:"dive"`
}
type CustomPageModelUpdate struct {
	ID      string                       `json:"id" binding:"required"`
	Name    string                       `json:"name" binding:"required"`
	Active  int                          `json:"active"`
	Image   []CustomPageImageModelInput  `json:"image" binding:"dive"`
	Content []CustomPageModelContentSave `json:"content" binding:"dive"`
}

type CustomPageModelDelete struct {
	Language_Id string `json:"id" binding:"required"`
}

type CustomPageModelInfo struct {
	ID string `json:"id" binding:"required"`
}
