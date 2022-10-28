package models

type ProductListingMenuLanguageModel struct {
	ID                   string `gorm:"column:id;primary_key" json:"id"`
	Name                 string `gorm:"column:name" json:"name"`
	Slug                 string `gorm:"column:slug" json:"slug"`
	Country_Id           string `gorm:"column:country_id" json:"country_id"`
	ProductListingMenuID string `gorm:"column:mkt_product_listing_menu_id" json:"mkt_product_listing_menu_id"`
}

type ProductListingMenuLanguageModelFrontend struct {
	Display_Text string `gorm:"column:display_text;default:null" json:"display_text"`
}

// TableName sets the insert table name for this struct type
func (c *ProductListingMenuLanguageModel) TableName() string {
	return "mkt_product_listing_menu_language"
}

type ProductListingMenuLanguageModelSave struct {
	ID         string `json:"id"`
	Name       string `json:"name" binding:"required"`
	Slug       string `json:"slug" binding:"required"`
	Country_Id string `json:"country_id" binding:"required"`
}

type ProductListingMenuLanguageModelInfo struct {
	ID string `json:"id" binding:"required"`
}
