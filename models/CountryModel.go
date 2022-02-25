package golibraries/models

import (
	"time"

	"gorm.io/gorm"
)

type CountryModel struct {
	ID             string     `gorm:"column:id;primary_key" json:"id"`
	Name           string     `gorm:"column:name;default:null" json:"name"`
	Description    string     `gorm:"column:description;default:null" json:"description"`
	TwoCodeLetter  string     `gorm:"column:two_code_letter;default:null" json:"twocodeletter"`
	CountryCode    string     `gorm:"column:country_code;default:null" json:"countrycode"`
	Flag           NullString `gorm:"column:flag;default:null" json:"flag"`
	Active         uint64     `gorm:"column:active;default:1" json:"active"`
	EnableContent  uint64     `gorm:"column:enable_content;default:0" json:"enablecontent"`
	CreatedBy      NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy      NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
	DefaultCountry uint64     `gorm:"column:default_country;default:0" json:"defaultcountry"`
	// DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt NullString `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (c *CountryModel) TableName() string {
	return "country"
}

func (c *CountryModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *CountryModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *CountryModel) FindAllCountries(tx *gorm.DB) (*[]CountryModel, error) {
	var err error
	countries := []CountryModel{}
	err = tx.Debug().Model(&CountryModel{}).Limit(100).Find(&countries).Error
	if err != nil {
		return &[]CountryModel{}, err
	}
	return &countries, err
}

// func (c *CountryModel) BeforeDelete(tx *gorm.DB) (err error) {
// 	tt := time.Now()
// 	tFormat := tt.Format("2006-01-02 15:04:05")
// 	c.DeletedAt = NullStringInput(tFormat)
// 	return
// }