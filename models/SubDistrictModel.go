package models

import (
	"time"

	"gorm.io/gorm"
)

type SubDistrictModel struct {
	ID              string     `gorm:"column:id;primary_key" json:"id"`
	Name            string     `gorm:"column:name;default:null" json:"name"`
	SubDistrictCode string     `gorm:"column:subdistrict_code;default:null" json:"subdistrict_code"`
	Active          uint64     `gorm:"column:active;default:1" json:"active"`
	CountryId       string     `gorm:"column:country_id;default:null" json:"country_id"`
	ProvinceId      string     `gorm:"column:province_id;default:null" json:"province_id"`
	CityId          string     `gorm:"column:city_id;default:null" json:"city_id"`
	DistrictId      string     `gorm:"column:district_id;default:null" json:"district_id"`
	ZipCode         string     `gorm:"column:zip_code;default:null" json:"zip_code"`
	CreatedBy       NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt       time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy       NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt       time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type SubDistrictModelPreload struct {
	ID              string `gorm:"column:id;primary_key" json:"id"`
	Name            string `gorm:"column:name;default:null" json:"name"`
	SubDistrictCode string `gorm:"column:subdistrict_code;default:null" json:"subdistrict_code"`
	Active          uint64 `gorm:"column:active;default:1" json:"active"`
	CountryId       string `gorm:"column:country_id;default:null" json:"country_id"`
	ProvinceId      string `gorm:"column:province_id;default:null" json:"province_id"`
	CityId          string `gorm:"column:city_id;default:null" json:"city_id"`
	DistrictId      string `gorm:"column:district_id;default:null" json:"district_id"`
	ZipCode         string `gorm:"column:zip_code;default:null" json:"zip_code"`
}

// info subdistrict
type SubDistrictInfoForm struct {
	ID string `json:"id" binding:"required"`
}

type SubDistrictInfoDetail struct {
	CountryName     string `json:"country_name"`
	CountryId       string `json:"country_id"`
	ProvinceName    string `json:"province_name"`
	ProvinceId      string `json:"province_id"`
	CityName        string `json:"city_name"`
	CityId          string `json:"city_id"`
	DistrictName    string `json:"district_name"`
	DistrictId      string `json:"district_id"`
	Active          uint64 `json:"active"`
	ID              string `json:"id"`
	Name            string `json:"subdistrict_name"`
	SubDistrictCode string `json:"subdistrict_code"`
	ZipCode         string `json:"zip_code"`
}

// TableName sets the insert table name for this struct type
func (s *SubDistrictModel) TableName() string {
	return "subdistrict"
}

func (s *SubDistrictModelPreload) TableName() string {
	return "subdistrict"
}

func (s *SubDistrictModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (s *SubDistrictModel) BeforeUpdate(tx *gorm.DB) (err error) {

	return
}

func (s *SubDistrictModel) FindAllCities(tx *gorm.DB) (*[]SubDistrictModel, error) {
	var err error
	subdistricts := []SubDistrictModel{}
	err = tx.Debug().Model(&SubDistrictModel{}).Limit(100).Find(&subdistricts).Error
	if err != nil {
		return &[]SubDistrictModel{}, err
	}
	return &subdistricts, err
}

// func (s *SubDistrictModel) BeforeDelete(tx *gorm.DB) (err error) {
// 	tt := time.Now()
// 	tFormat := tt.Format("2006-01-02 15:04:05")
// 	s.DeletedAt = NullStringInput(tFormat)
// 	return
// }
