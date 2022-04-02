package models

import (
	"gorm.io/gorm"
	"time"
	"gorm.io/plugin/soft_delete"
)

type ProductAccordionModel struct {
	ID           	string     		`gorm:"column:id;primary_key" json:"id"`
	Name         	NullString     	`gorm:"column:name;default:null" json:"name"`
	OrderSequence	NullInt64		`gorm:"column:order_sequence" json:"order_sequence"`
	Active	 		int				`gorm:"column:active" json:"active"`
	Inc	 			int				`gorm:"column:inc" json:"inc"`
	CreatedBy    	NullString  	`gorm:"column:created_by" json:"created_by"`
	CreatedAt    	time.Time 		`gorm:"column:created_at" json:"crated_at"`
	UpdatedBy    	NullString 		`gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    	time.Time 		`gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy	 	NullString 		`gorm:"column:updated_by" json:"deleted_by"`
	DeletedAt 		soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type ProductAccordionModelPreload struct {
	ID           	string     		`gorm:"column:id;primary_key" json:"id"`
	Name         	NullString     	`gorm:"column:name;default:null" json:"name"`
	OrderSequence	NullInt64		`gorm:"column:order_sequence" json:"order_sequence"`
	Active	 		int				`gorm:"column:active" json:"active"`
	Inc	 			int				`gorm:"column:inc" json:"inc"`
	Language	  	[]LanguageTableModelPreload 	`gorm:"foreignKey:TableTargetID;references:ID" json:"language"`
}



// TableName sets the insert table name for this struct type
func (p *ProductAccordionModel) TableName() string {
	return "product_accordion"
}
func (p *ProductAccordionModelPreload) TableName() string {
	return "product_accordion"
}

func (p *ProductAccordionModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProductAccordionModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductAccordionModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *ProductAccordionModel) AfterDelete(tx *gorm.DB) (err error) {
	// fmt.Println("Before Delete")
	_,con,_ := DbConnect()
	var model ProductAccordionModel
	con.Model(&model).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return
}

//strcture input
type ProductAccordionModelInput struct {
	Name         	string     `json:"name" binding:"required"`
	OrderSequence	int64		`json:"order_sequence"`
	Active	 		int			`json:"active"`
	CountryID	 	[]string	`json:"country_id"`
	Title	 		[]string	`json:"title"`
	Description	 	[]string	`json:"description"`
}
type ProductAccordionModelUpdate struct {
	ID				string		`json:"id" binding:"required"`
	Name         	string     `json:"name" binding:"required"`
	OrderSequence	int64		`json:"order_sequence"`
	Active	 		int			`json:"active"`
	CountryID	 	[]string	`json:"country_id"`
	Title	 		[]string	`json:"title"`
	Description	 	[]string	`json:"description"`
}
type ProductAccordionModelInfo struct {
	ID	string	`json:"id" binding:"required"`
}