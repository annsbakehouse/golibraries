package models

import (
	"time"
	"gorm.io/gorm"
	//"fmt"
	// "database/sql"
)

type EmployeeModel struct {
	ID        				string		`gorm:"column:id;primaryKey" json:"id"`     //
	EmployeeIdCode			int			`gorm:"column:employee_id_code" json:"employee_id_code"` //
	Company					NullString	`gorm:"column:company" json:"company"` //
	Fullname				NullString	`gorm:"column:fullname" json:"fullname"` //
	Initial					NullString	`gorm:"column:initial" json:"initial"` //
	Active					int			`gorm:"column:active" json:"active"` //
	CountryIdPhone1			NullString	`gorm:"column:country_id_phone_1" json:"country_id_phone_1"` //
	Phone1					NullString	`gorm:"column:phone_1" json:"phone_1"` //
	CountryIdPhone2			NullString	`gorm:"column:country_id_phone_2" json:"country_id_phone_2"` //
	Phone2					NullString	`gorm:"column:phone_2" json:"phone_2"` //
	PlaceOfBirth			NullString	`gorm:"column:place_of_birth" json:"place_of_birth"` //
	DateOfBirth				time.Time	`gorm:"column:date_of_birth" json:"date_of_birth"` //
	Email      				NullString 	`gorm:"column:email" json:"email"`               //
	Gender					int			`gorm:"column:gender" json:"gender"` //
	MaritalStatus			int			`gorm:"column:marital_status" json:"marital_status"` //
	BloodTypeId				NullString	`gorm:"column:blood_type_id" json:"blood_type_id"` //
	ReligionId				NullString	`gorm:"column:religion_id" json:"religion_id"` //
	Avatar					NullString	`gorm:"column:avatar" json:"avatar"` //
	JoinDate				time.Time	`gorm:"column:join_date" json:"join_date"` //
	TerminationDate			time.Time	`gorm:"column:termination_date" json:"termination_date"` //
	WarehouseId				NullString	`gorm:"column:warehouse_id" json:"warehouse_id"` //
	DepartmentId			NullString	`gorm:"column:department_id" json:"department_id"` //
	JobLevelId				NullString	`gorm:"column:job_level_id" json:"job_level_id"` //
	UserId					NullString	`gorm:"column:user_id" json:"user_id"` //
	EmploymentStatus		int			`gorm:"column:employment_status" json:"employment_status"` //
	CreatedBy				NullString 	`gorm:"column:created_by" json:"created_by"` //
	CreatedAt 				time.Time   `gorm:"column:created_at" json:"created_at"` //
	UpdatedBy 				NullString 	`gorm:"column:updated_by" json:"updated_by"` //
	UpdatedAt 				time.Time   `gorm:"column:updated_at" json:"updated_at"` //
}

// TableName sets the insert table name for this struct type
func (u *EmployeeModel) TableName() string {
	return "employee"
}

func (u *EmployeeModel) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (u *EmployeeModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}