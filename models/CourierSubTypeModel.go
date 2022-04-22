package models

import (
	"time"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type CourierSubTypeModel struct {
	ID         			string     `gorm:"column:id;primary_key" json:"id"`
	Name       	string     `gorm:"column:name;size:100" json:"name"`
	CourierTypeID   		string     `gorm:"column:courier_type_id" json:"courier_type_id"`
	INC     			int     `gorm:"column:inc;default:0" json:"inc"`
	Active     			uint64     `gorm:"column:active;" json:"active"`
	CreatedBy  			NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt  			time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy  			NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt  			time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedBy  			NullString `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt  			soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type CourierSubTypeModelPreload struct {
	ID         			string     `gorm:"column:id;primary_key" json:"id"`
	Name       	string     `gorm:"column:name" json:"name"`
	CourierTypeID   		string     `gorm:"column:courier_type_id" json:"courier_type_id"`
	INC     			int     `gorm:"column:inc;default:0" json:"inc"`
	Active     			uint64     `gorm:"column:active" json:"active"`
	DeletedAt  			soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (c *CourierSubTypeModel) TableName() string {
	return "courier_subtype"
}
func (c *CourierSubTypeModelPreload) TableName() string {
	return "courier_subtype"
}

func (c *CourierSubTypeModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *CourierSubTypeModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *CourierSubTypeModel) BeforeDelete(tx *gorm.DB) (err error) {
	var model CourierSubTypeModel
	tx.Model(&model).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return
}

type CourierSubTypeModelInput struct {
	Name       			string     `json:"name" binding:"required"`
	ShortName   		string     `json:"short_name" binding:"required"`
	CummulativeRange  	int     	`json:"cummulative_range"`
	TimeRange  			int     	`json:"time_range"`
	Active     			uint64     `json:"active"`
	PlatformID			[]string	`json:"platform"`
	CourierSubTypeID		[]string	`json:"courier_sub_type_id"`
	CourierSubType		[]string	`json:"courier_sub_type"`
	CourierSubTypeActive		[]int	`json:"courier_sub_type_active"`
}

type CourierSubTypeModelUpdate struct {
	ID         			string     	`json:"id" binding:"required"`
	Name       			string     	`json:"name" binding:"required,max:100"`
	Active     			uint64     	`json:"active"`
}