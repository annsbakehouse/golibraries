package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
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
	Email               NullString `gorm:"column:email;unique" json:"email"`                        //
	ValidEmail          NullString `gorm:"column:valid_email;unique;not null" json:"valid_email"`   //
	Password            NullString `gorm:"column:password" json:"password"`                         //
	InternalShortNote   NullString `gorm:"column:internal_short_note" json:"internal_short_note"`   //
	InternalLongNote    NullString `gorm:"column:internal_long_note" json:"internal_long_note"`     //
	CustomersLevelID    NullString `gorm:"column:customers_level_id" json:"customers_level_id"`     //
	ValidUntil          NullString `gorm:"column:valid_until" json:"valid_until"`                   //
	INC                 int        `gorm:"<-:false;column:inc" json:"inc"`                          //
	NewsletterSubscribe int        `gorm:"column:newsletter_subscribe" json:"newsletter_subscribe"` //
	AllowLogin          int        `gorm:"column:allow_login" json:"allow_login"`                   //
	Secret              string     `gorm:"column:secret" json:"secret"`                             //
	Active              int        `gorm:"column:active" json:"active"`                             //
	ForgotCount         int        `gorm:"column:forgot_count" json:"forgot_count"`                 //
	ForgotTime          time.Time  `gorm:"column:forgot_time" json:"forgot_time"`                   //
	ForgotToken         NullString `gorm:"column:forgot_token" json:"forgot_token"`                 //
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
	Email               NullString             `gorm:"column:email;unique" json:"email"`                        //
	ValidEmail          NullString             `gorm:"column:valid_email;unique;not null" json:"valid_email"`   //
	Password            NullString             `gorm:"column:password" json:"password"`                         //
	InternalShortNote   NullString             `gorm:"column:internal_short_note" json:"internal_short_note"`   //
	InternalLongNote    NullString             `gorm:"column:internal_long_note" json:"internal_long_note"`     //
	CustomersLevelID    NullString             `gorm:"column:customers_level_id" json:"customers_level_id"`     //
	ValidUntil          NullString             `gorm:"column:valid_until" json:"valid_until"`                   //
	INC                 int                    `gorm:"<-:false;column:inc" json:"inc"`                          //
	NewsletterSubscribe int                    `gorm:"column:newsletter_subscribe" json:"newsletter_subscribe"` //
	AllowLogin          int                    `gorm:"column:allow_login" json:"allow_login"`                   //
	Secret              string                 `gorm:"column:secret" json:"secret"`                             //
	Active              int                    `gorm:"column:active" json:"active"`                             //
	ForgotCount         int                    `gorm:"column:forgot_count" json:"forgot_count"`                 //
	ForgotTime          time.Time              `gorm:"column:forgot_time" json:"forgot_time"`                   //
	ForgotToken         NullString             `gorm:"column:forgot_token" json:"forgot_token"`                 //
	PhoneCountry1Info   CountryModelPreload    `gorm:"foreignKey:PhoneCountryID1;references:ID" json:"phone1_country_info"`
	PhoneCountry2Info   CountryModelPreload    `gorm:"foreignKey:PhoneCountryID2;references:ID" json:"phone2_country_info"`
	PhoneCountry3Info   CountryModelPreload    `gorm:"foreignKey:PhoneCountryID3;references:ID" json:"phone3_country_info"`
	CustomersLevelInfo  CustomerLevelModelInfo `gorm:"foreignKey:CustomersLevelID;references:ID" json:"customer_level_info"`
}

type CustomerRetailContactModelPreload struct {
	ID                string     `gorm:"column:id;primary_key" json:"id"` //
	CRID              NullString `gorm:"column:crid" json:"crid"`         //
	Name              NullString `gorm:"column:name" json:"name"`         //
	AltName           NullString `gorm:"column:alt_name" json:"alt_name"` //
	PhoneCountryID1   NullString `gorm:"column:phone_1_country_id" json:"phone_1_country_id"`
	PhoneCountryCode1 NullString `gorm:"column:phone_1_country_code" json:"phone_1_country_code"` //
	Phone1            NullString `gorm:"column:phone_1" json:"phone_1"`                           //
	PhoneCountryID2   NullString `gorm:"column:phone_2_country_id" json:"phone_2_country_id"`     //
	PhoneCountryCode2 NullString `gorm:"column:phone_2_country_code" json:"phone_2_country_code"` //
	Phone2            NullString `gorm:"column:phone_2" json:"phone_2"`                           //
	PhoneNote2        NullString `gorm:"column:phone_2_note" json:"phone_2_note"`                 //
	PhoneCountryID3   NullString `gorm:"column:phone_3_country_id" json:"phone_3_country_id"`     //
	PhoneCountryCode3 NullString `gorm:"column:phone_3_country_code" json:"phone_3_country_code"` //
	Phone3            NullString `gorm:"column:phone_3" json:"phone_3"`                           //
	PhoneNote3        NullString `gorm:"column:phone_3_note" json:"phone_3_note"`                 //
	Birthday          NullString `gorm:"column:birthday" json:"birthday"`                         //
	Gender            int        `gorm:"column:gender" json:"gender"`                             //
	Active            int        `gorm:"column:active" json:"active"`                             //
	Email             NullString `gorm:"column:email;unique" json:"email"`                        //
	ValidEmail        NullString `gorm:"column:valid_email;unique;not null" json:"valid_email"`   //
	Password          NullString `gorm:"column:password" json:"password"`                         //
	InternalShortNote NullString `gorm:"column:internal_short_note" json:"internal_short_note"`   //
	InternalLongNote  NullString `gorm:"column:internal_long_note" json:"internal_long_note"`     //
	CustomersLevelID  NullString `gorm:"column:customers_level_id" json:"customers_level_id"`     //
	ValidUntil        NullString `gorm:"column:valid_until" json:"valid_until"`                   //
}

