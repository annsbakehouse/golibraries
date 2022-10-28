package models

import (
	"time"

	"gorm.io/gorm"
)

type MultiProductCampaignModel struct {
	ID         string     `gorm:"column:id;primary_key" json:"id"`
	Name       NullString `gorm:"column:name" json:"name"`
	Start      string     `gorm:"column:start" json:"start"`
	End        string     `gorm:"column:end" json:"end"`
	Usage      int        `gorm:"column:usage" json:"usage"`
	Active     int        `gorm:"column:active" json:"active"`
	Created_By NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated_By NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type MultiProductCampaignModelPreload struct {
	ID             string                                    `gorm:"column:id;primary_key" json:"id"`
	Name           NullString                                `gorm:"column:name" json:"name"`
	Start          string                                    `gorm:"column:start" json:"start"`
	End            string                                    `gorm:"column:end" json:"end"`
	Usage          int                                       `gorm:"column:usage" json:"usage"`
	Active         int                                       `gorm:"column:active" json:"active"`
	ProductContent []MultiProductCampaignProductModelPreload `gorm:"foreignKey:MultiProductCampaignID;references:ID" json:"product_content"`
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
func (c *MultiProductCampaignModel) TableName() string {
	return "mkt_multi_product_campaign"
}

func (c *MultiProductCampaignModelPreload) TableName() string {
	return "mkt_multi_product_campaign"
}
func (p *MultiProductCampaignModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *MultiProductCampaignModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *MultiProductCampaignModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *MultiProductCampaignModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type MultiProductCampaignModeContent struct {
	CountryID   string `json:"country_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	LandingPage string `json:"landing_page" binding:"required"`
	ButtonText  string `json:"button_text" binding:"required"`
}

type MultiProductCampaignModelSave struct {
	Name           string                                 `json:"name" binding:"required"`
	Start          string                                 `json:"start" binding:"required"`
	End            string                                 `json:"end" binding:"required"`
	Usage          int                                    `json:"usage" binding:"required"`
	Active         int                                    `json:"active"`
	Content        []MultiProductCampaignModeContent      `json:"content" binding:"required,dive"`
	ProductContent []MultiProductCampaignProductModelSave `json:"product_content" binding:"required,dive"`
}

type MultiProductCampaignModelUpdate struct {
	ID             string                                 `json:"id" binding:"required"`
	Name           string                                 `json:"name" binding:"required"`
	Start          string                                 `json:"start" binding:"required"`
	End            string                                 `json:"end" binding:"required"`
	Usage          int                                    `json:"usage" binding:"required"`
	Active         int                                    `json:"active"`
	Content        []MultiProductCampaignModeContent      `json:"content" binding:"required,dive"`
	ProductContent []MultiProductCampaignProductModelSave `json:"product_content" binding:"required,dive"`
}

type MultiProductCampaignModelInfo struct {
	ID string `json:"id" binding:"required"`
}
type MultiProductCampaignModelUsageInfo struct {
	Usage           int    `json:"usage" binding:"required"`
	CountryId       string `json:"country_id" binding:"required"`
	CustomerLevelID string `json:"customer_level_id"`
	Limit           int    `json:"limit" binding:"required"`
}
