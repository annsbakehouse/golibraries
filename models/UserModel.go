package models

import (
	"time"

	"gorm.io/gorm"

	//"database/sql"
	"fmt"

	"github.com/annsbakehouse/golibraries/library"
	"github.com/google/uuid"
)

type UserModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`             //
	UID       NullString `gorm:"column:UID" json:"UID,string"`                //
	Fullname  NullString `gorm:"column:fullname" json:"fullname"`             //
	Username  NullString `gorm:"column:username;uniqueIndex" json:"username"` //
	Password  NullString `gorm:"column:password" json:"password"`             //
	Email     NullString `gorm:"column:email" json:"email"`                   //
	Active    int        `gorm:"column:active" json:"active"`                 //
	Qr        NullString `gorm:"column:qr" json:"qr"`                         //
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`         //
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_on"`         //
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`         //
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_on"`         //
	IsAdmin   int        `gorm:"column:is_admin" json:"is_admin"`             //
}

// TableName sets the insert table name for this struct type
func (u *UserModel) TableName() string {
	return "user"
}
func UserModelQR(data string, id string) string {
	if len(data) == 0 {
		data = uuid.New().String()
	}
	var model []UserModel
	db, dbReader, _ := DbConnect()
	defer db.Close()
	res := dbReader.Model(&model).Where("qr=?", data)
	if len(id) > 0 {
		res.Where("id!=?", id)
	}
	res.Find(&model)
	if res.RowsAffected == 0 {
		return data
	}
	uid := uuid.New().String()
	return UserModelQR(uid, id)
}
func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	_, dbReader, _ := DbConnect()
	var sequence int
	dbReader.Raw("SELECT last_value FROM user_inc_seq").Scan(&sequence)
	genericUID := library.GenUID(sequence+1, 2, 2)
	u.UID = NullStringInput(fmt.Sprintf("%v", genericUID))
	u.Qr = NullStringInput(UserModelQR(u.Qr.String, ""))
	return nil
}

func (u *UserModel) BeforeUpdate(tx *gorm.DB) (err error) {
	// tt := time.Now()
	// tFormat := tt.Format("2006-01-02 15:04:05")
	// u.UpdatedOn = tFormat
	return nil
}

func (u *UserModel) BeforeDelete(tx *gorm.DB) (err error) {
	// tt := time.Now()
	// tFormat := tt.Format("2006-01-02 15:04:05")
	// u.DeletedOn = tFormat
	return
}
