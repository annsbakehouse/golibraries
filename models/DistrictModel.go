package golibraries/models

import (
	"gorm.io/gorm"
	"time"
)

type DistrictModel struct {
	ID           string     `gorm:"column:id;primary_key" json:"id"`
	Name         string     `gorm:"column:name;default:null" json:"name"`
	DistrictCode string     `gorm:"column:district_code;default:null" json:"districtcode"`
	Active       uint64     `gorm:"column:active;default:1" json:"active"`
	CountryId    string     `gorm:"column:country_id;default:null" json:"countryid"`
	ProvinceId   string     `gorm:"column:province_id;default:null" json:"provinceid"`
	CityId       string     `gorm:"column:city_id;default:null" json:"cityid"`
	CreatedBy    NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy    NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	// DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt NullString `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (d *DistrictModel) TableName() string {
	return "district"
}

func (d *DistrictModel) BeforeCreate(tx *gorm.DB) (err error) {
	
	return
}

func (d *DistrictModel) BeforeUpdate(tx *gorm.DB) (err error) {
	
	return
}

func (d *DistrictModel) FindAllCities(tx *gorm.DB) (*[]DistrictModel, error) {
	var err error
	districts := []DistrictModel{}
	err = tx.Debug().Model(&DistrictModel{}).Limit(100).Find(&districts).Error
	if err != nil {
		return &[]DistrictModel{}, err
	}
	return &districts, err
}

// func (d *DistrictModel) BeforeDelete(tx *gorm.DB) (err error) {
// 	tt := time.Now()
// 	tFormat := tt.Format("2006-01-02 15:04:05")
// 	d.DeletedAt = NullStringInput(tFormat)
// 	return
// }
