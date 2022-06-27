package models

import (
	"time"

	"gorm.io/gorm"
)

type SliderModel struct {
	ID            string     `gorm:"column:id;primary_key" json:"id"`
	Campaign_Name NullString `gorm:"column:campaign_name;default:null" json:"campaign_name"`
	Active        int        `gorm:"column:active;default:0" json:"active"`
	Display_Start NullString `gorm:"column:display_start;default:null" json:"display_start"`
	Display_Until NullString `gorm:"column:display_until;default:null" json:"display_until"`
	Usage         int        `gorm:"column:usage" json:"usage"`
	Created_By    NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt     time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated_By    NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt     time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type SliderModelPreload struct {
	ID            string                       `gorm:"column:id;primary_key" json:"id"`
	Campaign_Name NullString                   `gorm:"column:campaign_name;default:null" json:"campaign_name"`
	Active        int                          `gorm:"column:active;default:null" json:"active"`
	Display_Start NullString                   `gorm:"column:display_start;default:null" json:"display_start"`
	Display_Until NullString                   `gorm:"column:display_until;default:null" json:"display_until"`
	Usage         int                          `gorm:"column:usage" json:"usage"`
	Display_Text  []SliderLanguageModelPreload `gorm:"foreignKey:slider_id;references:ID" json:"slider_data"`
}

// type SliderModelPreload struct {
// 	ID             string     `gorm:"column:id;primary_key" json:"id"`
// 	Campaign_Name           NullString     `gorm:"column:campaign_name;default:null" json:"campaign_name"`
// 	Active    NullString     `gorm:"column:active;default:null" json:"active"`
// 	Display_Start  NullString     `gorm:"column:display_start;default:null" json:"display_start"`
// 	Display_Until    NullString     `gorm:"column:display_until;default:null" json:"display_until"`
// 	Display_Text      []SliderLanguageModelPreload `gorm:"foreignKey:slider_id;references:ID" json:"display_text"`

// }

// TableName sets the insert table name for this struct type
func (c *SliderModel) TableName() string {
	return "mkt_slider"
}

func (c *SliderModelPreload) TableName() string {
	return "mkt_slider"
}
func (p *SliderModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *SliderModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *SliderModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *SliderModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type SliderModelSave struct {
	Campaign_Name string                    `json:"campaign_name" binding:"required"`
	Active        int                       `json:"active"`
	Display_Start string                    `json:"display_start" binding:"required"`
	Display_Until string                    `json:"display_until" binding:"required"`
	Usage         int                       `json:"usage" binding:"required"`
	SliderData    []SliderLanguageModelSave `json:"slider_data" binding:"dive"`
}
type SliderModelUpdate struct {
	ID            string                      `json:"id" binding:"required"`
	Campaign_Name string                      `json:"campaign_name" binding:"required"`
	Active        int                         `json:"active"`
	Display_Start string                      `json:"display_start" binding:"required"`
	Display_Until string                      `json:"display_until" binding:"required"`
	Usage         int                         `json:"usage" binding:"required"`
	SliderData    []SliderLanguageModelUpdate `json:"slider_data" binding:"dive"`
}

type SliderModelInfo struct {
	ID string `json:"id" binding:"required"`
}

type FrontEndSliderGet struct {
	Country_Id string `json:"country_id" binding:"required"`
}
