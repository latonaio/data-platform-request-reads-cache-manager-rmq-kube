package apiOutputFormatter

type Orders struct {
	OrdersHeader                                       []OrdersHeader                                       `json:"Header"`
	OrdersHeaderWithItem                               []OrdersHeaderWithItem                               `json:"HeaderWithItem"`
	OrdersPartner                                      []OrdersPartner                                      `json:"Partner"`
	OrdersItem                                         []OrdersItem                                         `json:"Item"`
	OrdersSingleUnit                                   []OrdersSingleUnit                                   `json:"SingleUnit"`
	OrdersItemScheduleLine                             []OrdersItemScheduleLine                             `json:"ItemScheduleLine"`
	OrdersItemPricingElement                           []OrdersItemPricingElement                           `json:"ItemPricingElement"`
	OrdersItemSingleUnitMillSheetHeader                []OrdersItemSingleUnitMillSheetHeader                `json:"OrdersItemSingleUnitMillSheetHeader"`
	OrdersItemSingleUnitMillSheetHeaderInspectionLot   []OrdersItemSingleUnitMillSheetHeaderInspectionLot   `json:"OrdersItemSingleUnitMillSheetHeaderInspectionLot"`
	OrdersItemSingleUnitMillSheetSpecDetails           []OrdersItemSingleUnitMillSheetSpecDetails           `json:"OrdersItemSingleUnitMillSheetSpecDetails"`
	OrdersItemSingleUnitMillSheetComponentCompositions []OrdersItemSingleUnitMillSheetComponentCompositions `json:"OrdersItemSingleUnitMillSheetComponentCompositions"`
	OrdersItemSingleUnitMillSheetInspections           []OrdersItemSingleUnitMillSheetInspections           `json:"OrdersItemSingleUnitMillSheetInspections"`
	MillSheetPdfMountPath                              *string                                              `json:"MillSheetPdfMountPath"`
	MountPath                                          *string                                              `json:"mount_path"`
	Accepter                                           []string                                             `json:"Accepter"`
}

type OrdersHeader struct {
	OrderID              int     `json:"OrderID"`
	Buyer                int     `json:"Buyer"`
	BuyerName            string  `json:"BuyerName"`
	Seller               int     `json:"Seller"`
	SellerName           string  `json:"SellerName"`
	HeaderDeliveryStatus *string `json:"HeaderDeliveryStatus"`
	OrderType            string  `json:"OrderType"`
	IsCancelled          *bool   `json:"IsCancelled"`
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
}

type OrdersPartner struct {
	OrderID                 int     `json:"OrderID"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
}

type OrdersHeaderWithItem struct {
	OrderID                         int     `json:"OrderID"`
	OrderStatus                     string  `json:"OrderStatus"`
	OrderDate                       string  `json:"OrderDate"`
	OrderType                       string  `json:"OrderType"`
	Buyer                           int     `json:"Buyer"`
	BuyerName                       string  `json:"BuyerName"`
	Seller                          int     `json:"Seller"`
	SellerName                      string  `json:"SellerName"`
	RequestedDeliveryDate           string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime           string  `json:"RequestedDeliveryTime"`
	TotalGrossAmount                float32 `json:"TotalGrossAmount"`
	Contract                        *int    `json:"Contract"`
	ContractItem                    *int    `json:"ContractItem"`
	Project                         *int    `json:"Project"`
	ProjectDescription              *string `json:"ProjectDescription"`
	WBSElement                      *int    `json:"WBSElement"`
	WBSElementResponsiblePersonName *string `json:"WBSElementResponsiblePersonName"`
	Incoterms                       *string `json:"Incoterms"`
	IncotermsName                   *string `json:"IncotermsName"`
	PricingDate                     string  `json:"PricingDate"`
	PaymentTerms                    string  `json:"PaymentTerms"`
	PaymentTermsName                string  `json:"PaymentTermsName"`
	PaymentMethod                   string  `json:"PaymentMethod"`
	TransactionCurrency             string  `json:"TransactionCurrency"`
	HeaderText						*string	`json:"HeaderText"`
}

type OrdersItem struct {
	OrderID                     int      `json:"OrderID"`
	OrderItem                   int      `json:"OrderItem"`
	OrderStatus                 string   `json:"OrderStatus"`
	OrderItemCategory           string   `json:"OrderItemCategory"`
	Buyer                       int      `json:"Buyer"`
	BuyerName                   string   `json:"BuyerName"`
	Seller                      int      `json:"Seller"`
	SellerName                  string   `json:"SellerName"`
	DeliverToParty              int		 `json:"DeliverToParty"`
	DeliverToPartyName          string   `json:"DeliverToPartyName"`
	DeliverToPlant              string   `json:"DeliverToPlant"`
	DeliverToPlantName          string   `json:"DeliverToPlantName"`
	DeliverFromParty            int		 `json:"DeliverFromParty"`
	DeliverFromPartyName        string   `json:"DeliverFromPartyName"`
	DeliverFromPlant            string   `json:"DeliverFromPlant"`
	DeliverFromPlantName        string   `json:"DeliverFromPlantName"`
	Product                     string   `json:"Product"`
	ProductSpecification        *string  `json:"ProductSpecification"`
	SizeOrDimensionText         *string  `json:"SizeOrDimensionText"`
	OrderItemText               string   `json:"OrderItemText"`
	OrderItemTextByBuyer        string   `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller       string   `json:"OrderItemTextBySeller"`
	OrderQuantityInBaseUnit     float32  `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit float32  `json:"OrderQuantityInDeliveryUnit"`
	BaseUnit                    string   `json:"BaseUnit"`
	DeliveryUnit                string   `json:"DeliveryUnit"`
	RequestedDeliveryDate       string   `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime       string   `json:"RequestedDeliveryTime"`
	NetAmount                   float32  `json:"NetAmount"`
	TaxAmount                   float32  `json:"TaxAmount"`
	GrossAmount                 float32  `json:"GrossAmount"`
	ProductNetWeight            *float32 `json:"ProductNetWeight"`
	IsCancelled                 *bool    `json:"IsCancelled"`
	IsMarkedForDeletion         *bool    `json:"IsMarkedForDeletion"`
	Images                      Images   `json:"Images"`
}

