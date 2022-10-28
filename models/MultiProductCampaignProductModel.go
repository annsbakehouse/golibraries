package models

import (
	"time"

	"gorm.io/gorm"
)

type MultiProductCampaignProductModel struct {
	ID                     string     `gorm:"column:id;primary_key" json:"id"`
	MultiProductCampaignID string     `gorm:"column:mkt_multi_product_campaign_id" json:"mkt_multi_product_campaign_id"`
	ProductID              string     `gorm:"column:product_id" json:"product_id"`
	ProductModelID         string     `gorm:"column:product_model_id" json:"product_model_id"`
	ProductModelDataID     string     `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	Ordering               int        `gorm:"column:ordering" json:"ordering"`
	Created_By             NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt              time.Time  `gorm:"column:created_at" json:"created_at"`
}

type MultiProductCampaignProductModelPreload struct {
	ID                     string                   `gorm:"column:id;primary_key" json:"id"`
	MultiProductCampaignID string                   `gorm:"column:mkt_multi_product_campaign_id" json:"mkt_multi_product_campaign_id"`
	ProductID              string                   `gorm:"column:product_id" json:"product_id"`
	ProductModelID         string                   `gorm:"column:product_model_id" json:"product_model_id"`
	ProductModelDataID     string                   `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	Ordering               int                      `gorm:"column:ordering" json:"ordering"`
	ProductInfo            ProductModelPreload      `gorm:"foreignKey:ID;references:ProductID" json:"product_info"`
	ProductModelInfo       ProductModelModelPreload `gorm:"foreignKey:ID;references:ProductModelID" json:"product_model_info"`
	ProductModelDataInfo   ProductModelDataPreload  `gorm:"foreignKey:ID;references:ProductModelDataID" json:"product_model_data_info"`
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
func (c *MultiProductCampaignProductModel) TableName() string {
	return "mkt_multi_product_campaign_product"
}

func (c *MultiProductCampaignProductModelPreload) TableName() string {
	return "mkt_multi_product_campaign_product"
}
func (p *MultiProductCampaignProductModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *MultiProductCampaignProductModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *MultiProductCampaignProductModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *MultiProductCampaignProductModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type MultiProductCampaignProductModelSave struct {
	ProductID          string `json:"product_id" binding:"required"`
	ProductModelID     string `json:"product_model_id" binding:"required"`
	ProductModelDataID string `json:"product_model_data_id" binding:"required"`
	Ordering           int    `json:"ordering" binding:"required"`
}
type MultiProductCampaignProductModelUpdate struct {
	ID                     string `json:"id" binding:"required"`
	MultiProductCampaignID string `json:"mkt_multi_product_campaign_id"`
	ProductID              string `json:"product_id"`
	ProductModelID         string `json:"product_model_id"`
	ProductModelDataID     string `json:"product_model_data_id"`
	Ordering               int    `json:"ordering"`
}

type MultiProductCampaignProductModelInfo struct {
	ID string `json:"id" binding:"required"`
}
