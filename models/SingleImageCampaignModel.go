package models

import (
	"time"

	"gorm.io/gorm"
)

type SingleImageCampaignModel struct {
	ID            string     `gorm:"column:id;primary_key" json:"id"`
	Campaign_Name NullString `gorm:"column:campaign_name" json:"campaign_name"`
	Active        int        `gorm:"column:active;default:0" json:"active"`
	Display_From  NullString `gorm:"column:display_from" json:"display_from"`
	Display_Until NullString `gorm:"column:display_until" json:"display_until"`
	Usage         int        `gorm:"column:usage" json:"usage"`
	DesktopImage  NullString `gorm:"column:desktop_image" json:"desktop_image"`
	TabletImage   NullString `gorm:"column:tablet_image" json:"tablet_image"`
	MobileImage   NullString `gorm:"column:mobile_image" json:"mobile_image"`
	ImageShowCase NullString `gorm:"column:image_show_case" json:"image_show_case"`
	Created_By    NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt     time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated_By    NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt     time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type SingleImageCampaignModelPreload struct {
	ID                  string                                    `gorm:"column:id;primary_key" json:"id"`
	Campaign_Name       NullString                                `gorm:"column:campaign_name" json:"campaign_name"`
	Active              int                                       `gorm:"column:active;default:0" json:"active"`
	Display_From        NullString                                `gorm:"column:display_from" json:"display_from"`
	Display_Until       NullString                                `gorm:"column:display_until" json:"display_until"`
	Usage               int                                       `gorm:"column:usage" json:"usage"`
	DesktopImage        NullString                                `gorm:"column:desktop_image" json:"desktop_image"`
	TabletImage         NullString                                `gorm:"column:tablet_image" json:"tablet_image"`
	MobileImage         NullString                                `gorm:"column:mobile_image" json:"mobile_image"`
	ImageShowCase       NullString                                `gorm:"column:image_show_case" json:"image_show_case"`
	SingleImageLanguage []SingleImageCampaignLanguageModelPreload `gorm:"foreignKey:SingleImageCampaign_id;references:ID" json:"single_image_language"`
}

// type SingleImageCampaignModelPreload struct {
// 	ID             string     `gorm:"column:id;primary_key" json:"id"`
// 	Campaign_Name           NullString     `gorm:"column:campaign_name;default:null" json:"campaign_name"`
// 	Active    NullString     `gorm:"column:active;default:null" json:"active"`
// 	Display_Start  NullString     `gorm:"column:display_start;default:null" json:"display_start"`
// 	Display_Until    NullString     `gorm:"column:display_until;default:null" json:"display_until"`
// 	Display_Text      []SliderLanguageModelPreload `gorm:"foreignKey:slider_id;references:ID" json:"display_text"`

// }

// TableName sets the insert table name for this struct type
func (c *SingleImageCampaignModel) TableName() string {
	return "mkt_single_image_campaign"
}

func (c *SingleImageCampaignModelPreload) TableName() string {
	return "mkt_single_image_campaign"
}
func (p *SingleImageCampaignModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *SingleImageCampaignModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *SingleImageCampaignModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *SingleImageCampaignModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type SingleImageCampaignModelSave struct {
	Campaign_Name       string                                 `json:"campaign_name"  binding:"required"`
	Active              int                                    `json:"active"`
	Display_From        string                                 `json:"display_from"  binding:"required"`
	Display_Until       string                                 `json:"display_until"  binding:"required"`
	Usage               int                                    `json:"usage"`
	DesktopImage        string                                 `json:"desktop_image"`
	TabletImage         string                                 `json:"tablet_image"`
	MobileImage         string                                 `json:"mobile_image"`
	ImageShowCase       string                                 `json:"image_show_case"`
	SingleImageLanguage []SingleImageCampaignLanguageModelSave `json:"single_image_language" binding:"dive"`
}
type SingleImageCampaignModelUpdate struct {
	ID                  string                                 `json:"id" binding:"required"`
	Campaign_Name       string                                 `json:"campaign_name"`
	Active              int                                    `json:"active"  binding:"omitempty"`
	Display_From        string                                 `json:"display_from"  binding:"required"`
	Display_Until       string                                 `json:"display_until"  binding:"required"`
	Usage               int                                    `json:"usage"`
	DesktopImage        string                                 `json:"desktop_image"`
	TabletImage         string                                 `json:"tablet_image"`
	MobileImage         string                                 `json:"mobile_image"`
	ImageShowCase       string                                 `json:"image_show_case"`
	SingleImageLanguage []SingleImageCampaignLanguageModelSave `json:"single_image_language" binding:"dive"`
}

type SingleImageCampaignModelInfo struct {
	ID string `json:"id" binding:"required"`
}
