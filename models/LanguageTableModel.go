package models

import (
	"time"

	"gorm.io/gorm"
)

type LanguageTableModel struct {
	ID            string     `gorm:"column:id;primary_key" json:"id"`
	Word          NullString `gorm:"column:word;default:null" json:"word"`
	CountryId     string     `gorm:"column:country_id;default:null" json:"country_id"`
	TableTarget   NullString `gorm:"column:table_target;default:null" json:"table_target"`
	TableTargetID string     `gorm:"column:table_target_id;default:null" json:"table_target_id"`
	ColumnName    string     `gorm:"column:column_name;default:null" json:"column_name"`
	CreatedBy     NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt     time.Time  `gorm:"column:created_at" sql:"type:timestamp without time zone" json:"crated_at"`
	UpdatedBy     NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt     time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
}

type LanguageTableModelPreload struct {
	ID            string              `gorm:"column:id;primary_key" json:"id"`
	Word          NullString          `gorm:"column:word;default:null" json:"word"`
	CountryId     string              `gorm:"column:country_id;default:null" json:"country_id"`
	TableTarget   NullString          `gorm:"column:table_target;default:null" json:"table_target"`
	TableTargetID string              `gorm:"column:table_target_id;default:null" json:"table_target_id"`
	ColumnName    string              `gorm:"column:column_name;default:null" json:"column_name"`
	Country       CountryModelPreload `gorm:"foreignKey:ID;references:CountryId" json:"country_data"`
}

// TableName sets the insert table name for this struct type
func (p *LanguageTableModel) TableName() string {
	return "language_table"
}
func (p *LanguageTableModelPreload) TableName() string {
	return "language_table"
}

func (p *LanguageTableModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *LanguageTableModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *LanguageTableModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
