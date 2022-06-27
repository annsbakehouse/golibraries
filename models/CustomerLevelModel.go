package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerLevelModel struct {
	ID        string                    `gorm:"column:id;primary_key" json:"id"` //
	Name      string                    `gorm:"column:name" json:"name"`         //
	Parent    NullString                `gorm:"column:parent" json:"parent"`     //
	Active    int                       `gorm:"column:active;default:1" json:"active"`
	CreatedBy NullString                `gorm:"column:created_by" json:"created_by"` //
	CreatedAt time.Time                 `gorm:"column:created_at" json:"created_at"` //
	UpdatedBy NullString                `gorm:"column:updated_by" json:"updated_by"` //
	UpdatedAt time.Time                 `gorm:"column:updated_at" json:"updated_at"` //
	IsParent  int                       `gorm:"column:is_parent" json:"is_parent"`   //
	Child     []CustomerLevelModelChild `gorm:"foreignKey:ID;references:Parent" json:"child"`
}

type CustomerLevelModelInfo struct {
	ID         string                    `gorm:"column:id;primary_key" json:"id"` //
	Name       string                    `gorm:"column:name" json:"name"`         //
	Parent     NullString                `gorm:"column:parent" json:"parent"`     //
	Active     int                       `gorm:"column:active;default:1" json:"active"`
	IsParent   int                       `gorm:"column:is_parent" json:"is_parent"` //
	Child      []CustomerLevelModelChild `gorm:"foreignKey:Parent;references:ID" json:"child"`
	ParentInfo CustomerLevelModelChild   `gorm:"foreignKey:ID;references:Parent" json:"parent_info"`
}

type CustomerLevelModelChild struct {
	ID       string `gorm:"column:id;primary_key" json:"id"` //
	Name     string `gorm:"column:name" json:"name"`         //
	Parent   string `gorm:"column:parent" json:"parent"`     //
	Active   int    `gorm:"column:active;default:1" json:"active"`
	IsParent int    `gorm:"column:is_parent" json:"is_parent"` //
}

// TableName sets the insert table name for this struct type
func (c *CustomerLevelModel) TableName() string {
	return "customers_level"
}
func (c *CustomerLevelModelInfo) TableName() string {
	return "customers_level"
}
func (c *CustomerLevelModelChild) TableName() string {
	return "customers_level"
}

func (c *CustomerLevelModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *CustomerLevelModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

// func (c *CustomerLevelModel) BeforeDelete(tx *gorm.DB) (err error) {
// 	tt := time.Now()
// 	tFormat := tt.Format("2006-01-02 15:04:05")
// 	c.DeletedAt = NullStringInput(tFormat)
// 	return
// }

type CustomerLevelByParentForm struct {
	Parent string `json:"parent" binding:"required"`
}

type CustomerLevelData struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ProvinceCode string `json:"province_code"`
	Active       uint64 `json:"active"`
	CountryId    string `json:"country_id"`
	CountryName  string `json:"country_name"`
}

type CustomerLevelInfoForm struct {
	ID string `json:"id" binding:"required"`
}

type CustomerLevelSaveForm struct {
	Name     string `json:"name" binding:"required"`
	Parent   string `json:"parent"`
	Active   int    `json:"active"`
	IsParent int    `json:"is_parent"`
}
type CustomerLevelUpdateForm struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Parent   string `json:"parent"`
	Active   int    `json:"active"`
	IsParent int    `json:"is_parent"`
}
