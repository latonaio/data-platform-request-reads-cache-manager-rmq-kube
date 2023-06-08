package dpfm_api_output_formatter

type OrdersDetail struct {
	ProductName                           string                                `json:"ProductName"`
	ProductCode                           string                                `json:"ProductCode"`
	ProductInfo                           []ProductInfo                         `json:"ProductInfo"`
	ProductTag                            *[]ProductTag                         `json:"ProductTag"`
	Images                                Images                                `json:"Images"`
	Stock                                 Stock                                 `json:"Stock"`
	AvailabilityStock                     Stock                                 `json:"AvailabilityStock"`
	OrderQuantityInDelivery               OrderQuantityInDelivery               `json:"OrderQuantityInDelivery"`
	OrderQuantityInBase                   OrderQuantityInBase                   `json:"OrderQuantityInBase"`
	ConfirmedOrderQuantityByPDTAvailCheck ConfirmedOrderQuantityByPDTAvailCheck `json:"ConfirmedOrderQuantityByPDTAvailCheck"`
	Popup                                 Popup                                 `json:"Popup"`
	Pulldown                              OrdersDetailPullDown                  `json:"Pulldown"`
}

type OrdersDetailPullDown struct {
	SupplyChains map[int]Deliver `json:"SupplyChains"`
}

type ProductTagReads struct {
	ConnectionKey string `json:"Connection_key"`
	Result        bool   `json:"Result"`
	RedisKey      string `json:"Redis_key"`
	Filepath      string `json:"Filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"Api_schema"`
	MaterialCode  string `json:"Material_code"`
	Deleted       string `json:"Deleted"`
}

type ProductTag struct {
	Key      string `json:"Key"`
	DocCount int    `json:"Doc_count"`
}

type OrderQuantityInDelivery struct {
	Quantity *float32 `json:"Quantity"`
	Unit     string   `json:"Unit"`
}

type OrderQuantityInBase struct {
	Quantity *float32 `json:"Quantity"`
	Unit     string   `json:"Unit"`
}
type ConfirmedOrderQuantityByPDTAvailCheck struct {
	Quantity *float32 `json:"Quantity"`
	Unit     string   `json:"Unit"`
}

type ProductInfo struct {
	KeyName string      `json:"KeyName"`
	Key     string      `json:"Key"`
	Value   interface{} `json:"Value,omitempty"`
}
type Images struct {
	Equipment         *ProductImage `json:"Equipment,omitempty"`
	Product           *ProductImage `json:"Product,omitempty"`
	Barcord           *BarcordImage `json:"Barcode,omitempty"`
	ProductionVersion *ProductImage `json:"ProductionVersion,omitempty"`
	Operations        *ProductImage `json:"Operations,omitempty"`
}

type ProductImage struct {
	BusinessPartnerID int    `json:"BusinessPartnerID"`
	DocID             string `json:"DocID"`
	FileExtension     string `json:"FileExtension"`
}

type BarcordImage struct {
	ID          string `json:"Id"`
	Barcode     string `json:"Barcode"`
	BarcodeType string `json:"BarcodeType"`
}
type Stock struct {
	ProductStock    int    `json:"ProductStock"`
	StorageLocation string `json:"StorageLocation"`
}
type AvailabilityStock struct {
	AvailableProductStock    int    `json:"AvailableProductStock"`
	AvailableStorageLocation string `json:"AvailableStorageLocation"`
}

type Popup struct {
	RequestedDeliveryDate                           string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                           string  `json:"RequestedDeliveryTime"`
	ConfirmedDeliveryDate                           string  `json:"ConfirmedDeliveryDate"`
	ConfirmedDeliveryTime                           string  `json:"ConfirmedDeliveryTime"`
	OrderQuantityInBaseUnit                         int     `json:"OrderQuantityInBaseUnit"`
	BaseUnit                                        string  `json:"BaseUnit"`
	OrderQuantityInDeliveryUnit                     int     `json:"OrderQuantityInDeliveryUnit"`
	DeliveryUnit                                    string  `json:"DeliveryUnit"`
	ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit int     `json:"ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit"`
	DeliverToPlantBatch                             *string `json:"DeliverToPlantBatch"`
	BatchMgmtPolicyInDeliverToPlant                 *string `json:"BatchMgmtPolicyInDeliverToPlant"`
	DeliverToPlantBatchValidityStartDate            *string `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityStartTime            *string `json:"DeliverToPlantBatchValidityStartTime"`
	DeliverToPlantBatchValidityEndDate              *string `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverToPlantBatchValidityEndTime              *string `json:"DeliverToPlantBatchValidityEndTime"`
	DeliverFromPlantBatch                           *string `json:"DeliverFromPlantBatch"`
	BatchMgmtPolicyInDeliverFromPlant               *string `json:"BatchMgmtPolicyInDeliverFromPlant"`
	DeliverFromPlantBatchValidityStartDate          *string `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityStartTime          *string `json:"DeliverFromPlantBatchValidityStartTime"`
	DeliverFromPlantBatchValidityEndDate            *string `json:"DeliverFromPlantBatchValidityEndDate"`
	DeliverFromPlantBatchValidityEndTime            *string `json:"DeliverFromPlantBatchValidityEndTime"`
}
