package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductListingMenuModel struct {
	ID         string     `gorm:"column:id;primary_key" json:"id"`
	Name       NullString `gorm:"column:name" json:"name"`
	Active     int        `gorm:"column:active" json:"active"`
	Ordering   int        `gorm:"column:ordering" json:"ordering"`
	Created_By NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated_By NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type ProductListingMenuModelPreload struct {
	ID                        string                            `gorm:"column:id;primary_key" json:"id"`
	Name                      NullString                        `gorm:"column:name" json:"name"`
	Active                    int                               `gorm:"column:active" json:"active"`
	Ordering                  int                               `gorm:"column:ordering" json:"ordering"`
	ProductListingMenLanguage []ProductListingMenuLanguageModel `gorm:"foreignKey:ProductListingMenuID;references:ID" json:"product_listing_language"`
}

// TableName sets the insert table name for this struct type
func (c *ProductListingMenuModel) TableName() string {
	return "mkt_product_listing_menu"
}

func (c *ProductListingMenuModelPreload) TableName() string {
	return "mkt_product_listing_menu"
}
func (p *ProductListingMenuModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductListingMenuModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductListingMenuModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductListingMenuModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type ProductListingMenuModelSave struct {
	Name                       string                                `json:"name" binding:"required"`
	Active                     int                                   `json:"active"`
	Ordering                   int                                   `json:"ordering" binding:"required"`
	ProductListingMenuLanguage []ProductListingMenuLanguageModelSave `json:"product_listing_language" binding:"dive"`
}
type ProductListingMenuModelUpdate struct {
	ID                         string                                `json:"id" binding:"required"`
	Name                       string                                `json:"name" binding:"required"`
	Active                     int                                   `json:"active"`
	Ordering                   int                                   `json:"ordering" binding:"required"`
	ProductListingMenuLanguage []ProductListingMenuLanguageModelSave `json:"product_listing_language" binding:"dive"`
}

type ProductListingMenuModelInfo struct {
	ID string `json:"id" binding:"required"`
}

type FrontEndProductListingMenuGet struct {
	Country_Id string `json:"country_id" binding:"required"`
}
