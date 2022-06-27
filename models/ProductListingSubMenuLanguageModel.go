package models

type ProductListingSubMenuLanguageModel struct {
	ID                      string `gorm:"column:id;primary_key" json:"id"`
	Name                    string `gorm:"column:name" json:"name"`
	Description             string `gorm:"column:description" json:"description"`
	MetaKeyword             string `gorm:"column:meta_keyword" json:"meta_keyword"`
	MetaTitle               string `gorm:"column:meta_title" json:"meta_title"`
	MetaDescription         string `gorm:"column:meta_description" json:"meta_description"`
	SeoFooter               string `gorm:"column:seo_footer" json:"seo_footer"`
	ImageAlt                string `gorm:"column:image_alt" json:"image_alt"`
	Slug                    string `gorm:"column:slug" json:"slug"`
	Country_Id              string `gorm:"column:country_id" json:"country_id"`
	ProductListingSubmenuID string `gorm:"column:mkt_product_listing_submenu_id" json:"mkt_product_listing_submenu_id"`
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
	Slug       string `json:"slug" binding:"required"`
	Country_Id string `json:"country_id" binding:"required"`
}

type ProductListingSubMenuLanguageModelInfo struct {
	ID string `json:"id" binding:"required"`
}
