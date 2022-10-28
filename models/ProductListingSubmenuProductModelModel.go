package models

type ProductListingSubmenuProductModelModel struct {
	INC                     string `gorm:"column:inc;primary_key" json:"inc"`
	ProductID               string `gorm:"column:product_id" json:"product_id"`
	ProductModelID          string `gorm:"column:product_model_id" json:"product_model_id"`
	ProductModelDataID      string `gorm:"column:product_model_data_id" json:"product_model_data_id"`
	ProductListingSubmenuID string `gorm:"column:mkt_product_listing_submenu_id" json:"mkt_product_listing_submenu_id"`
	ProductListingMenuID    string `gorm:"column:mkt_product_listing_menu_id" json:"mkt_product_listing_menu_id"`
}

type ProductListingSubmenuProductModelModelFrontend struct {
	Display_Text string `gorm:"column:display_text;default:null" json:"display_text"`
}

// TableName sets the insert table name for this struct type
func (c *ProductListingSubmenuProductModelModel) TableName() string {
	return "mkt_product_listing_submenu_product_model"
}

type ProductListingSubmenuProductModelModelSave struct {
	INC                     string `json:"inc"`
	ProductID               string `json:"product_id" binding:"required"`
	ProductModelID          string `json:"product_model_id" binding:"required"`
	ProductModelDataID      string `json:"product_model_data_id" binding:"required"`
	ProductListingSubmenuID string `json:"mkt_product_listing_submenu_id" binding:"required"`
	ProductListingMenuID    string `json:"mkt_product_listing_menu_id" binding:"required"`
}

type ProductListingSubmenuProductModelModelInfo struct {
	ID string `json:"id" binding:"required"`
}
