package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductFilterTagProductModel struct {
	ID                    string     `gorm:"column:id;primary_key" json:"id"`
	MktProductFilterTagID string     `gorm:"column:mkt_product_filter_tag_id" json:"mkt_product_filter_tag_id"`
	ProductID             string     `gorm:"column:product_id" json:"product_id"`
	CreatedBy             NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt             time.Time  `gorm:"column:created_at" json:"crated_at"`
	UpdatedBy             NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt             time.Time  `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	// DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductFilterTagProductModelPreload struct {
	ID                    string       `gorm:"column:id;primary_key" json:"id"`
	MktProductFilterTagID string       `gorm:"column:mkt_product_filter_tag_id" json:"mkt_product_filter_tag_id"`
	ProductID             string       `gorm:"column:product_id" json:"product_id"`
	ProductInfo           ProductModel `gorm:"foreignKey:ID;references:ProductID" json:"product_info"`
	// DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (p *ProductFilterTagProductModel) TableName() string {
	return "mkt_product_filter_tag_product"
}
func (p *ProductFilterTagProductModelPreload) TableName() string {
	return "mkt_product_filter_tag_product"
}

func (p *ProductFilterTagProductModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductFilterTagProductModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductFilterTagProductModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductFilterTagProductModel) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println(p.ID)
	// var model ProductContentParameterTitleModel
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return
}

//strcture input
type ProductFilterTagProductModelInput struct {
	ID                    string `json:"id"`
	MktProductFilterTagID string `json:"mkt_product_filter_tag_id" binding:"required"`
	ProductID             string `json:"product_id" binding:"required"`
}
type ProductFilterTagProductModelInfo struct {
	ID string `json:"id" binding:"required"`
}
