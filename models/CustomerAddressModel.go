package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerAddressModel struct {
	ID                     string     `gorm:"column:id;primary_key" json:"id"`                     //
	Title                  NullString `gorm:"column:title" json:"title"`                           //
	CustomerRetailID       string     `gorm:"column:customer_retail_id" json:"customer_retail_id"` //
	CountryID              string     `gorm:"column:country_id" json:"country_id"`                 //
	ProvinceID             string     `gorm:"column:province_id" json:"province_id"`
	CityID                 string     `gorm:"column:city_id" json:"city_id"`                                     //
	DistrictID             string     `gorm:"column:district_id" json:"district_id"`                             //
	SubdistrictID          string     `gorm:"column:subdistrict_id" json:"subdistrict_id"`                       //
	PostalCode             string     `gorm:"column:postal_code" json:"postal_code"`                             //
	CountryCode            NullString `gorm:"column:country_code" json:"country_code"`                           //
	AddressDetail1         NullString `gorm:"column:address_detail1" json:"address_detail1"`                     //
	AddressDetail2         NullString `gorm:"column:address_detail2" json:"address_detail2"`                     //
	AddressDetail3         NullString `gorm:"column:address_detail3" json:"address_detail3"`                     //
	ReceiverName           NullString `gorm:"column:receiver_name" json:"receiver_name"`                         //
	ReceiverPhone          NullString `gorm:"column:receiver_phone" json:"receiver_phone"`                       //
	ReceiverPhoneCountryID string     `gorm:"column:receiver_phone_country_id" json:"receiver_phone_country_id"` //
	CreatedBy              NullString `gorm:"column:created_by" json:"created_by"`                               //
	CreatedAt              time.Time  `gorm:"column:created_at" json:"created_at"`                               //
	UpdatedBy              NullString `gorm:"column:updated_by" json:"updated_by"`                               //
	UpdatedAt              time.Time  `gorm:"column:updated_at" json:"updated_at"`                               //
}

type CustomerAddressModelPreload struct {
	ID                       string                     `gorm:"column:id;primary_key" json:"id"`                     //
	Title                    NullString                 `gorm:"column:title" json:"title"`                           //
	CustomerRetailID         string                     `gorm:"column:customer_retail_id" json:"customer_retail_id"` //
	CountryID                string                     `gorm:"column:country_id" json:"country_id"`                 //
	ProvinceID               string                     `gorm:"column:province_id" json:"province_id"`
	CityID                   string                     `gorm:"column:city_id" json:"city_id"`                                     //
	DistrictID               string                     `gorm:"column:district_id" json:"district_id"`                             //
	SubdistrictID            string                     `gorm:"column:subdistrict_id" json:"subdistrict_id"`                       //
	PostalCode               string                     `gorm:"column:postal_code" json:"postal_code"`                             //
	CountryCode              NullString                 `gorm:"column:country_code" json:"country_code"`                           //
	AddressDetail1           NullString                 `gorm:"column:address_detail1" json:"address_detail1"`                     //
	AddressDetail2           NullString                 `gorm:"column:address_detail2" json:"address_detail2"`                     //
	AddressDetail3           NullString                 `gorm:"column:address_detail3" json:"address_detail3"`                     //
	ReceiverName             NullString                 `gorm:"column:receiver_name" json:"receiver_name"`                         //
	ReceiverPhone            NullString                 `gorm:"column:receiver_phone" json:"receiver_phone"`                       //
	ReceiverPhoneCountryID   string                     `gorm:"column:receiver_phone_country_id" json:"receiver_phone_country_id"` //
	CustomerRetailInfo       CustomerRetailModelPreload `gorm:"foreignKey:CustomerRetailID;references:ID" json:"customer_retail_info"`
	CountryInfo              CountryModelPreload        `gorm:"foreignKey:CountryID;references:ID" json:"country_info"`
	ReceiverPhoneCountryInfo CountryModelPreload        `gorm:"foreignKey:ReceiverPhoneCountryID;references:ID" json:"receiver_phone_country_info"`
	ProvinceInfo             ProvinceModelPreload       `gorm:"foreignKey:ProvinceID;references:ID" json:"province_info"`
	CityInfo                 CityModelPreload           `gorm:"foreignKey:CityID;references:ID" json:"city_info"`
	DistrictInfo             DistrictModelPreload       `gorm:"foreignKey:DistrictID;references:ID" json:"district_info"`
	SubdistrictInfo          SubDistrictModelPreload    `gorm:"foreignKey:SubdistrictID;references:ID" json:"subdistrict_info"`
}

// TableName sets the insert table name for this struct type
func (c *CustomerAddressModel) TableName() string {
	return "customer_address"
}
func (c *CustomerAddressModelPreload) TableName() string {
	return "customer_address"
}

func (c *CustomerAddressModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *CustomerAddressModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *CustomerAddressModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}

type CustomerAddressModelInfoForm struct {
	ID string `json:"id" binding:"required"`
}
type CustomerAddressByCustomerModelInfoForm struct {
	ID string `json:"customer_id" binding:"required"`
}

type CustomerAddressModelSaveForm struct {
	Title                  string `json:"title" binding:"required"`                     //
	CustomerRetailID       string `json:"customer_retail_id" binding:"required"`        //
	CountryID              string `json:"country_id"`                                   //
	ProvinceID             string `json:"province_id" binding:"required"`               //
	CityID                 string `json:"city_id" binding:"required"`                   //
	DistrictID             string `json:"district_id" binding:"required"`               //
	SubdistrictID          string `json:"subdistrict_id" binding:"required"`            //
	PostalCode             string `json:"postal_code"`                                  //
	CountryCode            string `json:"country_code" binding:"required"`              //
	AddressDetail1         string `json:"address_detail1" binding:"required"`           //
	AddressDetail2         string `json:"address_detail2"`                              //
	AddressDetail3         string `json:"address_detail3"`                              //
	ReceiverName           string `json:"receiver_name" binding:"required"`             //
	ReceiverPhone          string `json:"receiver_phone" binding:"required"`            //
	ReceiverPhoneCountryID string `json:"receiver_phone_country_id" binding:"required"` //
}
type CustomerAddressModellUpdateForm struct {
	ID                     string `json:"id" binding:"required"`                        //
	Title                  string `json:"title" binding:"required"`                     //
	CustomerRetailID       string `json:"customer_retail_id" binding:"required"`        //
	CountryID              string `json:"country_id"`                                   //
	ProvinceID             string `json:"province_id" binding:"required"`               //
	CityID                 string `json:"city_id" binding:"required"`                   //
	DistrictID             string `json:"district_id" binding:"required"`               //
	SubdistrictID          string `json:"subdistrict_id" binding:"required"`            //
	PostalCode             string `json:"postal_code"`                                  //
	CountryCode            string `json:"country_code" binding:"required"`              //
	AddressDetail1         string `json:"address_detail1" binding:"required"`           //
	AddressDetail2         string `json:"address_detail2"`                              //
	AddressDetail3         string `json:"address_detail3"`                              //
	ReceiverName           string `json:"receiver_name" binding:"required"`             //
	ReceiverPhone          string `json:"receiver_phone" binding:"required"`            //
	ReceiverPhoneCountryID string `json:"receiver_phone_country_id" binding:"required"` //
}
