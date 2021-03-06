package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerRetailModel struct {
	ID                  string     `gorm:"column:id;primary_key" json:"id"` //
	CRID                NullString `gorm:"column:crid" json:"crid"`         //
	Name                NullString `gorm:"column:name" json:"name"`         //
	AltName             NullString `gorm:"column:alt_name" json:"alt_name"` //
	PhoneCountryID1     NullString `gorm:"column:phone_1_country_id" json:"phone_1_country_id"`
	PhoneCountryCode1   NullString `gorm:"column:phone_1_country_code" json:"phone_1_country_code"` //
	Phone1              NullString `gorm:"column:phone_1" json:"phone_1"`                           //
	PhoneCountryID2     NullString `gorm:"column:phone_2_country_id" json:"phone_2_country_id"`     //
	PhoneCountryCode2   NullString `gorm:"column:phone_2_country_code" json:"phone_2_country_code"` //
	Phone2              NullString `gorm:"column:phone_2" json:"phone_2"`                           //
	PhoneNote2          NullString `gorm:"column:phone_2_note" json:"phone_2_note"`                 //
	PhoneCountryID3     NullString `gorm:"column:phone_3_country_id" json:"phone_3_country_id"`     //
	PhoneCountryCode3   NullString `gorm:"column:phone_3_country_code" json:"phone_3_country_code"` //
	Phone3              NullString `gorm:"column:phone_3" json:"phone_3"`                           //
	PhoneNote3          NullString `gorm:"column:phone_3_note" json:"phone_3_note"`                 //
	Birthday            NullString `gorm:"column:birthday" json:"birthday"`                         //
	Gender              int        `gorm:"column:gender" json:"gender"`                             //
	Email               NullString `gorm:"column:email" json:"email"`                               //
	Password            NullString `gorm:"column:password" json:"password"`                         //
	InternalShortNote   NullString `gorm:"column:internal_short_note" json:"internal_short_note"`   //
	InternalLongNote    NullString `gorm:"column:internal_long_note" json:"internal_long_note"`     //
	CustomersLevelID    NullString `gorm:"column:customers_level_id" json:"customers_level_id"`     //
	ValidUntil          NullString `gorm:"column:valid_until" json:"valid_until"`                   //
	INC                 int        `gorm:"<-:false;column:inc" json:"inc"`                          //
	NewsletterSubscribe int        `gorm:"column:newsletter_subscribe" json:"newsletter_subscribe"` //
	AllowLogin          int        `gorm:"column:allow_login" json:"allow_login"`                   //
	CreatedBy           NullString `gorm:"column:created_by" json:"created_by"`                     //
	CreatedAt           time.Time  `gorm:"column:created_at" json:"created_at"`                     //
	UpdatedBy           NullString `gorm:"column:updated_by" json:"updated_by"`                     //
	UpdatedAt           time.Time  `gorm:"column:updated_at" json:"updated_at"`                     //
	DeletedBy           NullString `gorm:"column:deleted_by" json:"deleted_by"`                     //
	// DeletedAt         soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type CustomerRetailModelPreload struct {
	ID                  string                 `gorm:"column:id;primary_key" json:"id"` //
	CRID                NullString             `gorm:"column:crid" json:"crid"`         //
	Name                NullString             `gorm:"column:name" json:"name"`         //
	AltName             NullString             `gorm:"column:alt_name" json:"alt_name"` //
	PhoneCountryID1     NullString             `gorm:"column:phone_1_country_id" json:"phone_1_country_id"`
	PhoneCountryCode1   NullString             `gorm:"column:phone_1_country_code" json:"phone_1_country_code"` //
	Phone1              NullString             `gorm:"column:phone_1" json:"phone_1"`                           //
	PhoneCountryID2     NullString             `gorm:"column:phone_2_country_id" json:"phone_2_country_id"`     //
	PhoneCountryCode2   NullString             `gorm:"column:phone_2_country_code" json:"phone_2_country_code"` //
	Phone2              NullString             `gorm:"column:phone_2" json:"phone_2"`                           //
	PhoneNote2          NullString             `gorm:"column:phone_2_note" json:"phone_2_note"`                 //
	PhoneCountryID3     NullString             `gorm:"column:phone_3_country_id" json:"phone_3_country_id"`     //
	PhoneCountryCode3   NullString             `gorm:"column:phone_3_country_code" json:"phone_3_country_code"` //
	Phone3              NullString             `gorm:"column:phone_3" json:"phone_3"`                           //
	PhoneNote3          NullString             `gorm:"column:phone_3_note" json:"phone_3_note"`                 //
	Birthday            NullString             `gorm:"column:birthday" json:"birthday"`                         //
	Gender              int                    `gorm:"column:gender" json:"gender"`                             //
	Email               NullString             `gorm:"column:email" json:"email"`                               //
	Password            NullString             `gorm:"column:password" json:"password"`                         //
	InternalShortNote   NullString             `gorm:"column:internal_short_note" json:"internal_short_note"`   //
	InternalLongNote    NullString             `gorm:"column:internal_long_note" json:"internal_long_note"`     //
	CustomersLevelID    NullString             `gorm:"column:customers_level_id" json:"customers_level_id"`     //
	ValidUntil          NullString             `gorm:"column:valid_until" json:"valid_until"`                   //
	INC                 int                    `gorm:"<-:false;column:inc" json:"inc"`                          //
	NewsletterSubscribe int                    `gorm:"column:newsletter_subscribe" json:"newsletter_subscribe"` //
	AllowLogin          int                    `gorm:"column:allow_login" json:"allow_login"`                   //
	PhoneCountry1Info   CountryModelPreload    `gorm:"foreignKey:PhoneCountryID1;references:ID" json:"phone1_country_info"`
	PhoneCountry2Info   CountryModelPreload    `gorm:"foreignKey:PhoneCountryID2;references:ID" json:"phone2_country_info"`
	PhoneCountry3Info   CountryModelPreload    `gorm:"foreignKey:PhoneCountryID3;references:ID" json:"phone3_country_info"`
	CustomersLevelInfo  CustomerLevelModelInfo `gorm:"foreignKey:CustomersLevelID;references:ID" json:"customer_level_info"`
}

