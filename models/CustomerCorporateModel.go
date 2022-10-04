package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerCorporateModel struct {
	ID                 string     `gorm:"column:id;primary_key" json:"id"`                   //
	CCID               NullString `gorm:"column:ccid" json:"ccid"`                           //
	AltName            NullString `gorm:"column:alt_name" json:"alt_name"`                   //
	AutoOrderNote      NullString `gorm:"column:auto_order_note" json:"auto_order_note"`     //
	AutoInvoiceNote    NullString `gorm:"column:auto_invoice_note" json:"auto_invoice_note"` //
	Active             int        `gorm:"column:active" json:"active"`                       //
	EmployeeId         NullString `gorm:"column:employee_id" json:"employee_id"`             //
	CoaId              NullString `gorm:"column:coa_id" json:"coa_id"`                       //
	MaxAr              float32    `gorm:"column:max_ar" json:"max_ar"`                       //
	BillingAddress     NullString `gorm:"column:billing_address" json:"billing_address"`     //
	CustomerRetailId   string     `gorm:"column:customer_retail_id" json:"customer_retail_id"`
	OverdueLimit       int        `gorm:"column:overdue_limit" json:"overdue_limit"`
	AccountExecutiveId NullString `gorm:"column:account_executive_id" json:"account_executive_id"`
	PaymentTypeId      NullString `gorm:"column:payment_type_id" json:"payment_type_id"`
	CreatedBy          NullString `gorm:"column:created_by" json:"created_by"` //
	CreatedAt          time.Time  `gorm:"column:created_at" json:"created_at"` //
	UpdatedBy          NullString `gorm:"column:updated_by" json:"updated_by"` //
	UpdatedAt          time.Time  `gorm:"column:updated_at" json:"updated_at"` //
	DeletedBy          NullString `gorm:"column:deleted_by" json:"deleted_by"` //
	// DeletedAt         soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type CustomerCorporateModelPreload struct {
	ID                   string                       `gorm:"column:id;primary_key" json:"id"`                   //
	CCID                 NullString                   `gorm:"column:ccid" json:"ccid"`                           //
	AltName              NullString                   `gorm:"column:alt_name" json:"alt_name"`                   //
	AutoOrderNote        NullString                   `gorm:"column:auto_order_note" json:"auto_order_note"`     //
	AutoInvoiceNote      NullString                   `gorm:"column:auto_invoice_note" json:"auto_invoice_note"` //
	Active               int                          `gorm:"column:active" json:"active"`                       //
	EmployeeId           NullString                   `gorm:"column:employee_id" json:"employee_id"`             //
	CoaId                NullString                   `gorm:"column:coa_id" json:"coa_id"`                       //
	MaxAr                float32                      `gorm:"column:max_ar" json:"max_ar"`                       //
	BillingAddress       NullString                   `gorm:"column:billing_address" json:"billing_address"`     //
	CustomerRetailId     string                       `gorm:"column:customer_retail_id" json:"customer_retail_id"`
	OverdueLimit         int                          `gorm:"column:overdue_limit" json:"overdue_limit"`
	CustomerRetailInfo   CustomerRetailModelPreload   `gorm:"foreignKey:CustomerRetailId;references:ID" json:"customer_retail_info"`
	AccountExecutiveId   NullString                   `gorm:"column:account_executive_id" json:"account_executive_id"`
	PaymentTypeId        NullString                   `gorm:"column:payment_type_id" json:"payment_type_id"`
	ContactPersonInfo    []ContactPersonPreload       `gorm:"foreignKey:CustomerCorporateId;references:ID" json:"contact_persons"`
	PaymentTypeInfo      PaymentTypeModelPreload      `gorm:"foreignKey:PaymentTypeId;references:ID" json:"payment_type_info"`
	AccountExecutiveInfo AccountExecutiveModelPreload `gorm:"foreignKey:AccountExecutiveId;references:ID" json:"account_executive_info"`
}

func (c *CustomerCorporateModel) TableName() string {
	return "customer_corporate"
}

func (c *CustomerCorporateModelPreload) TableName() string {
	return "customer_corporate"
}

func (c *CustomerCorporateModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *CustomerCorporateModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (c *CustomerCorporateModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

type CustomerCorporateInfoForm struct {
	ID string `json:"id" binding:"required"`
}

type CustomerCorporateSaveForm struct {
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
	// Gender            int    `json:"gender"`                                  //
	Email              string  `json:"email" binding:"required"` //
	InternalShortNote  string  `json:"internal_short_note"`      //
	InternalLongNote   string  `json:"internal_long_note"`
	CustomersLevelID   string  `json:"customers_level_id" binding:"required"` //
	AutoOrderNote      string  `json:"auto_order_note"`                       //
	AutoInvoiceNote    string  `json:"auto_invoice_note"`                     //
	Active             int     `json:"active"`                                //
	EmployeeId         string  `json:"employee_id"`                           //
	CoaId              string  `json:"coa_id"`                                //
	MaxAr              float32 `json:"max_ar"`                                //
	BillingAddress     string  `json:"biling_address"`                        //
	OverdueLimit       int     `json:"overdue_limit"`                         //
	AccountExecutiveId string  `json:"account_executive_id"`                  //
	PaymentTypeId      string  `json:"payment_type_id"`                       //
}

type CustomerCorporateFormUpdate struct {
	ID                 string  `json:"id" binding:"required"`
	CustomerRetailID   string  `json:"customer_retail_id" binding:"required"`
	Name               string  `json:"name" binding:"required"`                 //
	AltName            string  `json:"alt_name"`                                //
	PhoneCountryID1    string  `json:"phone_1_country_id" binding:"required"`   //
	PhoneCountryCode1  string  `json:"phone_1_country_code" binding:"required"` //
	Phone1             string  `json:"phone_1" binding:"required"`              //
	PhoneCountryID2    string  `json:"phone_2_country_id"`                      //
	PhoneCountryCode2  string  `json:"phone_2_country_code"`                    //
	Phone2             string  `json:"phone_2"`                                 //
	PhoneNote2         string  `json:"phone_2_note"`                            //
	PhoneCountryID3    string  `json:"phone_3_country_id"`                      //
	PhoneCountryCode3  string  `json:"phone_3_country_code"`                    //
	Phone3             string  `json:"phone_3"`                                 //
	PhoneNote3         string  `json:"phone_3_note"`                            //
	Birthday           string  `json:"birthday"`                                //
	Gender             int     `json:"gender"`                                  //
	Email              string  `json:"email" binding:"required"`                //
	InternalShortNote  string  `json:"internal_short_note"`                     //
	InternalLongNote   string  `json:"internal_long_note"`
	CustomersLevelID   string  `json:"customers_level_id" binding:"required"` //
	AutoOrderNote      string  `json:"auto_order_note"`                       //
	Active             int     `json:"active"`                                //
	EmployeeId         string  `json:"employee_id"`                           //
	CoaId              string  `json:"coa_id"`                                //
	MaxAr              float32 `json:"max_ar"`                                //
	BillingAddress     string  `json:"biling_address"`                        //
	OverdueLimit       int     `json:"overdue_limit"`                         //
	AccountExecutiveId string  `json:"account_executive_id"`                  //
	PaymentTypeId      string  `json:"payment_type_id"`                       //
}
