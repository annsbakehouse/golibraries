package models

import (
	"time"

	"gorm.io/gorm"
	//"fmt"
	// "database/sql"
)

type EmployeeModel struct {
	ID               string     `gorm:"column:id;primaryKey" json:"id"`                      //
	EmployeeIdCode   int        `gorm:"column:employee_id_code" json:"employee_id_code"`     //
	Company          NullString `gorm:"column:company" json:"company"`                       //
	Fullname         NullString `gorm:"column:fullname" json:"fullname"`                     //
	Initial          NullString `gorm:"column:initial" json:"initial"`                       //
	Active           int        `gorm:"column:active" json:"active"`                         //
	CountryIdPhone1  NullString `gorm:"column:country_id_phone_1" json:"country_id_phone_1"` //
	Phone1           NullString `gorm:"column:phone_1" json:"phone_1"`                       //
	CountryIdPhone2  NullString `gorm:"column:country_id_phone_2" json:"country_id_phone_2"` //
	Phone2           NullString `gorm:"column:phone_2" json:"phone_2"`                       //
	PlaceOfBirth     NullString `gorm:"column:place_of_birth" json:"place_of_birth"`         //
	DateOfBirth      NullDate   `gorm:"column:date_of_birth" json:"date_of_birth"`           //
	Email            NullString `gorm:"column:email" json:"email"`                           //
	Gender           int        `gorm:"column:gender" json:"gender"`                         //
	MaritalStatus    int        `gorm:"column:marital_status" json:"marital_status"`         //
	BloodTypeId      NullString `gorm:"column:blood_type_id" json:"blood_type_id"`           //
	ReligionId       NullString `gorm:"column:religion_id" json:"religion_id"`               //
	Avatar           NullString `gorm:"column:avatar" json:"avatar"`                         //
	JoinDate         NullDate   `gorm:"column:join_date" json:"join_date"`                   //
	TerminationDate  NullDate   `gorm:"column:termination_date" json:"termination_date"`     //
	WarehouseId      NullString `gorm:"column:warehouse_id" json:"warehouse_id"`             //
	DepartmentId     NullString `gorm:"column:department_id" json:"department_id"`           //
	JobLevelId       NullString `gorm:"column:job_level_id" json:"job_level_id"`             //
	UserId           NullString `gorm:"column:user_id" json:"user_id"`                       //
	EmploymentStatus int        `gorm:"column:employment_status" json:"employment_status"`   //
	CreatedBy        NullString `gorm:"column:created_by" json:"created_by"`                 //
	CreatedAt        *time.Time `gorm:"column:created_at" json:"created_at"`                 //
	UpdatedBy        NullString `gorm:"column:updated_by" json:"updated_by"`                 //
	UpdatedAt        *time.Time `gorm:"column:updated_at" json:"updated_at"`                 //
}

type EmployeeModelPreload struct {
	ID                  string                 `gorm:"column:id;primaryKey" json:"id"`                      //
	EmployeeIdCode      int                    `gorm:"column:employee_id_code" json:"employee_id_code"`     //
	Company             NullString             `gorm:"column:company" json:"company"`                       //
	Fullname            NullString             `gorm:"column:fullname" json:"fullname"`                     //
	Initial             NullString             `gorm:"column:initial" json:"initial"`                       //
	Active              int                    `gorm:"column:active" json:"active"`                         //
	CountryIdPhone1     NullString             `gorm:"column:country_id_phone_1" json:"country_id_phone_1"` //
	Phone1              NullString             `gorm:"column:phone_1" json:"phone_1"`                       //
	CountryIdPhone2     NullString             `gorm:"column:country_id_phone_2" json:"country_id_phone_2"` //
	Phone2              NullString             `gorm:"column:phone_2" json:"phone_2"`                       //
	PlaceOfBirth        NullString             `gorm:"column:place_of_birth" json:"place_of_birth"`         //
	DateOfBirth         NullDate               `gorm:"column:date_of_birth" json:"date_of_birth"`           //
	Email               NullString             `gorm:"column:email" json:"email"`                           //
	Gender              int                    `gorm:"column:gender" json:"gender"`                         //
	MaritalStatus       int                    `gorm:"column:marital_status" json:"marital_status"`         //
	BloodTypeId         NullString             `gorm:"column:blood_type_id" json:"blood_type_id"`           //
	ReligionId          NullString             `gorm:"column:religion_id" json:"religion_id"`               //
	Avatar              NullString             `gorm:"column:avatar" json:"avatar"`                         //
	JoinDate            NullDate               `gorm:"column:join_date" json:"join_date"`                   //
	TerminationDate     NullDate               `gorm:"column:termination_date" json:"termination_date"`     //
	WarehouseId         NullString             `gorm:"column:warehouse_id" json:"warehouse_id"`             //
	DepartmentId        NullString             `gorm:"column:department_id" json:"department_id"`           //
	JobLevelId          NullString             `gorm:"column:job_level_id" json:"job_level_id"`             //
	UserId              NullString             `gorm:"column:user_id" json:"user_id"`                       //
	EmploymentStatus    int                    `gorm:"column:employment_status" json:"employment_status"`   //
	UserInfo            UserModel              `gorm:"foreignKey:UserId;references:ID" json:"user_info"`
	CountryIdPhone1Info CountryModelPreload    `gorm:"foreignKey:CountryIdPhone1;references:ID" json:"country_phone1_info"`
	CountryIdPhone2Info CountryModelPreload    `gorm:"foreignKey:CountryIdPhone2;references:ID" json:"country_phone2_info"`
	WarehouseInfo       WarehousePreloadModel  `gorm:"foreignKey:WarehouseId;references:ID" json:"warehouse_info"`
	DepartmentInfo      DepartmentPreloadModel `gorm:"foreignKey:DepartmentId;references:ID" json:"department_info"`
	JobLevelInfo        JobLevelModel          `gorm:"foreignKey:JobLevelId;references:ID" json:"joblevel_info"`
	BloodTypeInfo       BloodTypeModel         `gorm:"foreignKey:BloodTypeId;references:ID" json:"bloodtype_info"`
	ReligionInfo        ReligionModel          `gorm:"foreignKey:ReligionId;references:ID" json:"religion_info"`
}

// TableName sets the insert table name for this struct type
func (u *EmployeeModel) TableName() string {
	return "employee"
}
func (u *EmployeeModelPreload) TableName() string {
	return "employee"
}

func (u *EmployeeModel) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (u *EmployeeModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}
