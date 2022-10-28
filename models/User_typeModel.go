package models

import (
	"time"
	"gorm.io/gorm"
	"fmt"
	// "database/sql"
)

type User_typeModel struct {
	ID        string        `gorm:"column:id;primaryKey" json:"id"`     //
	Name      NullString 	`gorm:"column:name" json:"name"`             //
	CreatedBy string 		`gorm:"column:created_by" json:"created_by"` //
	CreatedAt time.Time    `gorm:"column:created_at" json:"created_on"` //
	UpdatedBy NullString 	`gorm:"column:updated_by" json:"updated_by"` //
	UpdatedAt time.Time    `gorm:"column:updated_at" json:"updated_on"` //
}

// TableName sets the insert table name for this struct type
func (u *User_typeModel) TableName() string {
	fmt.Println("a")
	return "user_type"
}

func (u *User_typeModel) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (u *User_typeModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}