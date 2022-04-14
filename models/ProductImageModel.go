package models

import (
	"gorm.io/gorm"
	"time"
	"gorm.io/plugin/soft_delete"
)

type ProductImageModel struct {
	ID           	string     		`gorm:"column:id;primary_key" json:"id"`
	ProductID       string     	`gorm:"column:product_id;default:null" json:"product_id"`
	ProductImageUsageID			string			`gorm:"column:product_image_usage_id" json:"product_image_usage_id"`
	Image			string			`gorm:"column:image" json:"image"`
	Inc				int			`gorm:"column:inc" json:"inc"`
	CreatedBy    	NullString  	`gorm:"column:created_by" json:"created_by"`
	CreatedAt    	time.Time 		`gorm:"column:created_at" json:"crated_at"`
	UpdatedBy    	NullString 		`gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    	time.Time 		`gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy	 	NullString 		`gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductImageModelPreload struct {
	ID           	string     		`gorm:"column:id;primary_key" json:"id"`
	ProductID       string     	`gorm:"column:product_id;default:null" json:"product_id"`
	ProductImageUsageID			string			`gorm:"column:product_image_usage_id" json:"product_image_usage_id"`
	Image			string			`gorm:"column:image" json:"image"`
	Inc				int			`gorm:"column:inc" json:"inc"`
	DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
	DescriptionLang	  	[]LanguageTableModelPreload 	`gorm:"foreignKey:TableTargetID;references:ID" json:"image_description"`
	AltLang	  	[]LanguageTableModelPreload 	`gorm:"foreignKey:TableTargetID;references:ID" json:"image_alt"`
	ImageUsage ProductImageUsageModelPreload `gorm:"foreignKey:ID;references:ProductImageUsageID" json:"image_usage"`
}

// TableName sets the insert table name for this struct type
func (p *ProductImageModel) TableName() string {
	return "product_image"
}
func (p *ProductImageModelPreload) TableName() string {
	return "product_image"
}

func (p *ProductImageModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductImageModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductImageModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductImageModel) BeforeDelete(tx *gorm.DB) (err error) {
	var model ProductImageModel
	tx.Model(&model).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return 
}
func (p *ProductImageModel) AfterDelete(tx *gorm.DB) (err error) {
	
	return
}

//strcture input
type ProductImageModelInput struct {
	ProductID       			string    	`json:"product_id" binding:"required"`
	ProductImageUsageID			[]string		`json:"product_image_usage_id" binding:"required"`
	Image						[]string		`json:"image" binding:"required"`
	Country						[][]string		`json:"country" binding:"required"`
	ImageDescription			[][]string 		`json:"image_description"`
	ImageAlt					[][]string 		`json:"image_alt"`
}
type ProductImageModelUpdate struct {
	ID							[]string		`json:"id"`
	ProductID       			string    		`json:"product_id" binding:"required"`
	ProductImageUsageID			[]string		`json:"product_image_usage_id" binding:"required"`
	Image						[]string		`json:"image"`
	Country						[][]string		`json:"country" binding:"required"`
	ImageDescription			[][]string 		`json:"image_description"`
	ImageAlt					[][]string 		`json:"image_alt"`
}
type ProductImageModelInfo struct {
	ID	string	`json:"id" binding:"required"`
}