package models

import "time"

// TableName sets the insert table name for this struct type
func (c *InvoiceModel) TableName() string {
	return "order_invoice"
}

func (c *InvoiceOrderModel) TableName() string {
	return "order_data"
}

func (c *InvoiceDeliveryModel) TableName() string {
	return "order_delivery"
}

func (c *InvoiceProductModel) TableName() string {
	return "order_product"
}

func (c *InvoiceNestedModel) TableName() string {
	return "order_nested_product"
}

func (c *InvoiceBarcodeModel) TableName() string {
	return "order_barcode"
}

func (c *InvoiceTrackingModel) TableName() string {
	return "order_invoice_tracking"
}

func (c *InvoicePaymentModel) TableName() string {
	return "order_payment"
}

type InvoiceModel struct {
	Id                     string     `gorm:"column:id;primary_key" json:"id"`
	InvoiceDate            NullString `gorm:"column:invoice_date" json:"invoice_date"`
	DueDate                NullString `gorm:"column:due_date" json:"due_date"`
	FirstDeliveryDate      NullString `gorm:"column:first_delivery_date" json:"first_delivery_date"`
	FirstDeliveryTimeStart NullString `gorm:"column:first_delivery_time_start" json:"first_delivery_time_start"`
	FirstDeliveryTimeEnd   NullString `gorm:"column:first_delivery_time_end" json:"first_delivery_time_end"`
	Invid                  NullString `gorm:"column:invid" json:"invid"`
	CustomerRetailId       NullString `gorm:"column:customer_retail_id" json:"customer_retail_id"`
	CustomerInfo           NullString `gorm:"column:customer_info" json:"customer_info"`
	CustomerCorporateId    NullString `gorm:"column:customer_corporate_id" json:"customer_corporate_id"`
	ReferralId             NullString `gorm:"column:referral_id" json:"referral_id"`
	CustomerCorporatePicId NullString `gorm:"column:customer_corporate_pic_id" json:"customer_corporate_pic_id"`
	ShortCode              NullString `gorm:"column:short_code" json:"short_code"`
	StatusInv              int        `gorm:"column:status_inv" json:"status_inv"`
	PlatformId             NullString `gorm:"column:platform_id" json:"platform_id"`
	CustomersLevelId       NullString `gorm:"column:customers_level_id" json:"customer_level_id"`
	ContactMethodId        NullString `gorm:"column:contact_method_id" json:"contact_method_id"`
	SumGrandTotal          float64    `gorm:"column:sum_grand_total" json:"sum_grand_total"`
	Note                   string     `gorm:"column:note" json:"note"`
	Inc                    int        `gorm:"<-:false;column:inc" json:"inc"`
	CreatedBy              NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt              time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy              NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt              time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type InvoiceOrderModel struct {
	Id                 string     `gorm:"column:id;primary_key" json:"id"`
	OrderInvoiceId     string     `gorm:"column:order_invoice_id" json:"order_invoice_id"`
	Orderid            int        `gorm:"column:orderid" json:"orderid"`
	StatusOrder        int        `gorm:"column:status_order" json:"status_order"`
	ShortCode          NullString `gorm:"column:short_code" json:"short_code"`
	Notes              NullString `gorm:"column:notes" json:"notes"`
	OnHold             int        `gorm:"column:on_hold" json:"on_hold"`
	TotalPriceDiscount float64    `gorm:"column:total_price_discount" json:"total_price_discount"`
	TotalPrice         float64    `gorm:"column:total_price" json:"total_price"`
	TotalRetailPriceDm float64    `gorm:"column:total_retail_price_dm" json:"total_retail_price_dm"`
	DeliveryFee        float64    `gorm:"column:delivery_fee" json:"delivery_fee"`
	DeliveryDiscount   float64    `gorm:"column:delivery_discount" json:"delivery_discount"`
	PromoValue         float64    `gorm:"column:promo_value" json:"promo_value"`
	PromoDelivery      float64    `gorm:"column:promo_delivery" json:"promo_delivery"`
	GrandTotal         float64    `gorm:"column:grand_total" json:"grand_total"`
	PromoCode          NullString `gorm:"column:promo_code" json:"promo_code"`
	PromoName          NullString `gorm:"column:promo_name" json:"promo_name"`
	PromoInfo          NullString `gorm:"column:promo_info" json:"promo_info"`
	CreatedBy          NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt          time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy          NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type InvoiceDeliveryModel struct {
	Id                string     `gorm:"column:id;primary_key" json:"id"`
	OrderInvoiceId    string     `gorm:"column:order_invoice_id" json:"order_invoice_id"`
	OrderDataId       string     `gorm:"column:order_data_id" json:"order_data_id"`
	CustomerAddressId NullString `gorm:"column:customer_address_id" json:"customer_address_id"`
	WarehouseId       NullString `gorm:"column:warehouse_id" json:"warehouse_id"`
	CourierTypeId     string     `gorm:"column:courier_type_id" json:"courier_type_id"`
	CourierName       string     `gorm:"column:courier_name" json:"courier_name"`
	DeliveryTimeId    string     `gorm:"column:delivery_time_id" json:"delivery_time_id"`
	DeliveryDate      string     `gorm:"column:delivery_date" json:"delivery_date"`
	DeliveryStart     string     `gorm:"column:delivery_start" json:"delivery_start"`
	DeliveryEnd       string     `gorm:"column:delivery_end" json:"delivery_end"`
	SenderName        string     `gorm:"column:sender_name" json:"sender_name"`
	SenderPhone       string     `gorm:"column:sender_phone" json:"sender_phone"`
	OrderDeliveryInfo string     `gorm:"column:order_delivery_info" json:"order_delivery_info"`
	CreatedBy         NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt         time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy         NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt         time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type InvoiceProductModel struct {
	Id               string     `gorm:"column:id;primary_key" json:"id"`
	OrderInvoiceId   string     `gorm:"column:order_invoice_id" json:"order_invoice_id"`
	OrderDataId      string     `gorm:"column:order_data_id" json:"order_data_id"`
	ProductId        string     `gorm:"column:product_id" json:"product_id"`
	ProductModelId   string     `gorm:"column:product_model_id" json:"product_model_id"`
	Name             string     `gorm:"column:name" json:"name"`
	Price            float64    `gorm:"column:price" json:"price"`
	PriceRetail      float64    `gorm:"column:price_retail" json:"price_retail"`
	Discount         float64    `gorm:"column:discount" json:"discount"`
	Qty              int        `gorm:"column:qty" json:"qty"`
	W2               string     `gorm:"column:w2" json:"w2"`
	W4               string     `gorm:"column:w4" json:"w4"`
	W2Info           string     `gorm:"column:w2_info" json:"w2_info"`
	W4Info           string     `gorm:"column:w4_info" json:"w4_info"`
	ProductInfo      string     `gorm:"column:product_info" json:"product_info"`
	ModelInfo        string     `gorm:"column:model_info" json:"model_info"`
	CourierSubtypeId NullString `gorm:"column:courier_subtype_id" json:"courier_subtype_id"`
	CreatedBy        NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt        time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy        NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt        time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type InvoiceNestedModel struct {
	Id             string     `gorm:"column:id;primary_key" json:"id"`
	OrderInvoiceId string     `gorm:"column:order_invoice_id" json:"order_invoice_id"`
	OrderDataId    string     `gorm:"column:order_data_id" json:"order_data_id"`
	OrderProductId string     `gorm:"column:order_product_id" json:"order_product_id"`
	ProductId      string     `gorm:"column:product_id" json:"product_id"`
	ProductModelId string     `gorm:"column:product_model_id" json:"product_model_id"`
	ProductName    string     `gorm:"column:product_name" json:"product_name"`
	Price          float64    `gorm:"column:price" json:"price"`
	PriceRetail    float64    `gorm:"column:price_retail" json:"price_retail"`
	Discount       float64    `gorm:"column:discount" json:"discount"`
	Qty            int        `gorm:"column:qty" json:"qty"`
	ProductInfo    string     `gorm:"column:product_info" json:"product_info"`
	ModelInfo      string     `gorm:"column:model_info" json:"model_info"`
	CreatedBy      NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt      time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy      NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt      time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type InvoiceBarcodeModel struct {
	Id                   string     `gorm:"column:id;primary_key" json:"id"`
	OrderBarcode         int        `gorm:"<-:false;column:order_barcode" json:"order_barcode"`
	OrderNestedProductId string     `gorm:"column:order_nested_product_id" json:"order_nested_product_id"`
	CreatedBy            NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt            time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy            NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt            time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type InvoiceTrackingModel struct {
	Id                  string     `gorm:"column:id;primary_key" json:"id"`
	OrderInvoiceId      string     `gorm:"order_invoice_id" json:"order_invoice_id"`
	TrackingTitleId     string     `gorm:"tracking_title_id" json:"tracking_title_id"`
	TrackingSelectionId string     `gorm:"tracking_selection_id" json:"tracking_selection_id"`
	CreatedBy           NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt           time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy           NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt           time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type InvoicePaymentModel struct {
	Id             string     `gorm:"column:id;primary_key" json:"id"`
	OrderInvoiceId string     `gorm:"column:order_invoice_id" json:"order_invoice_id"`
	PaymentTypeId  string     `gorm:"column:payment_type_id" json:"payment_type_id"`
	PaymentType    string     `gorm:"column:payment_type" json:"payment_type"`
	PaymentDate    string     `gorm:"column:payment_date" json:"payment_date"`
	Value          float64    `gorm:"column:value" json:"value"`
	Status         int        `gorm:"column:status" json:"status"`
	Note           string     `gorm:"column:note" json:"note"`
	PaidOn         NullString `gorm:"column:paid_on" json:"paid_on"`
	PaymentId      string     `gorm:"column:payment_id" json:"payment_id"`
	PaymentInfo    string     `gorm:"column:payment_info" json:"payment_info"`
	PaymentStatus  int        `gorm:"column:payment_status" json:"payment_status"`
	MaxTimePay     NullString `gorm:"column:max_time_pay" json:"max_time_pay"`
	CreatedBy      NullString `gorm:"column:created_by" json:"created_by"`
	CreatedAt      time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy      NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt      time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type InvoiceSaveForm struct {
	InvoiceDate            string                    `json:"invoice_date" binding:"required"`
	DueDate                string                    `json:"due_date"`
	FirstDeliveryDate      string                    `json:"first_delivery_date"`
	FirstDeliveryTimeStart string                    `json:"first_delivery_time_start"`
	FirstDeliveryTimeEnd   string                    `json:"first_delivery_time_end"`
	CustomerRetailId       string                    `json:"customer_retail_id"`
	CustomerInfo           string                    `json:"customer_info"`
	CustomerCorporateId    string                    `json:"customer_corporate_id"`
	ReferralId             string                    `json:"referral_id"`
	CustomerCorporatePicId string                    `json:"customer_corporate_pic_id"`
	ShortCode              string                    `json:"short_code"`
	StatusInv              int                       `json:"status_inv"`
	PlatformId             string                    `json:"platform_id"`
	CustomersLevelId       string                    `json:"customers_level_id"`
	ContactMethodId        string                    `json:"contact_method_id"`
	SumGrandTotal          float64                   `json:"sum_grand_total"`
	Note                   string                    `json:"note"`
	Order                  []InvoiceOrderSaveForm    `json:"order"`
	Tracking               []InvoiceTrackingSaveForm `json:"tracking"`
	Payment                []InvoicePaymentSaveForm  `json:"payment"`
}

type InvoiceOrderSaveForm struct {
	StatusOrder         int                      `json:"status_order"`
	ShortCode           string                   `json:"short_code"`
	Notes               string                   `json:"notes"`
	OnHold              int                      `json:"on_hold"`
	TotalPriceDiscount  float64                  `json:"total_price_discount"`
	TotalPrice          float64                  `json:"total_price"`
	TotalRetailPriceDm  float64                  `json:"total_retail_price_dm"`
	DeliveryFee         float64                  `json:"delivery_fee"`
	DeliveryDiscount    float64                  `json:"delivery_discount"`
	PromoValue          float64                  `json:"promo_value"`
	PromoDelivery       float64                  `json:"promo_delivery"`
	GrandTotal          float64                  `json:"grand_total"`
	PromoCode           string                   `json:"promo_code"`
	PromoName           string                   `json:"promo_name"`
	PromoInfo           string                   `json:"promo_info"`
	DeliveryAddressId   string                   `json:"delivery_address_id"`
	WarehouseId         string                   `json:"warehouse_id"`
	CourierTypeId       string                   `json:"courier_type_id"`
	CourierName         string                   `json:"courier_name"`
	DeliveryTimeId      string                   `json:"delivery_time_id"`
	DeliveryDate        string                   `json:"delivery_date"`
	DeliveryTimeStart   string                   `json:"delivery_time_start"`
	DeliveryTimeEnd     string                   `json:"delivery_time_end"`
	SenderName          string                   `json:"sender_name"`
	SenderPhone         string                   `json:"sender_phone"`
	DeliveryAddressInfo string                   `json:"delivery_address_info"`
	Product             []InvoiceProductSaveForm `json:"product"`
}

type InvoiceAddressSaveForm struct {
	CourierTypeId     string `json:"courier_type_id"`
	CustomerAddressId string `json:"customer_address_id"`
	WarehouseId       string `json:"warehouse_id"`
	DeliveryDate      string `json:"delivery_date"`
	DeliveryStart     string `json:"delivery_start"`
	DeliveryEnd       string `json:"delivery_end"`
	SenderName        string `json:"sender_name"`
	SenderPhone       string `json:"sender_phone"`
	OrderDeliveryInfo string `json:"order_delivery_info"`
	DeliveryTimeId    string `json:"delivery_time_id"`
}

type InvoiceProductSaveForm struct {
	ProductTypeId      string                         `json:"product_type_id"`
	ProductModelDataId string                         `json:"product_model_data_id"`
	ProductId          string                         `json:"product_id"`
	ProductModelId     string                         `json:"product_model_id"`
	Name               string                         `json:"name"`
	Price              float64                        `json:"price"`
	PriceRetail        float64                        `json:"price_retail"`
	PriceDiscount      float64                        `json:"price_discount"`
	W2                 string                         `json:"w2"`
	W4                 string                         `json:"w4"`
	W2Info             string                         `json:"w2_info"`
	W4Info             string                         `json:"w4_info"`
	ProductInfo        string                         `json:"product_info"`
	ModelInfo          string                         `json:"model_info"`
	Qty                int                            `json:"qty"`
	CourierSubtypeId   string                         `json:"courier_subtype_id"`
	Nested             []InvoiceNestedProductSaveForm `json:"nested"`
}

type InvoiceNestedProductSaveForm struct {
	LibraryId             string  `json:"library_id"`
	LibraryName           string  `json:"library_name"`
	LibraryPropertiesId   string  `json:"library_properties_id"`
	LibraryPropertiesName string  `json:"library_properties_name"`
	ProductId             string  `json:"product_id"`
	ProductName           string  `json:"product_name"`
	ModelId               string  `json:"model_id"`
	ModelName             string  `json:"model_name"`
	Qty                   int     `json:"qty"`
	Price                 float64 `json:"price"`
	PriceRetail           float64 `json:"price_retail"`
	Discount              float64 `json:"discount"`
	TotalPrice            float64 `json:"total_price"`
	Note                  string  `json:"note"`
	IsPackage             int     `json:"is_package"`
	ProductInfo           string  `json:"product_info"`
	ModelInfo             string  `json:"model_info"`
}

type InvoiceTrackingSaveForm struct {
	TrackingTitleId     string `json:"tracking_title_id"`
	TrackingSelectionId string `json:"tracking_selection_id"`
}

type InvoicePaymentSaveForm struct {
	PaymentTypeId string  `json:"payment_type_id"`
	PaymentType   string  `json:"payment_type"`
	PaymentDate   string  `json:"payment_date"`
	Value         float64 `json:"value"`
	Status        int     `json:"status"`
	Note          string  `json:"note"`
	PaidOn        string  `json:"paid_on"`
	PaymentId     string  `json:"payment_id"`
	PaymentInfo   string  `json:"payment_info"`
	PaymentStatus int     `json:"payment_status"`
	MaxTimePay    string  `json:"max_time_pay"`
}

type InvoiceListSearchForm struct {
	FirstDeliveryDateStart string `json:"first_delivery_date_start"`
	FirstDeliveryDateEnd   string `json:"first_delivery_date_end"`
	InvoiceDateStart       string `json:"invoice_date_start"`
	InvoiceDateEnd         string `json:"invoice_date_end"`
	DueDateStart           string `json:"due_date_start"`
	DueDateEnd             string `json:"due_date_end"`
	DataPlatform           string `json:"data_platform"`
	ContactMethod          string `json:"contact_method"`
	CustomerId             string `json:"customer_id"`
	FirstHandler           string `json:"first_handler"`
	InvoicePayment         string `json:"invoice_payment"`
	InvoiceStatus          string `json:"invoice_status"`
}

// Used in View Invoice

type InvoiceListData struct {
	Id                  string  `json:"id"`
	FirstDeliveryDate   string  `json:"first_delivery_date"`
	InvoiceDate         string  `json:"invoice_date"`
	DueDate             string  `json:"due_date"`
	Invid               string  `json:"invid"`
	Inc                 int     `json:"inc"`
	CustomerType        string  `json:"customer_type"`
	CustomerRetailId    string  `json:"customer_retail_id"`
	CustomerCorporateId string  `json:"customer_corporate_id"`
	CustomerInfo        string  `json:"customer_info"`
	GrandTotal          float64 `json:"grand_total"`
	Paid                float64 `json:"paid"`
	Status              int     `json:"status"`
}

type InvoiceInfoForm struct {
	ID string `json:"id" binding:"required"`
}

type InvoiceInfoData struct {
	Id                  string  `json:"id"`
	Invid               string  `json:"invid"`
	InvoiceDate         string  `json:"invoice_date"`
	DueDate             string  `json:"due_date"`
	PlatformId          string  `json:"platform_id"`
	PlatformName        string  `json:"platform_name"`
	CustomerRetailId    string  `json:"customer_retail_id"`
	CustomerCorporateId string  `json:"customer_corporate_id"`
	CustomerInfo        string  `json:"customer_info"`
	ShortCode           string  `json:"short_code"`
	StatusInv           int     `json:"status_inv"`
	ContactMethodId     string  `json:"contact_method_id"`
	GrandTotal          float64 `json:"grand_total"`
	Note                string  `json:"note"`
	CreatedBy           string  `json:"created_by"`
	CreatedByName       string  `json:"created_by_name"`
	UpdatedBy           string  `json:"updated_by"`
	UpdatedByName       string  `json:"updated_by_name"`
}

type InvoiceInfoOrderData struct {
	Id                 string  `json:"id"`
	OrderInvoiceId     string  `json:"order_invoice_id"`
	Orderid            int     `json:"orderid"`
	StatusOrder        int     `json:"status_order"`
	ShortCode          string  `json:"short_code"`
	Notes              string  `json:"notes"`
	OnHold             int     `json:"on_hold"`
	TotalPriceDiscount float64 `json:"total_price_discount"`
	TotalPrice         float64 `json:"total_price"`
	TotalRetailPriceDm float64 `json:"total_retail_price_dm"`
	DeliveryFee        float64 `json:"delivery_fee"`
	DeliveryDiscount   float64 `json:"delivery_discount"`
	PromoValue         float64 `json:"promo_value"`
	PromoDelivery      float64 `json:"promo_delivery"`
	GrandTotal         float64 `json:"grand_total"`
	PromoCode          string  `json:"promo_code"`
	PromoName          string  `json:"promo_name"`
	PromoInfo          string  `json:"promo_info"`
	CustomerAddressId  string  `json:"customer_address_id"`
	WarehouseId        string  `json:"warehouse_id"`
	OrderDeliveryInfo  string  `json:"order_delivery_info"`
	CourierTypeId      string  `json:"courier_type_id"`
	CourierSubtypeId   string  `json:"courier_subtype_id"`
	DeliveryDate       string  `json:"delivery_date"`
	DeliveryTimeId     string  `json:"delivery_time_id"`
	DeliveryStart      string  `json:"delivery_start"`
	DeliveryEnd        string  `json:"delivery_end"`
	SenderName         string  `json:"sender_name"`
	SenderPhone        string  `json:"sender_phone"`
	StatusDeliver      string  `json:"status_deliver"`
}

type InvoiceInfoProductData struct {
	Id             string  `json:"id"`
	OrderInvoiceId string  `json:"order_invoice_id"`
	OrderDataId    string  `json:"order_data_id"`
	ProductId      string  `json:"product_id"`
	ProductModelId string  `json:"product_model_id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	PriceRetail    float64 `json:"price_retail"`
	Discount       float64 `json:"discount"`
	Qty            int     `json:"qty"`
	W2             string  `json:"w2"`
	W4             string  `json:"w4"`
	W2Info         string  `json:"w2_info"`
	W4Info         string  `json:"w4_info"`
	ProductInfo    string  `json:"product_info"`
	ModelInfo      string  `json:"model_info"`
}

type InvoiceInfoNestedData struct {
	Id             string  `json:"id"`
	OrderInvoiceId string  `json:"order_invoice_id"`
	OrderDataId    string  `json:"order_data_id"`
	OrderProductId string  `json:"order_product_id"`
	ProductId      string  `json:"product_id"`
	ProductModelId string  `json:"product_model_id"`
	ProductName    string  `json:"product_name"`
	ProductContent string  `json:"product_content"`
	OrderBarcode   string  `json:"order_barcode"`
	Status         string  `json:"status"`
	Price          float64 `json:"price"`
	PriceRetail    float64 `json:"price_retail"`
	Discount       float64 `json:"discount"`
	Qty            int     `json:"qty"`
	ProductInfo    string  `json:"product_info"`
	ModelInfo      string  `json:"model_info"`
}

type InvoiceInfoPaymentData struct {
	Id             string  `json:"id"`
	OrderInvoiceId string  `json:"order_invoice_id"`
	PaymentTypeId  string  `json:"payment_type_id"`
	PaymentType    string  `json:"payment_type"`
	PaymentDate    string  `json:"payment_date"`
	Value          float64 `json:"value"`
	Status         int     `json:"status"`
	Note           string  `json:"note"`
	PaidOn         string  `json:"paid_on"`
	PaymentId      string  `json:"payment_id"`
	PaymentInfo    string  `json:"payment_info"`
	PaymentStatus  int     `json:"payment_status"`
	MaxTimePay     string  `json:"max_time_pay"`
}

type InvoiceInfoTrackingData struct {
	Id                    string `json:"id"`
	OrderInvoiceId        string `json:"order_invoice_id"`
	TrackingTitleId       string `json:"tracking_title_id"`
	TrackingTitleName     string `json:"tracking_title_name"`
	TrackingSelectionId   string `json:"tracking_selection_id"`
	TrackingSelectionName string `json:"tracking_selected_name"`
}

func (c *InvoiceInfoData) TableName() string {
	return "oi"
}

func (c *InvoiceEditPaymentModel) TableName() string {
	return "order_invoice"
}

type InvoiceEditPaymentForm struct {
	Id          string                   `json:"id" binding:"required"`
	InvoiceDate string                   `json:"invoice_date" binding:"required"`
	DueDate     string                   `json:"due_date"`
	StatusInv   int                      `json:"status_inv"`
	Note        string                   `json:"note"`
	Payment     []InvoicePaymentSaveForm `json:"payment"`
}

type InvoiceEditPaymentModel struct {
	InvoiceDate NullString `gorm:"column:invoice_date" json:"invoice_date"`
	DueDate     NullString `gorm:"column:due_date" json:"due_date"`
	StatusInv   int        `gorm:"column:status_inv" json:"status_inv"`
	Note        NullString `gorm:"column:note" json:"note"`
	UpdatedBy   NullString `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at"`
}
