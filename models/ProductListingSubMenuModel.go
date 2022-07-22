package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductListingSubMenuModel struct {
	ID                   string     `gorm:"column:id;primary_key" json:"id"`
	Name                 NullString `gorm:"column:name" json:"name"`
	DisplayFrom          NullString `gorm:"column:display_from" json:"display_from"`
	DisplayUntil         NullString `gorm:"column:display_until" json:"display_until"`
	Active               int        `gorm:"column:active" json:"active"`
	Ordering             int        `gorm:"column:ordering" json:"ordering"`
	DrawerDesktopImage   NullString `gorm:"column:drawer_desktop_image" json:"drawer_desktop_image"`
	ProductListingMenuID string     `gorm:"column:mkt_product_listing_menu_id" json:"mkt_product_listing_menu_id"`
	Created_By           NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt            time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated_By           NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt            time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type ProductListingSubMenuModelPreload struct {
	ID                           string                               `gorm:"column:id;primary_key" json:"id"`
	Name                         NullString                           `gorm:"column:name" json:"name"`
	DisplayFrom                  NullString                           `gorm:"column:display_from" json:"display_from"`
	DisplayUntil                 NullString                           `gorm:"column:display_until" json:"display_until"`
	Active                       int                                  `gorm:"column:active" json:"active"`
	Ordering                     int                                  `gorm:"column:ordering" json:"ordering"`
	DrawerDesktopImage           NullString                           `gorm:"column:drawer_desktop_image" json:"drawer_desktop_image"`
	ProductListingMenuID         string                               `gorm:"column:mkt_product_listing_menu_id" json:"mkt_product_listing_menu_id"`
	ProductListingSubMenLanguage []ProductListingSubMenuLanguageModel `gorm:"foreignKey:ProductListingSubmenuID;references:ID" json:"product_listing_submenu_language"`
}

// TableName sets the insert table name for this struct type
func (c *ProductListingSubMenuModel) TableName() string {
	return "mkt_product_listing_submenu"
}

func (c *ProductListingSubMenuModelPreload) TableName() string {
	return "mkt_product_listing_submenu"
}
func (p *ProductListingSubMenuModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductListingSubMenuModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductListingSubMenuModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductListingSubMenuModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type ProductListingSubMenuModelSave struct {
	Name                          string                                   `json:"name" binding:"required"`
	DisplayFrom                   string                                   `json:"display_from" binding:"required"`
	DisplayUntil                  string                                   `json:"display_until" binding:"required"`
	Active                        int                                      `json:"active"`
	Ordering                      int                                      `json:"ordering" binding:"required"`
	DrawerDesktopImage            string                                   `json:"drawer_desktop_image"`
	ProductListingMenuID          string                                   `json:"mkt_product_listing_menu_id"`
	ProductListingSubMenuLanguage []ProductListingSubMenuLanguageModelSave `json:"product_listing_submenu_language" binding:"dive"`
	//ProductListingSubmenuProductModel []ProductListingSubmenuProductModelModelSave `json:"product_listing_submenu_product_model" binding:"dive"`
}
type ProductListingSubMenuModelUpdate struct {
	ID                            string                                   `json:"id" binding:"required"`
	Name                          string                                   `json:"name" binding:"required"`
	DisplayFrom                   string                                   `json:"display_from" binding:"required"`
	DisplayUntil                  string                                   `json:"display_until" binding:"required"`
	Active                        int                                      `json:"active"`
	Ordering                      int                                      `json:"ordering" binding:"required"`
	DrawerDesktopImage            string                                   `json:"drawer_desktop_image"`
	ProductListingMenuID          string                                   `json:"mkt_product_listing_menu_id"`
	ProductListingSubMenuLanguage []ProductListingSubMenuLanguageModelSave `json:"product_listing_submenu_language" binding:"dive"`
}

type ProductListingSubMenuModelInfo struct {
	ID string `json:"id" binding:"required"`
}
type ProductAllSubMenuModelInfo struct {
	MenuListingID string `json:"menu_listing_id" binding:"required"`
}

type FrontEndProductListingSubMenuGet struct {
	Country_Id string `json:"country_id" binding:"required"`
}
