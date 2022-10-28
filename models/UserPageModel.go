package models

import (
	"time"
	"gorm.io/gorm"
	//"database/sql"
	//"annsbakehouse.com/erp/library"
	//"fmt"
)

type UserPageModel struct {
	ID         string         `gorm:"column:id;primary_key" json:"id"`         //
	PageName        NullString 		`gorm:"column:UID" json:"UID,string"`                   //
	PageUrl   NullString `gorm:"column:fullname" json:"fullname"`         //
	CreatedBy  NullString `gorm:"column:created_by" json:"created_by"`     //
	CreatedAt  time.Time   `gorm:"column:created_at" json:"created_on"`     //
	UpdatedBy  NullString `gorm:"column:updated_by" json:"updated_by"`     //
	UpdatedAt  time.Time   `gorm:"column:updated_at" json:"updated_on"`     //
	IsAdmin    int  `gorm:"column:is_admin" json:"is_admin"`         //
}

// TableName sets the insert table name for this struct type
func (u *UserPageModel) TableName() string {
	return "user_page"
}

func (u *UserPageModel) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (u *UserPageModel) BeforeUpdate(tx *gorm.DB) (err error) {
	// tt := time.Now()
	// tFormat := tt.Format("2006-01-02 15:04:05")
	// u.UpdatedOn = tFormat
	return nil
}

func (u *UserPageModel) BeforeDelete(tx *gorm.DB) (err error) {
	// tt := time.Now()
	// tFormat := tt.Format("2006-01-02 15:04:05")
	// u.DeletedOn = tFormat
	return
}