type OrdersSingleUnit struct {
	OrderID               int     `json:"OrderID"`
	OrderItem             int     `json:"OrderItem"`
	OrderStatus           string  `json:"OrderStatus"`
	Product               string  `json:"Product"`
	RequestedDeliveryDate string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime string  `json:"RequestedDeliveryTime"`
	GrossAmount           float32 `json:"GrossAmount"`
	TransactionCurrency   string  `json:"TransactionCurrency"`
	OrderType             *string `json:"OrderType"`
	Buyer                 int     `json:"Buyer"`
	BuyerName             string  `json:"BuyerName"`
	Seller                int     `json:"Seller"`
	SellerName            string  `json:"SellerName"`
	OrderItemTextByBuyer  string  `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller string  `json:"OrderItemTextBySeller"`
	ConditionCurrency     string  `json:"ConditionCurrency"`
	Images                Images  `json:"Images"`
}

type OrdersItemScheduleLine struct {
	OrderID                                         int      `json:"OrderID"`
	OrderItem                                       int      `json:"OrderItem"`
	ScheduleLine                                    int      `json:"ScheduleLine"`
	Product                                         string   `json:"Product"`
	RequestedDeliveryDate                           string   `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                           string   `json:"RequestedDeliveryTime"`
	Buyer                                           *int     `json:"Buyer"`
	BuyerName                                       string   `json:"BuyerName"`
	Seller                                          *int     `json:"Seller"`
	SellerName                                      string   `json:"SellerName"`
	ScheduleLineOrderQuantityInBaseUnit             float32  `json:"ScheduleLineOrderQuantityInBaseUnit"`
	ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit *float32 `json:"ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit"`
	StockConfirmationBusinessPartner                int      `json:"StockConfirmationBusinessPartner"`
	StockConfirmationBusinessPartnerName            string   `json:"StockConfirmationBusinessPartnerName"`
	StockConfirmationPlant                          string   `json:"StockConfirmationPlant"`
	StockConfirmationPlantName                      string   `json:"StockConfirmationPlantName"`
	DeliveredQuantityInBaseUnit                     *float32 `json:"DeliveredQuantityInBaseUnit"`
	UndeliveredQuantityInBaseUnit                   *float32 `json:"UndeliveredQuantityInBaseUnit"`
}

type OrdersItemPricingElement struct {
	OrderID                 int     `json:"OrderID"`
	OrderItem               int     `json:"OrderItem"`
	PricingProcedureCounter int     `json:"PricingProcedureCounter"`
	ConditionRateValue      float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit  int     `json:"ConditionRateValueUnit"`
	ConditionScaleQuantity  int     `json:"ConditionScaleQuantity"`
	ConditionCurrency       string  `json:"ConditionCurrency"`
	ConditionQuantity       float32 `json:"ConditionQuantity"`
	ConditionAmount         float32 `json:"ConditionAmount"`
	ConditionType           string  `json:"ConditionType"`
}

