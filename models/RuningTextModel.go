package models

import (
	"time"

	"gorm.io/gorm"
)

type RunningTextModel struct {
	ID            string     `gorm:"column:id;primary_key" json:"id"`
	Running_text1 NullString `gorm:"column:running_text1" json:"running_text1"`
	Running_text2 NullString `gorm:"column:running_text2;default:null" json:"running_text2"`
	Display_start NullString `gorm:"column:display_start;default:null" json:"display_start"`
	Display_until NullString `gorm:"column:display_until;default:null" json:"display_until"`
	Status        int        `gorm:"column:status;default:null" json:"status"`
	Inc           NullString `gorm:"column:inc" json:"inc"`
	CreatedAt     time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy     NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedBy     NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt     time.Time  `gorm:"column:updated_at" json:"updated_at"`
	// DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt NullString `gorm:"column:deleted_at" json:"deleted_at"`
}

type RunningTextModelPreload struct {
	ID            string       `gorm:"column:id;primary_key" json:"id"`
	Running_text1 NullString   `gorm:"column:running_text1" json:"running_text1"`
	Running_text2 NullString   `gorm:"column:running_text2;default:null" json:"running_text2"`
	Display_start NullDateTime `gorm:"column:display_start;default:null" json:"display_start"`
	Display_until NullDateTime `gorm:"column:display_until;default:null" json:"display_until"`
	Status        int          `gorm:"column:status;default:null" json:"status"`
	Inc           NullString   `gorm:"column:inc" json:"inc"`
	// DeletedBy NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt NullString `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (c *RunningTextModel) TableName() string {
	return "running_text"
}

func (c *RunningTextModelPreload) TableName() string {
	return "running_text"
}

func (c *RunningTextModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *RunningTextModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

type RunningTextModelSave struct {
	Running_text1 string `json:"running_text1" binding:"required"`
	Running_text2 string `json:"running_text2"  binding:"required"`
	Display_start string `json:"display_start" `
	Display_until string `json:"display_until" `
	// Status        int    `json:"status"`
}

type RunningTextModelUpdate struct {
	ID            string `json:"id" binding:"required"`
	Running_text1 string `json:"running_text1" binding:"required"`
	Running_text2 string `json:"running_text2" binding:"required"`
	Display_start string `json:"display_start"`
	Display_until string `json:"display_until"`
	// Status        int    `json:"status"`
}

type RunningTextModelDelete struct {
	ID string `json:"id" binding:"required"`
}

// func (c *CityModel) BeforeDelete(tx *gorm.DB) (err error) {
// 	tt := time.Now()
// 	tFormat := tt.Format("2006-01-02 15:04:05")
// 	c.DeletedAt = NullStringInput(tFormat)
// 	return
// }
