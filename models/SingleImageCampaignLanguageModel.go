package models

import (
	"time"

	"gorm.io/gorm"
)

type SingleImageCampaignLanguageModel struct {
	ID                     string     `gorm:"column:id;primary_key" json:"id"`
	Country_Id             string     `gorm:"column:country_id" json:"country_id"`
	SingleImageCampaign_id string     `gorm:"column:single_image_campaign_id" json:"single_image_campaign_id"`
	Title                  NullString `gorm:"column:title" json:"title"`
	Content                NullString `gorm:"column:content" json:"content"`
	ImgAlt                 NullString `gorm:"column:img_alt" json:"img_alt"`
	LandingPage            NullString `gorm:"column:landing_page" json:"landing_page"`
	ButtonText             NullString `gorm:"column:button_text" json:"button_text"`
	Created_By             NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt              time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated_By             NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt              time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type SingleImageCampaignLanguageModelPreload struct {
	ID                     string     `gorm:"column:id;primary_key" json:"id"`
	Country_Id             string     `gorm:"column:country_id" json:"country_id"`
	SingleImageCampaign_id string     `gorm:"column:single_image_campaign_id" json:"single_image_campaign_id"`
	Title                  NullString `gorm:"column:title" json:"title"`
	Content                NullString `gorm:"column:content" json:"content"`
	ImgAlt                 NullString `gorm:"column:img_alt" json:"img_alt"`
	LandingPage            NullString `gorm:"column:landing_page" json:"landing_page"`
	ButtonText             NullString `gorm:"column:button_text" json:"button_text"`
}

// TableName sets the insert table name for this struct type
func (c *SingleImageCampaignLanguageModel) TableName() string {
	return "mkt_single_image_campaign_language"
}

func (c *SingleImageCampaignLanguageModelPreload) TableName() string {
	return "mkt_single_image_campaign_language"
}

func (p *SingleImageCampaignLanguageModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *SingleImageCampaignLanguageModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *SingleImageCampaignLanguageModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *SingleImageCampaignLanguageModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type SingleImageCampaignLanguageModelSave struct {
	ID                     string `json:"id"`
	Country_Id             string `json:"country_id" binding:"required"`
	SingleImageCampaign_id string `json:"single_image_campaign_id"`
	Title                  string `json:"title" binding:"required"`
	Content                string `json:"content"`
	ImgAlt                 string `json:"img_alt"`
	LandingPage            string `json:"landing_page"`
	ButtonText             string `json:"button_text"`
}
type SingleImageCampaignLanguageModelUpdate struct {
	ID                     string `json:"id" binding:"required"`
	Country_Id             string `json:"country_id" binding:"required"`
	SingleImageCampaign_id string `json:"single_image_campaign_id"`
	Title                  string `json:"title" binding:"required"`
	Content                string `json:"content"`
	ImgAlt                 string `json:"img_alt"`
	LandingPage            string `json:"landing_page"`
	ButtonText             string `json:"button_text"`
}

type SingleImageCampaignLanguageModelDelete struct {
	Language_Id string `json:"id" binding:"required"`
}