type OrdersItemSingleUnitMillSheetHeader struct {
	OrderID                 int     `json:"OrderID"`
	OrderItem               int     `json:"OrderItem"`
	OrderType               *string `json:"OrderType"`
	OrderStatus             string  `json:"OrderStatus"`
	Buyer                   int     `json:"Buyer"`
	BuyerName               string  `json:"BuyerName"`
	Seller                  int     `json:"Seller"`
	SellerName              string  `json:"SellerName"`
	Product                 string  `json:"Product"`
	SizeOrDimensionText     string  `json:"SizeOrDimensionText"`
	OrderItemTextByBuyer    string  `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller   string  `json:"OrderItemTextBySeller"`
	OrderQuantityInBaseUnit float32 `json:"OrderQuantityInBaseUnit"`
	RequestedDeliveryDate   string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime   string  `json:"RequestedDeliveryTime"`
	ProductSpecification    string  `json:"ProductSpecification"`
	MarkingOfMaterial       string  `json:"MarkingOfMaterial"`
	ProductionVersion       *int    `json:"ProductionVersion"`
	ProductionVersionItem   *int    `json:"ProductionVersionItem"`
	ProductionOrder         *int    `json:"ProductionOrder"`
	ProductionOrderItem     *int    `json:"ProductionOrderItem"`
	Contract                *int    `json:"Contract"`
	ContractItem            *int    `json:"ContractItem"`
	Project                 *int    `json:"Project"`
	WBSElement              *int    `json:"WBSElement"`
	GrossAmount             float32 `json:"GrossAmount"`
	//ConditionCurrency     	*string `json:"ConditionCurrency"`
	InspectionLot           int     `json:"InspectionLot"`
	InspectionLotDate       string  `json:"InspectionLotDate"`
	InspectionSpecification *string `json:"InspectionSpecification"`
	//Images                  Images  `json:"Images"`
}

type OrdersItemSingleUnitMillSheetHeaderInspectionLot struct {
	OrderID                 int     `json:"OrderID"`
	OrderItem               int     `json:"OrderItem"`
	InspectionLot           int     `json:"InspectionLot"`
	InspectionLotDate       string  `json:"InspectionLotDate"`
	InspectionSpecification *string `json:"InspectionSpecification"`
}

type OrdersItemSingleUnitMillSheetSpecDetails struct {
	OrderID         int      `json:"OrderID"`
	OrderItem       int      `json:"OrderItem"`
	InspectionLot   int      `json:"InspectionLot"`
	SpecType        string   `json:"SpecType"`
	UpperLimitValue *float32 `json:"UpperLimitValue"`
	LowerLimitValue *float32 `json:"LowerLimitValue"`
	StandardValue   *float32 `json:"StandardValue"`
	SpecTypeUnit    *string  `json:"SpecTypeUnit"`
}

type OrdersItemSingleUnitMillSheetComponentCompositions struct {
	OrderID                                    int      `json:"OrderID"`
	OrderItem                                  int      `json:"OrderItem"`
	InspectionLot                              int      `json:"InspectionLot"`
	ComponentCompositionType                   string   `json:"ComponentCompositionType"`
	ComponentCompositionUpperLimitInPercent    *float32 `json:"ComponentCompositionUpperLimitInPercent"`
	ComponentCompositionLowerLimitInPercent    *float32 `json:"ComponentCompositionLowerLimitInPercent"`
	ComponentCompositionStandardValueInPercent *float32 `json:"ComponentCompositionStandardValueInPercent"`
}

type OrdersItemSingleUnitMillSheetInspections struct {
	OrderID                                  int      `json:"OrderID"`
	OrderItem                                int      `json:"OrderItem"`
	InspectionLot                            int      `json:"InspectionLot"`
	Inspection                               int      `json:"Inspection"`
	InspectionType                           string   `json:"InspectionType"`
	InspectionTypeCertificateValueInText     *string  `json:"InspectionTypeCertificateValueInText"`
	InspectionTypeCertificateValueInQuantity *float32 `json:"InspectionTypeCertificateValueInQuantity"`
	InspectionTypeValueUnit                  *string  `json:"InspectionTypeValueUnit"`
}
