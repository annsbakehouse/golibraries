package models

import (
	"gorm.io/gorm"
	"time"
)

type ProvinceModel struct {
	ID           string     `gorm:"column:id;primary_key" json:"id"`
	Name         string     `gorm:"column:name;default:null" json:"name"`
	ProvinceCode string     `gorm:"column:province_code;default:null" json:"provincecode"`
	Active       uint64     `gorm:"column:active;default:1" json:"active"`
	CountryId    string     `gorm:"column:country_id;default:null" json:"countryid"`
	CreatedBy    NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy    NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (p *ProvinceModel) TableName() string {
	return "province"
}

func (p *ProvinceModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProvinceModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *ProvinceModel) FindAllProvinces(tx *gorm.DB) (*[]ProvinceModel, error) {
	var err error
	provinces := []ProvinceModel{}
	err = tx.Debug().Model(&ProvinceModel{}).Limit(100).Find(&provinces).Error
	if err != nil {
		return &[]ProvinceModel{}, err
	}
	return &provinces, err
}

// func (p *ProvinceModel) BeforeDelete(tx *gorm.DB) (err error) {
// 	tt := time.Now()
// 	tFormat := tt.Format("2006-01-02 15:04:05")
// 	p.DeletedAt = NullStringInput(tFormat)
// 	return
// }
