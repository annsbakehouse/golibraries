package models

import (
	"gorm.io/gorm"
	"time"
	"gorm.io/plugin/soft_delete"
)

type ProductModelImageModel struct {
	ID           	string     		`gorm:"column:id;primary_key" json:"id"`
	ProductModelDataID       string     	`gorm:"column:product_model_data_id;default:null" json:"product_model_data_id"`
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

type ProductModelImageModelPreload struct {
	ID           	string     		`gorm:"column:id;primary_key" json:"id"`
	ProductModelDataID       string     	`gorm:"column:product_model_data_id;default:null" json:"product_model_data_id"`
	ProductImageUsageID			string			`gorm:"column:product_image_usage_id" json:"product_image_usage_id"`
	Image			string			`gorm:"column:image" json:"image"`
	Inc				int			`gorm:"column:inc" json:"inc"`
	DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
	DescriptionLang	  	[]LanguageTableModelPreload 	`gorm:"foreignKey:TableTargetID;references:ID" json:"image_description"`
	AltLang	  	[]LanguageTableModelPreload 	`gorm:"foreignKey:TableTargetID;references:ID" json:"image_alt"`
	ImageUsage ProductImageUsageModelPreload `gorm:"foreignKey:ID;references:ProductImageUsageID" json:"image_usage"`
}

// TableName sets the insert table name for this struct type
func (p *ProductModelImageModel) TableName() string {
	return "product_model_data_image"
}
func (p *ProductModelImageModelPreload) TableName() string {
	return "product_model_data_image"
}

func (p *ProductModelImageModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductModelImageModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductModelImageModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductModelImageModel) BeforeDelete(tx *gorm.DB) (err error) {
	var model ProductModelImageModel
	tx.Model(&model).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return 
}
func (p *ProductModelImageModel) AfterDelete(tx *gorm.DB) (err error) {
	
	return
}

//strcture input
type ProductModelImageModelInput struct {
	ProductModelDataID       			string    	`json:"product_model_data_id" binding:"required"`
	ProductImageUsageID			[]string		`json:"product_image_usage_id" binding:"required"`
	Image						[]string		`json:"image" binding:"required"`
	Country						[][]string		`json:"country" binding:"required"`
	ImageDescription			[][]string 		`json:"image_description"`
	ImageAlt					[][]string 		`json:"image_alt"`
}
type ProductModelImageModelUpdate struct {
	ID							[]string		`json:"id"`
	ProductModelDataID       			string    	`json:"product_model_data_id" binding:"required"`
	ProductImageUsageID			[]string		`json:"product_image_usage_id" binding:"required"`
	Image						[]string		`json:"image"`
	Country						[][]string		`json:"country" binding:"required"`
	ImageDescription			[][]string 		`json:"image_description"`
	ImageAlt					[][]string 		`json:"image_alt"`
}
type ProductModelImageModelInfo struct {
	ID	string	`json:"id" binding:"required"`
}