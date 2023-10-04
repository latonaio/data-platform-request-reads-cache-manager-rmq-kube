package apiOutputFormatter

type ProductMaster struct {
	ProductMasterGeneral               []ProductMasterGeneral               `json:"Generals"`
	ProductMasterGeneralWithOthers     []ProductMasterGeneralWithOthers     `json:"GeneralWithOthers"`
	ProductMasterDetailGeneral         []ProductMasterDetailGeneral         `json:"DetailGeneral"`
	ProductMasterDetailBusinessPartner []ProductMasterDetailBusinessPartner `json:"DetailBusinessPartner"`
	ProductMasterDetailBPPlant         []ProductMasterDetailBPPlant         `json:"DetailBPPlant"`
	ProductMasterSingleUnit            []ProductMasterSingleUnit            `json:"SingleUnit"`
}

type ProductMasterGeneral struct {
	Product             string  `json:"Product"`
	ProductName         *string `json:"ProductName"`
	ProductGroup        *string `json:"ProductGroup"`
	BaseUnit            string  `json:"BaseUnit"`
	ValidityStartDate   string  `json:"ValidityStartDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Images              Images  `json:"Images"`
}

type ProductMasterGeneralWithOthers struct {
	Product           string  `json:"Product"`
	ProductName       *string `json:"ProductName"`
	ProductGroup      *string `json:"ProductGroup"`
	BaseUnit          string  `json:"BaseUnit"`
	ValidityStartDate string  `json:"ValidityStartDate"`
	Images            Images  `json:"Images"`
}

type ProductMasterDetailGeneral struct {
	ProductType                   string   `json:"ProductType"`
	GrossWeight                   *float32 `json:"GrossWeight"`
	NetWeight                     *float32 `json:"NetWeight"`
	WeightUnit                    *string  `json:"WeightUnit"`
	InternalCapacityQuantity      *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit  *string  `json:"InternalCapacityQuantityUnit"`
	SizeOrDimensionText           *string  `json:"SizeOrDimensionText"`
	ProductStandardID             *string  `json:"ProductStandardID"`
	IndustryStandardName          *string  `json:"IndustryStandardName"`
	CountryOfOrigin               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage       *string  `json:"CountryOfOriginLanguage"`
	BarcodeType                   *string  `json:"BarcodeType"`
	ProductAccountAssignmentGroup *string  `json:"ProductAccountAssignmentGroup"`
	CreationDate                  string   `json:"CreationDate"`
	LastChangeDate                string   `json:"LastChangeDate"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
	Images                        Images   `json:"Images"`
}

type ProductMasterDetailBusinessPartner struct {
	BusinessPartner        int     `json:"BusinessPartner"`
	BusinessPartnerName    string  `json:"BusinessPartnerName"`
	ValidityStartDate      string  `json:"ValidityStartDate"`
	ValidityEndDate        string  `json:"ValidityEndDate"`
	BusinessPartnerProduct *string `json:"BusinessPartnerProduct"`
	CreationDate           string  `json:"CreationDate"`
	LastChangeDate         string  `json:"LastChangeDate"`
	IsMarkedForDeletion    *bool   `json:"IsMarkedForDeletion"`
	Images                 Images  `json:"Images"`
}

type ProductMasterDetailBPPlant struct {
	BusinessPartner                           int      `json:"BusinessPartner"`
	BusinessPartnerName                       string   `json:"BusinessPartnerName"`
	Plant                                     string   `json:"Plant"`
	PlantName                                 string   `json:"PlantName"`
	MRPType                                   string   `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	ReorderThresholdQuantityInBaseUnit        *float32 `json:"ReorderThresholdQuantityInBaseUnit"`
	PlanningTimeFenceInDays                   *int     `json:"PlanningTimeFenceInDays"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *float32 `json:"SafetyDuration"`
	SafetyDurationUnit                        *string  `json:"SafetyDurationUnit"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32 `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryQuantityInBaseUnit        *float32 `json:"StandardDeliveryQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDuration                  *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit              *string  `json:"StandardDeliveryDurationUnit"`
	IsBatchManagementRequired                 *bool    `json:"IsBatchManagementRequired"`
	BatchManagementPolicy                     *string  `json:"BatchManagementPolicy"`
	ProfitCenter                              *string  `json:"ProfitCenter"`
	CreationDate                              string   `json:"CreationDate"`
	LastChangeDate                            string   `json:"LastChangeDate"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
	Images                                    Images   `json:"Images"`
}

type ProductMasterSingleUnit struct {
	Product                       string   `json:"Product"`
	ProductName                   *string  `json:"ProductName"`
	ProductType                   *string  `json:"ProductType"`
	GrossWeight                   *float32 `json:"GrossWeight"`
	NetWeight                     *float32 `json:"NetWeight"`
	WeightUnit                    *string  `json:"WeightUnit"`
	InternalCapacityQuantity      *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit  *string  `json:"InternalCapacityQuantityUnit"`
	SizeOrDimensionText           *string  `json:"SizeOrDimensionText"`
	ProductStandardID             *string  `json:"ProductStandardID"`
	IndustryStandardName          *string  `json:"IndustryStandardName"`
	ItemCategory                  *string  `json:"ItemCategory"`
	CountryOfOrigin               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage       *string  `json:"CountryOfOriginLanguage"`
	LocalRegionOfOrigin           *string  `json:"LocalRegionOfOrigin"`
	LocalSubRegionOfOrigin        *string  `json:"LocalSubRegionOfOrigin"`
	MarkingOfMaterial             *string  `json:"MarkingOfMaterial"`
	BarcodeType                   *string  `json:"BarcodeType"`
	ProductAccountAssignmentGroup *string  `json:"ProductAccountAssignmentGroup"`
	ValidityEndDate               string   `json:"ValidityEndDate"`
	CreationDate                  string   `json:"CreationDate"`
	LastChangeDate                string   `json:"LastChangeDate"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
	Images                        *Images  `json:"Images"`
}
