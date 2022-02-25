package models

import (
	"time"
	"gorm.io/gorm"
)

type CityModel struct {
	ID         string     `gorm:"column:id;primary_key" json:"id"`
	Name       string     `gorm:"column:name;default:null" json:"name"`
	CityCode   string     `gorm:"column:city_code;default:null" json:"citycode"`
	Active     uint64     `gorm:"column:active;default:1" json:"active"`
	CountryId  string     `gorm:"column:country_id;default:null" json:"country_id"`
	ProvinceId string     `gorm:"column:province_id;default:null" json:"province_id"`
	CreatedBy  NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy  NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	// DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt NullString `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (c *CityModel) TableName() string {
	return "city"
}

func (c *CityModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *CityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *CityModel) FindAllCities(tx *gorm.DB) (*[]CityModel, error) {
	var err error
	cities := []CityModel{}
	err = tx.Debug().Model(&CityModel{}).Limit(100).Find(&cities).Error
	if err != nil {
		return &[]CityModel{}, err
	}
	return &cities, err
}

// func (c *CityModel) BeforeDelete(tx *gorm.DB) (err error) {
// 	tt := time.Now()
// 	tFormat := tt.Format("2006-01-02 15:04:05")
// 	c.DeletedAt = NullStringInput(tFormat)
// 	return
// }
