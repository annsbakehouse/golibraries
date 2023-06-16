package models

import (
	"time"
)

type BannerPopupModel struct {
	ID         string     `gorm:"column:id;primary_key" json:"id"`
	Name       NullString `gorm:"column:name;default:null" json:"name"`
	Image      string     `gorm:"column:image;" json:"image"`
	DateStart  string     `gorm:"column:date_start" json:"date_start"`
	DateEnd    string     `gorm:"column:date_end" json:"date_end"`
	Active     int        `gorm:"column:active" json:"active"`
	Created_By string     `gorm:"column:created_by" json:"created_by"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated_By string     `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type BannerPopupPreload struct {
	ID         string                      `gorm:"column:id;primary_key" json:"id"`
	Name       NullString                  `gorm:"column:name;default:null" json:"name"`
	Image      string                      `gorm:"column:image;" json:"image"`
	DateStart  string                      `gorm:"column:date_start" json:"date_start"`
	DateEnd    string                      `gorm:"column:date_end" json:"date_end"`
	Active     int                         `gorm:"column:active" json:"active"`
	Created_By string                      `gorm:"column:created_by" json:"created_by"`
	CreatedAt  time.Time                   `gorm:"column:created_at" json:"created_at"`
	Updated_By string                      `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt  time.Time                   `gorm:"column:updated_at" json:"updated_at"`
	Language   []LanguageTableModelPreload `gorm:"->;foreignKey:TableTargetID;references:ID" json:"language"`
}

// TableName sets the insert table name for this struct type
func (c *BannerPopupModel) TableName() string {
	return "mkt_banner_popup"
}

func (c *BannerPopupPreload) TableName() string {
	return "mkt_banner_popup"
}

type BannerPopupLanguage struct {
	CountryId string `json:"country_id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Link      string `json:"link" binding:"required"`
}
type BannerPopupSave struct {
	Name      string                `json:"name" binding:"required"`
	Image     string                `json:"image" binding:"required"`
	DateStart string                `json:"date_start" binding:"required"`
	DateEnd   string                `json:"date_end" binding:"required"`
	Active    int                   `json:"active"`
	Titles    []BannerPopupLanguage `json:"titles" binding:"required"`
}
type BannerPopupUpdate struct {
	ID        string                `json:"id" binding:"required"`
	Name      string                `json:"name" binding:"required"`
	Image     string                `json:"image"`
	DateStart string                `json:"date_start" binding:"required"`
	DateEnd   string                `json:"date_end" binding:"required"`
	Active    int                   `json:"active"`
	Titles    []BannerPopupLanguage `json:"titles" binding:"required"`
}

type BannerPopupInfo struct {
	ID string `json:"id" binding:"required"`
}
