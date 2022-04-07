package models

import (
	"gorm.io/gorm"
	"time"
	"gorm.io/plugin/soft_delete"
)

type ProductLibraryPropertyModel struct {
	ID           	string     `gorm:"column:id;primary_key" json:"id"`
	Name         	NullString     `gorm:"column:name;default:null" json:"name"`
	ProductLibraryCategoryId string `gorm:"column:product_library_category_id" json:"product_library_category_id"`
	ProductLibraryTitleId string `gorm:"column:product_library_title_id" json:"product_library_title_id"`
	CreatedBy    NullString  	`gorm:"column:created_by" json:"created_by"`
	CreatedAt    time.Time 		`gorm:"column:created_at" json:"crated_at"`
	UpdatedBy    NullString 	`gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy	 NullString `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductLibraryPropertyModelPreload struct {
	ID           	string     `gorm:"column:id;primary_key" json:"id"`
	Name         	NullString     `gorm:"column:name;default:null" json:"name"`
	ProductLibraryCategoryId string `gorm:"column:product_library_category_id" json:"product_library_category_id"`
	ProductLibraryTitleId string `gorm:"column:product_library_title_id" json:"product_library_title_id"`
	ProductLibraryCategory ProductLibraryCategoryModelPreload `gorm:"foreignKey:ID;references:ProductLibraryCategoryId" json:"product_library_category_data"`
	ProductLibraryTitle ProductLibraryTitleModelPreload `gorm:"foreignKey:ID;references:ProductLibraryTitleId" json:"product_library_title_data"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at" json:"-"`
}



// TableName sets the insert table name for this struct type
func (p *ProductLibraryPropertyModel) TableName() string {
	return "product_library_property"
}
func (p *ProductLibraryPropertyModelPreload) TableName() string {
	return "product_library_property"
}

func (p *ProductLibraryPropertyModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductLibraryPropertyModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductLibraryPropertyModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductLibraryPropertyModel) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println("Before Delete")
	_,con,_ := DbConnect()
	var model ProductLibraryPropertyModel
	con.Model(&model).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return
}

//strcture input
type ProductLibraryPropertyModelInput struct {
	Name         	string     	`json:"name" binding:"required"`
	ProductLibraryCategoryId string `json:"product_library_category_id" binding:"required"`
	ProductLibraryTitleId string `json:"product_library_title_id" binding:"required"`
}
type ProductLibraryPropertyModelUpdate struct {
	ID	string	`json:"id" binding:"required"`
	Name         	string     	`json:"name" binding:"required"`
	ProductLibraryCategoryId string `json:"product_library_category_id" binding:"required"`
	ProductLibraryTitleId string `json:"product_library_title_id" binding:"required"`
}
type ProductLibraryPropertyModelInfo struct {
	ID	string	`json:"id" binding:"required"`
}