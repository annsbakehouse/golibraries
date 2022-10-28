package frontend_models

type ProductCollectionView struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ProductCollectionPageFilterTagCategoryView struct {
	ID                         string `gorm:"<-:false;column:id;primary_key" json:"id"`
	ProductFilterTagCategoryID string `gorm:"<-:false;column:mkt_product_filter_tag_category_id" json:"product_filter_tag_category_id"`
}

// TableName sets the insert table name for this struct type
func (c *ProductCollectionView) TableName() string {
	return "mkt_product_collection_page"
}

// TableName sets the insert table name for this struct type
func (c *ProductCollectionPageFilterTagCategoryView) TableName() string {
	return "mkt_product_collection_page_filter_tag_category"
}
