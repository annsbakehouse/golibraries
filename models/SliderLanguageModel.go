package models

import (
	"time"

	"gorm.io/gorm"
)

type SliderLanguageModel struct {
	ID           string     `gorm:"column:id;primary_key" json:"id"`
	Country_Id   string     `gorm:"column:country_id" json:"country_id"`
	Display_Text NullString `gorm:"column:display_text" json:"display_text"`
	Slider_Id    string     `gorm:"column:slider_id" json:"slider_id"`
	Link         NullString `gorm:"column:link" json:"link"`
	Created_By   NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated_By   NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type SliderLanguageModelPreload struct {
	Language_Id  string     `gorm:"column:id;primary_key" json:"id"`
	Country_Id   string     `gorm:"column:country_id" json:"country_id"`
	Display_Text NullString `gorm:"column:display_text" json:"display_text"`
	Slider_Id    string     `gorm:"column:slider_id" json:"slider_id"`
	Link         string     `gorm:"column:link" json:"link"`
}

type SliderLanguageModelFrontend struct {
	Display_Text string `gorm:"column:display_text;default:null" json:"display_text"`
}

// TableName sets the insert table name for this struct type
func (c *SliderLanguageModel) TableName() string {
	return "mkt_slider_language"
}

func (c *SliderLanguageModelPreload) TableName() string {
	return "mkt_slider_language"
}

func (p *SliderLanguageModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *SliderLanguageModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *SliderLanguageModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *SliderLanguageModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type SliderLanguageModelSave struct {
	Display_Text string `json:"display_text" binding:"required"`
	Link         string `json:"link"`
	Country_Id   string `json:"country_id" binding:"required"`
}
type SliderLanguageModelUpdate struct {
	ID           string `json:"id"`
	Display_Text string `json:"display_text" binding:"required"`
	Link         string `json:"link"`
	Country_Id   string `json:"country_id" binding:"required"`
}

type SliderLanguageModelDelete struct {
	Language_Id string `json:"id" binding:"required"`
}
