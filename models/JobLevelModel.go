package models

import (
	"gorm.io/gorm"
	"time"
)

type JobLevelModel struct {
	ID        		string     			`gorm:"column:id;primary_key" json:"id"`
	Name			NullString			`gorm:"column:name" json:"name"`
	DepartmentId    string     			`gorm:"column:department_id" json:"department_id"`	
	Active			int					`gorm:"column:active" json:"active"`	
	CreatedBy		NullString			`gorm:"column:created_by" json:"created_by"`
	CreatedAt		time.Time			`gorm:"column:created_at" json:"created_at"`
	UpdatedBy		NullString			`gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt		time.Time			`gorm:"column:updated_at" json:"updated_at"`
	Departement 	DepartmentModel 	`gorm:"foreignKey:ID;references:DepartmentId" json:"departement"`
}

// TableName sets the insert table name for this struct type
func (d *JobLevelModel) TableName() string {
	return "job_level"
}

func (d *JobLevelModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (d *JobLevelModel) BeforeUpdate(tx *gorm.DB) (err error) {
	
	return
}
func (d *JobLevelModel) AfterUpdate(tx *gorm.DB) (err error) {
	
	return
}

type JobLevelInfoForm struct {
	ID string `json:"id" binding:"required"`
}

type JobLevelSaveForm struct {
	Name			string 	`json:"name" binding:"required"`
	DepartmentId    string 	`json:"department_id" binding:"required"`
	Active		    int 	`json:"active"`
}

type JobLevelUpdateForm struct {
	ID				string 	`json:"id" binding:"required"`
	Name			string 	`json:"name" binding:"required"`
	DepartmentId    string 	`json:"department_id"`
	Active		    int 	`json:"active"`
}