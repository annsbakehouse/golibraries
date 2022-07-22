package models

import (
	"time"

	"gorm.io/gorm"
)

type MKTProductAccordionModel struct {
	ID                 string     `gorm:"column:id;primary_key" json:"id"`
	ProductID          string     `gorm:"column:product_id;default:null" json:"product_id"`
	ProductModelID     NullString `gorm:"column:product_model_id;default:null" json:"product_model_id"`
	ProductModelDataID NullString `gorm:"column:product_model_data_id;default:null" json:"product_model_data_id"`
	ProductAccordionID string     `gorm:"column:product_accordion_id;default:null" json:"product_accordion_id"`
	CreatedBy          NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt          time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy          NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" json:"updated_at"`
	// DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type MKTProductAccordionModelPreload struct {
	ID                   string                       `gorm:"column:id;primary_key" json:"id"`
	ProductID            string                       `gorm:"column:product_id;default:null" json:"product_id"`
	ProductModelID       NullString                   `gorm:"column:product_model_id;default:null" json:"product_model_id"`
	ProductModelDataID   NullString                   `gorm:"column:product_model_data_id;default:null" json:"product_model_data_id"`
	ProductAccordionID   string                       `gorm:"column:product_accordion_id;default:null" json:"product_accordion_id"`
	ProductAccordionInfo ProductAccordionModelPreload `gorm:"foreignKey:ProductAccordionID;references:ID" json:"product_accordion_info"`
	// DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (p *MKTProductAccordionModel) TableName() string {
	return "mkt_product_accordion"
}
func (p *MKTProductAccordionModelPreload) TableName() string {
	return "mkt_product_accordion"
}

func (p *MKTProductAccordionModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *MKTProductAccordionModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *MKTProductAccordionModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *MKTProductAccordionModel) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println(p.ID)
	// var model ProductContentParameterTitleModel
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return
}

//strcture input
type MKTProductAccordionModelInput struct {
	ProductID          string     `gorm:"column:product_id;default:null" json:"product_id"`
	ProductModelID     NullString `gorm:"column:product_model_id;default:null" json:"product_model_id"`
	ProductModelDataID NullString `gorm:"column:product_model_data_id;default:null" json:"product_model_data_id"`
	ProductAccordionID string     `gorm:"column:product_accordion_id;default:null" json:"product_accordion_id"`
}
type MKTProductAccordionModelUpdate struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
type MKTProductAccordionModelInfo struct {
	ID string `json:"id" binding:"required"`
}
