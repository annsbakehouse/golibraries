package models

import (
	"gorm.io/gorm"
	"time"
	"gorm.io/plugin/soft_delete"
	"fmt"
)

type ProductContentParameterSelectionModel struct {
	ID           			string     		`gorm:"column:id;primary_key" json:"id"`
	Name         			NullString     	`gorm:"column:name;default:null" json:"name"`
	ProductContentTitleID	string			`gorm:"column:product_content_title_id" json:"product_content_title_id"`
	Inc	 			int				`gorm:"column:inc" json:"inc"`
	CreatedBy    	NullString  	`gorm:"column:created_by" json:"created_by"`
	CreatedAt    	time.Time 		`gorm:"column:created_at" json:"crated_at"`
	UpdatedBy    	NullString 		`gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    	time.Time 		`gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy	 	NullString 		`gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductContentParameterSelectionModelPreload struct {
	ID           					string     		`gorm:"column:id;primary_key" json:"id"`
	Name         					NullString     	`gorm:"column:name;default:null" json:"name"`
	ProductContentTitleID			string			`gorm:"column:product_content_title_id" json:"product_content_title_id"`
	Inc	 							int				`gorm:"column:inc" json:"inc"`
	DeletedAt 						soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
	ProductContentParameterTitle	ProductContentParameterTitleModelPreload 	`gorm:"foreignKey:ID;references:ProductContentTitleID" json:"product_content_title"`
}

// TableName sets the insert table name for this struct type
func (p *ProductContentParameterSelectionModel) TableName() string {
	return "product_content_selection"
}
func (p *ProductContentParameterSelectionModelPreload) TableName() string {
	return "product_content_selection"
}

func (p *ProductContentParameterSelectionModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductContentParameterSelectionModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductContentParameterSelectionModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductContentParameterSelectionModel) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println(p.ID)
	var model ProductContentParameterSelectionModel
	tx.Model(&model).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return
}

//strcture input
type ProductContentParameterSelectionModelInput struct {
	Name         	string     `json:"name" binding:"required"`
	ProductContentTitleID	string		`json:"product_content_title_id" binding:"required"`
}
type ProductContentParameterSelectionModelUpdate struct {
	ID				string		`json:"id" binding:"required"`
	Name         	string     `json:"name" binding:"required"`
	ProductContentTitleID	string		`json:"product_content_title_id" binding:"required"`
}
type ProductContentParameterSelectionModelInfo struct {
	ID	string	`json:"id" binding:"required"`
}