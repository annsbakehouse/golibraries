package models

import (
	"time"

	"gorm.io/gorm"
)

type DeliveryFeeModel struct {
	ID               string     `gorm:"column:id;primary_key" json:"id"`
	SubdistrictId    NullString `gorm:"column:subdistrict_id;default:null" json:"subdistrict_id"`
	DistrictId       NullString `gorm:"column:district_id;default:null" json:"district_id"`
	CityId           string     `gorm:"column:city_id;default:null" json:"city_id"`
	ProvinceId       string     `gorm:"column:province_id;default:null" json:"province_id"`
	CountryId        string     `gorm:"column:country_id;default:null" json:"country_id"`
	Price            int        `gorm:"column:price;default:0" json:"price"`
	WarehouseID      string     `gorm:"column:warehouse_id;default:null" json:"warehouse_id"`
	CourierSubtypeId string     `gorm:"column:courier_subtype_id" json:"courier_subtype_id"`
	CourierTypeId    string     `gorm:"column:courier_type_id" json:"courier_type_id"`
	Active           uint64     `gorm:"column:active" json:"active"`
	CreatedBy        NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt        time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy        NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt        time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedBy        NullString `gorm:"column:deleted_by" json:"deleted_by"`
	// DeletedAt        soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
}

type DeliveryFeeModelPreload struct {
	ID               string     `gorm:"column:id;primary_key" json:"id"`
	SubdistrictId    NullString `gorm:"column:subdistrict_id;default:null" json:"subdistrict_id"`
	DistrictId       NullString `gorm:"column:district_id;default:null" json:"district_id"`
	CityId           string     `gorm:"column:city_id;default:null" json:"city_id"`
	ProvinceId       string     `gorm:"column:province_id;default:null" json:"province_id"`
	CountryId        string     `gorm:"column:country_id;default:null" json:"country_id"`
	Price            int        `gorm:"column:price;default:0" json:"price"`
	WarehouseID      string     `gorm:"column:warehouse_id;default:null" json:"warehouse_id"`
	CourierSubtypeId string     `gorm:"column:courier_subtype_id" json:"courier_subtype_id"`
	CourierTypeId    string     `gorm:"column:courier_type_id" json:"courier_type_id"`
	Active           uint64     `gorm:"column:active" json:"active"`
	//DeletedAt        soft_delete.DeletedAt      `gorm:"uniqueIndex:udx_name;column:deleted_at" json:"deleted_at"`
	CountryInfo     CountryModelPreload        `gorm:"foreignKey:ID;references:CountryId" json:"country_info"`
	ProvinceInfo    ProvinceModelPreload       `gorm:"foreignKey:ID;references:ProvinceId" json:"province_info"`
	CityInfo        CityModelPreload           `gorm:"foreignKey:ID;references:CityId" json:"city_info"`
	DistrictInfo    DistrictModelPreload       `gorm:"foreignKey:ID;references:DistrictId" json:"district_info"`
	SubDistrictInfo SubDistrictModelPreload    `gorm:"foreignKey:ID;references:SubdistrictId" json:"subdistrict_info"`
	CourierInfo     CourierTypeModelPreload    `gorm:"foreignKey:ID;references:CourierTypeId" json:"courier_info"`
	SubCourierInfo  CourierSubTypeModelPreload `gorm:"foreignKey:ID;references:CourierSubtypeId" json:"subcourier_info"`
}

// TableName sets the insert table name for this struct type
func (c *DeliveryFeeModel) TableName() string {
	return "delivery_fee"
}

func (c *DeliveryFeeModelPreload) TableName() string {
	return "delivery_fee"
}

func (c *DeliveryFeeModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (c *DeliveryFeeModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *DeliveryFeeModel) BeforeDelete(tx *gorm.DB) (err error) {
	// var model DeliveryFeeModel
	// fmt.Println(model)
	// tx.Model(&model).Where("id=?", p.ID).Update("deleted_by", ActiveUser)
	return
}

type DeliveryFeeModelInput struct {
	SubdistrictId    string `json:"subdistrict_id"`
	DistrictId       string `json:"district_id"`
	CityId           string `json:"city_id" binding:"required"`
	ProvinceId       string `json:"province_id" binding:"required"`
	Price            int    `json:"price"`
	Active           uint64 `json:"active"`
	CountryId        string `json:"country_id" binding:"required"`
	WarehouseID      string `json:"warehouse_id"`
	CourierSubtypeId string `json:"courier_subtype_id" binding:"required"`
	CourierTypeId    string `json:"courier_type_id" binding:"required"`
}
type DeliveryFeeModelUpdate struct {
	ID               string `json:"id" binding:"required"`
	SubdistrictId    string `json:"subdistrict_id"`
	DistrictId       string `json:"district_id"`
	CityId           string `json:"city_id" binding:"required"`
	ProvinceId       string `json:"province_id" binding:"required"`
	Price            int    `json:"price"`
	Active           uint64 `json:"active"`
	CountryId        string `json:"country_id" binding:"required"`
	WarehouseID      string `json:"warehouse_id"`
	CourierSubtypeId string `json:"courier_subtype_id" binding:"required"`
	CourierTypeId    string `json:"courier_type_id" binding:"required"`
}
type DeliveryFeeModelInfo struct {
	ID string `json:"id" binding:"required"`
}
