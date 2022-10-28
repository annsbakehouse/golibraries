package models

type ProductListingSubMenuLanguageModel struct {
	ID                      string              `gorm:"column:id;primary_key" json:"id"`
	Name                    NullString          `gorm:"column:name" json:"name"`
	ImageAlt                NullString          `gorm:"column:image_alt" json:"image_alt"`
	Slug                    NullString          `gorm:"column:slug" json:"slug"`
	Country_Id              string              `gorm:"column:country_id" json:"country_id"`
	ProductListingSubmenuID string              `gorm:"column:mkt_product_listing_submenu_id" json:"mkt_product_listing_submenu_id"`
	ProductListingMenuID    string              `gorm:"column:mkt_product_listing_menu_id" json:"mkt_product_listing_menu_id"`
	CountryInfo             CountryModelPreload `gorm:"foreignKey:Country_Id;references:ID" json:"country_info"`
}

type ProductListingSubMenuLanguageModelFrontend struct {
	Display_Text string `gorm:"column:display_text;default:null" json:"display_text"`
}

// TableName sets the insert table name for this struct type
func (c *ProductListingSubMenuLanguageModel) TableName() string {
	return "mkt_product_listing_submenu_language"
}

type ProductListingSubMenuLanguageModelSave struct {
	ID         string `json:"id"`
	Name       string `json:"name" binding:"required"`
	ImageAlt   string `json:"image_alt" binding:"required"`
	Slug       string `json:"slug" binding:"required"`
	Country_Id string `json:"country_id" binding:"required"`
}

type ProductListingSubMenuLanguageModelInfo struct {
	ID string `json:"id" binding:"required"`
}