type CustomerRetailContactCustomModelPreload struct {
	ID                 string                 `gorm:"column:id;primary_key" json:"id"` //
	CRID               NullString             `gorm:"column:crid" json:"crid"`         //
	Name               NullString             `gorm:"column:name" json:"name"`         //
	AltName            NullString             `gorm:"column:alt_name" json:"alt_name"` //
	PhoneCountryID1    NullString             `gorm:"column:phone_1_country_id" json:"phone_1_country_id"`
	PhoneCountryCode1  NullString             `gorm:"column:phone_1_country_code" json:"phone_1_country_code"` //
	Phone1             NullString             `gorm:"column:phone_1" json:"phone_1"`                           //
	PhoneCountryID2    NullString             `gorm:"column:phone_2_country_id" json:"phone_2_country_id"`     //
	PhoneCountryCode2  NullString             `gorm:"column:phone_2_country_code" json:"phone_2_country_code"` //
	Phone2             NullString             `gorm:"column:phone_2" json:"phone_2"`                           //
	PhoneNote2         NullString             `gorm:"column:phone_2_note" json:"phone_2_note"`                 //
	PhoneCountryID3    NullString             `gorm:"column:phone_3_country_id" json:"phone_3_country_id"`     //
	PhoneCountryCode3  NullString             `gorm:"column:phone_3_country_code" json:"phone_3_country_code"` //
	Phone3             NullString             `gorm:"column:phone_3" json:"phone_3"`                           //
	PhoneNote3         NullString             `gorm:"column:phone_3_note" json:"phone_3_note"`                 //
	Birthday           NullString             `gorm:"column:birthday" json:"birthday"`                         //
	Gender             int                    `gorm:"column:gender" json:"gender"`                             //
	Active             int                    `gorm:"column:active" json:"active"`                             //
	Email              NullString             `gorm:"column:email;unique" json:"email"`                        //
	ValidEmail         NullString             `gorm:"column:valid_email;unique;not null" json:"valid_email"`   //
	Password           NullString             `gorm:"column:password" json:"password"`                         //
	InternalShortNote  NullString             `gorm:"column:internal_short_note" json:"internal_short_note"`   //
	InternalLongNote   NullString             `gorm:"column:internal_long_note" json:"internal_long_note"`     //
	CustomersLevelID   NullString             `gorm:"column:customers_level_id" json:"customers_level_id"`     //
	ValidUntil         NullString             `gorm:"column:valid_until" json:"valid_until"`                   //
	CustomersLevelInfo CustomerLevelModelInfo `gorm:"foreignKey:CustomersLevelID;references:ID" json:"customer_level_info"`
}

// TableName sets the insert table name for this struct type
func (c *CustomerRetailModel) TableName() string {
	return "customer_retail"
}
func (c *CustomerRetailModelPreload) TableName() string {
	return "customer_retail"
}
func (c *CustomerRetailContactModelPreload) TableName() string {
	return "customer_retail"
}
func (c *CustomerRetailContactCustomModelPreload) TableName() string {
	return "customer_retail"
}

func GetRandomSecret(tx *gorm.DB) string {
	length := 50
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	secret := fmt.Sprintf("%x", b)[:length]
	var model CustomerRetailModel
	res := tx.Model(&model).Where("secret=?", secret).First(&model)
	if res.RowsAffected == 0 {
		return secret
	}
	return GetRandomSecret(tx)
}
func (c *CustomerRetailModel) BeforeCreate(tx *gorm.DB) (err error) {
	c.Secret = GetRandomSecret(tx)
	if c.Email.Valid {
		c.ValidEmail = c.Email
	} else {
		c.ValidEmail = NullStringInput("emailkosong_" + uuid.New().String() + "@annsbakehouse.com")
	}
	return
}
func (c *CustomerRetailModel) BeforeUpdate(tx *gorm.DB) (err error) {
	if c.Email.Valid {
		c.ValidEmail = c.Email
	}
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
	Email             string `json:"email"`                                   //
	Active            int    `json:"active"`                                  //
	InternalShortNote string `json:"internal_short_note"`                     //
	InternalLongNote  string `json:"internal_long_note"`
	CustomersLevelID  string `json:"customers_level_id" binding:"required"` //
	//
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
	Email             string `json:"email"`                                   //
	Active            int    `json:"active"`                                  //
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
	Email               string `json:"email"`                //
	Password            string `json:"password"`             //
	AllowLogin          int    `json:"allow_login"`          //
	NewsletterSubscribe int    `json:"newsletter_subscribe"` //
	InternalShortNote   string `json:"internal_short_note"`  //
	InternalLongNote    string `json:"internal_long_note"`   //
}
