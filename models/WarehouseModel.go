package golibraries/models

import (
	"gorm.io/gorm"
	"time"
)

type WarehouseModel struct {
	ID           	string     		`gorm:"column:id;primary_key" json:"id"`
	LetterCode		NullString		`gorm:"column:letter_code;unique" json:"letter_code"`
	Name         	NullString     	`gorm:"column:name;default:null" json:"name"`
	Active       	uint64     		`gorm:"column:active;default:1" json:"active"`
	Address 		NullString     	`gorm:"column:address;" json:"address"`
	Longitude    	NullString     	`gorm:"column:longitude;default:null" json:"longitude,omitempty"`
	Latitude    	NullString     	`gorm:"column:latitude;default:null" json:"latitude,omitempty"`
	W1    			uint64     		`gorm:"column:w1;default:1" json:"w1"`
	W2   			uint64     		`gorm:"column:w2;default:1" json:"w2"`
	W3    			uint64     		`gorm:"column:w3;default:1" json:"w3"`
	W4    			uint64     		`gorm:"column:w4;default:1" json:"w4"`
	W5    			uint64     		`gorm:"column:w5;default:1" json:"w5"`
	Branch    		uint64     		`gorm:"column:branch;default:1" json:"branch"`
	CreatedBy    	NullString 		`gorm:"column:created_by" json:"created_by"`
	CreatedAt    	time.Time 		`gorm:"column:created_at" json:"created_at"`
	UpdatedBy    	NullString 		`gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    	time.Time 		`gorm:"column:updated_at" json:"updated_at"`
	Inc				uint64			`gorm:"column:inc;<-:false" json:"inc"`
	WhDepartmentWarehouse []DepartmentWarehouseModel `gorm:"foreignKey:WarehouseId;references:ID" json:"wh_departement_warehouse"`
}

type WarehousePreloadModel struct {
	ID           	string     		`gorm:"column:id;primary_key" json:"id"`
	LetterCode		string			`gorm:"column:letter_code;unique" json:"letter_code"`
	Name         	string     		`gorm:"column:name;default:null" json:"name"`
	Active       	uint64     		`gorm:"column:active;default:1" json:"active"`
	Address 		string     		`gorm:"column:address;" json:"address"`
	Longitude    	string     		`gorm:"column:longitude;default:null" json:"longitude,omitempty"`
	Latitude    	string     		`gorm:"column:latitude;default:null" json:"latitude,omitempty"`
	W1    			string     		`gorm:"column:w1;default:1" json:"w1"`
	W2   			string     		`gorm:"column:w2;default:1" json:"w2"`
	W3    			string     		`gorm:"column:w3;default:1" json:"w3"`
	W4    			string     		`gorm:"column:w4;default:1" json:"w4"`
	W5    			string     		`gorm:"column:w5;default:1" json:"w5"`
	Branch    		string     		`gorm:"column:branch;default:1" json:"branch"`
	Inc				uint64			`gorm:"column:inc;<-:false" json:"inc"`
	WhDepartmentWarehouse []DepartmentWarehouseModel `gorm:"foreignKey:WarehouseId;references:ID" json:"wh_departement_warehouse"`
}

// TableName sets the insert table name for this struct type
func (p *WarehousePreloadModel) TableName() string {
	return "warehouse"
}

func (p *WarehouseModel) TableName() string {
	return "warehouse"
}

func (p *WarehouseModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *WarehouseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

type WarehouseSaveData struct {
	LetterCode		string		`json:"letter_code" binding:"required"`
	Name         	string     	`json:"name" binding:"required"`
	Active       	uint64     	`json:"active"`
	Address 		string     	`json:"address"`
	Longitude    	string     	`json:"longitude"`
	Latitude    	string     	`json:"latitude"`
	W1    			uint64     	`json:"w1"`
	W2   			uint64     	`json:"w2"`
	W3    			uint64     	`json:"w3"`
	W4    			uint64     	`json:"w4"`
	W5    			uint64     	`json:"w5"`
	Branch    		uint64     	`json:"branch"`
}

type WarehouseUpdateData struct {
	ID				string		`json:"id" binding:"required"`
	LetterCode		string		`json:"letter_code" binding:"required"`
	Name         	string     	`json:"name" binding:"required"`
	Active       	uint64     	`json:"active"`
	Address 		string     	`json:"address"`
	Longitude    	string     	`json:"longitude"`
	Latitude    	string     	`json:"latitude"`
	W1    			uint64     	`json:"w1"`
	W2   			uint64     	`json:"w2"`
	W3    			uint64     	`json:"w3"`
	W4    			uint64     	`json:"w4"`
	W5    			uint64     	`json:"w5"`
	Branch    		uint64     	`json:"branch"`
	StockOrder    	string     	`json:"stock_order" binding:"required"`
}

type WarehouseInfoData struct {
	ID				string		`json:"id" binding:"required"`
}