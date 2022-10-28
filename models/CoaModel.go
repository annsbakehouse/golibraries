package models

import (
	"time"

	"gorm.io/gorm"
)

type CoaModel struct {
	ID             string     `gorm:"column:id;primary_key" json:"id"`
	CoaType        int        `gorm:"column:coa_type;" json:"coa_type"`
	Name           NullString `gorm:"column:name" json:"name"`
	NormalPosition int        `gorm:"column:normal_position" json:"normal_position"`
	Active         int        `gorm:"column:active" json:"active"`
	CoaCode        NullString `gorm:"column:coa_code" json:"coa_code"`
	ParentCoaID    NullString `gorm:"column:parent_coa_id" json:"parent_coa_id"`
	Ordering       int        `gorm:"column:ordering" json:"ordering"`
	AccountLevel   int        `gorm:"column:account_level" json:"account_level"`
	CreatedBy      NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt      time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy      NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt      time.Time  `gorm:"column:updated_at" json:"updated_at"`
}
type CoaModelPreload struct {
	ID             string           `gorm:"column:id;primary_key" json:"id"`
	CoaType        int              `gorm:"column:coa_type;" json:"coa_type"`
	Name           NullString       `gorm:"column:name" json:"name"`
	NormalPosition int              `gorm:"column:normal_position" json:"normal_position"`
	Active         int              `gorm:"column:active" json:"active"`
	CoaCode        NullString       `gorm:"column:coa_code" json:"coa_code"`
	ParentCoaID    NullString       `gorm:"column:parent_coa_id" json:"parent_coa_id"`
	Ordering       int              `gorm:"column:ordering" json:"ordering"`
	AccountLevel   int              `gorm:"column:account_level" json:"account_level"`
	ParentInfo     CoaModelNested   `gorm:"foreignKey:ID;references:ParentCoaID" json:"parent_info"`
	ChildInfo      []CoaModelNested `gorm:"foreignKey:ParentCoaID;references:ID" json:"child_info"`
}
type CoaModelNested struct {
	ID             string     `gorm:"column:id;primary_key" json:"id"`
	CoaType        int        `gorm:"column:coa_type;" json:"coa_type"`
	Name           NullString `gorm:"column:name" json:"name"`
	NormalPosition int        `gorm:"column:normal_position" json:"normal_position"`
	Active         int        `gorm:"column:active" json:"active"`
	CoaCode        NullString `gorm:"column:coa_code" json:"coa_code"`
	ParentCoaID    NullString `gorm:"column:parent_coa_id" json:"parent_coa_id"`
	Ordering       int        `gorm:"column:ordering" json:"ordering"`
	AccountLevel   int        `gorm:"column:account_level" json:"account_level"`
}

// TableName sets the insert table name for this struct type
func (p *CoaModel) TableName() string {
	return "coa"
}
func (p *CoaModelPreload) TableName() string {
	return "coa"
}
func (p *CoaModelNested) TableName() string {
	return "coa"
}

func (p *CoaModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *CoaModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

type CoaInputSave struct {
	CoaType        int    `json:"coa_type"`
	Name           string `json:"name" binding:"required"`
	NormalPosition int    `json:"normal_position" binding:"required"`
	Active         int    `json:"active"`
	CoaCode        string `json:"coa_code" binding:"required"`
	ParentCoaID    string `json:"parent_coa_id"`
	Ordering       int    `json:"ordering" binding:"required"`
	AccountLevel   int    `json:"account_level" binding:"required"`
}

type CoaInputUpdate struct {
	ID             string `json:"id" binding:"required"`
	CoaType        int    `json:"coa_type"`
	Name           string `json:"name" binding:"required"`
	NormalPosition int    `json:"normal_position" binding:"required"`
	Active         int    `json:"active"`
	CoaCode        string `json:"coa_code" binding:"required"`
	ParentCoaID    string `json:"parent_coa_id"`
	Ordering       int    `json:"ordering" binding:"required"`
	AccountLevel   int    `json:"account_level" binding:"required"`
}

type CoaInputInfo struct {
	ID string `json:"id" binding:"required"`
}