// TableName sets the insert table name for this struct type
func (c *CustomerRetailModel) TableName() string {
	return "customer_retail"
}
func (c *CustomerRetailModelPreload) TableName() string {
	return "customer_retail"
}

func (c *CustomerRetailModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *CustomerRetailModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *CustomerRetailModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model CustomerRetailModel
	// tx.Model(&model).Where("id=?", c.ID).Update("deleted_by", ActiveUser)
	return
}

type CustomerRetailInfoForm struct {
	ID string `json:"id" binding:"required"`
}

type CustomerRetailSaveForm struct {
	Name              string `json:"name" binding:"required"`                 //
	AltName           string `json:"alt_name"`                                //
	PhoneCountryID1   string `json:"phone_1_country_id" binding:"required"`   //
	PhoneCountryCode1 string `json:"phone_1_country_code" binding:"required"` //
	Phone1            string `json:"phone_1" binding:"required"`              //
	PhoneCountryID2   string `json:"phone_2_country_id"`                      //
	PhoneCountryCode2 string `json:"phone_2_country_code"`                    //
	Phone2            string `json:"phone_2"`                                 //
	PhoneNote2        string `json:"phone_2_note"`                            //
	PhoneCountryID3   string `json:"phone_3_country_id"`                      //
	PhoneCountryCode3 string `json:"phone_3_country_code"`                    //
	Phone3            string `json:"phone_3"`                                 //
	PhoneNote3        string `json:"phone_3_note"`                            //
	Birthday          string `json:"birthday"`                                //
	Gender            int    `json:"gender"`                                  //
	Email             string `json:"email" binding:"required"`                //
	InternalShortNote string `json:"internal_short_note"`                     //
	InternalLongNote  string `json:"internal_long_note"`                      //
}
type CustomerRetailUpdatePersonalForm struct {
	ID                string `json:"id" binding:"required"`
	Name              string `json:"name" binding:"required"`                 //
	AltName           string `json:"alt_name"`                                //
	PhoneCountryID1   string `json:"phone_1_country_id" binding:"required"`   //
	PhoneCountryCode1 string `json:"phone_1_country_code" binding:"required"` //
	Phone1            string `json:"phone_1" binding:"required"`              //
	PhoneCountryID2   string `json:"phone_2_country_id"`                      //
	PhoneCountryCode2 string `json:"phone_2_country_code"`                    //
	Phone2            string `json:"phone_2"`                                 //
	PhoneNote2        string `json:"phone_2_note"`                            //
	PhoneCountryID3   string `json:"phone_3_country_id"`                      //
	PhoneCountryCode3 string `json:"phone_3_country_code"`                    //
	Phone3            string `json:"phone_3"`                                 //
	PhoneNote3        string `json:"phone_3_note"`                            //
	Birthday          string `json:"birthday"`                                //
	Gender            int    `json:"gender"`                                  //
	Email             string `json:"email" binding:"required"`                //
	InternalShortNote string `json:"internal_short_note"`                     //
	InternalLongNote  string `json:"internal_long_note"`                      //
}

type CustomerRetailUpdateCustomerLevelForm struct {
	ID                string `json:"id" binding:"required"`
	CustomersLevelID  string `json:"customers_level_id" binding:"required"` //
	ValidUntil        string `json:"valid_until"  binding:"required"`       //        //                   //
	InternalShortNote string `json:"internal_short_note"`                   //
	InternalLongNote  string `json:"internal_long_note"`                    //
}

type CustomerRetailUpdateLoginForm struct {
	ID                  string `json:"id" binding:"required"`
	Email               string `json:"email" binding:"required"` //
	Password            string `json:"password"`                 //
	AllowLogin          int    `json:"allow_login"`              //
	NewsletterSubscribe int    `json:"newsletter_subscribe"`     //
	InternalShortNote   string `json:"internal_short_note"`      //
	InternalLongNote    string `json:"internal_long_note"`       //
}
