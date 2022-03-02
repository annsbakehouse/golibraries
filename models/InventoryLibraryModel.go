package models

import (
	"gorm.io/gorm"
	"time"
)

type InventoryLibraryModel struct {
	ID           string     `gorm:"column:id;primary_key" json:"id"`
	Name         string     `gorm:"column:name;default:null" json:"name"`
	IsParent	 int		`gorm:"column:is_parent" json:"is_parent"`
	ParentID	 NullString		`gorm:"column:parent_id" json:"parent_id"`
	Active	 	 int		 	`gorm:"column:active" json:"active"`
	CreatedBy    NullString  	`gorm:"column:created_by" json:"created_by"`
	CreatedAt    time.Time 		`gorm:"column:created_at" json:"_"`
	UpdatedBy    NullString 	`gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at" sql:"type:timestamp without time zone" json:"updated_at"`
	DeletedBy	 NullString `gorm:"column:updated_by" json:"deleted_by"`
	DeletedAt    time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Child  []InventoryLibraryModel `gorm:"foreignKey:ParentID;references:ID" json:"child"`
}


// TableName sets the insert table name for this struct type
func (p *InventoryLibraryModel) TableName() string {
	return "inventory_library"
}

func (p *InventoryLibraryModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *InventoryLibraryModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryLibraryModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (p *InventoryLibraryModel	) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println("Before Delete")
	_,con,_ := DbConnect()
	var model InventoryLibraryModel
	con.Model(&model).Where("id=?", p.ID).Update("deleted_by",ActiveUser)
	return
}

//strcture input
type InventoryLibraryInput struct {
	Name	string	`json:"name" binding:"required"`
	IsParent	 int	`json:"is_parent"`
	ParentID	 string	 `json:"parent_id"`
	Active	 	 int	 `json:"active"`
}
type InventoryLibraryUpdate struct {
	ID	string	`json:"id" binding:"required"`
	Name	string	`json:"name" binding:"required"`
	IsParent	 int	`json:"is_parent"`
	ParentID	 string	 `json:"parent_id"`
	Active	 	 int	 `json:"active"`
}
type InventoryLibraryInfo struct {
	ID	string	`json:"id" binding:"required"`
}
