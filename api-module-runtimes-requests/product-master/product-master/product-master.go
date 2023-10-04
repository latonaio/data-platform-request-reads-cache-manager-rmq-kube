package apiModuleRuntimesRequestsProductMaster

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ProductMasterReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	General           General   `json:"ProductMaster"`
	Generals          []General `json:"ProductMasters"`
	Accepter          []string  `json:"accepter"`
}

type General struct {
	Product                       string               `json:"Product"`
	ProductType                   *string              `json:"ProductType"`
	BaseUnit                      *string              `json:"BaseUnit"`
	ValidityStartDate             *string              `json:"ValidityStartDate"`
	ValidityEndDate               *string              `json:"ValidityEndDate"`
	ItemCategory                  *string              `json:"ItemCategory"`
	ProductGroup                  *string              `json:"ProductGroup"`
	GrossWeight                   *float32             `json:"GrossWeight"`
	NetWeight                     *float32             `json:"NetWeight"`
	WeightUnit                    *string              `json:"WeightUnit"`
	InternalCapacityQuantity      *float32             `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit  *string              `json:"InternalCapacityQuantityUnit"`
	SizeOrDimensionText           *string              `json:"SizeOrDimensionText"`
	ProductStandardID             *string              `json:"ProductStandardID"`
	IndustryStandardName          *string              `json:"IndustryStandardName"`
	MarkingOfMaterial             *string              `json:"MarkingOfMaterial"`
	CountryOfOrigin               *string              `json:"CountryOfOrigin"`
	CountryOfOriginLanguage       *string              `json:"CountryOfOriginLanguage"`
	LocalRegionOfOrigin           *string              `json:"LocalRegionOfOrigin"`
	LocalSubRegionOfOrigin        *string              `json:"LocalSubRegionOfOrigin"`
	BarcodeType                   *string              `json:"BarcodeType"`
	ProductAccountAssignmentGroup *string              `json:"ProductAccountAssignmentGroup"`
	CreationDate                  *string              `json:"CreationDate"`
	LastChangeDate                *string              `json:"LastChangeDate"`
	IsMarkedForDeletion           *bool                `json:"IsMarkedForDeletion"`
	ProductDescription            []ProductDescription `json:"ProductDescription"`
	BusinessPartner               []BusinessPartner    `json:"BusinessPartner"`
	GeneralDoc                    []GeneralDoc         `json:"GeneralDoc"`
	Tax                           []Tax                `json:"Tax"`
}

type BusinessPartner struct {
	Product                string            `json:"Product"`
	BusinessPartner        int               `json:"BusinessPartner"`
	ValidityStartDate      string            `json:"ValidityStartDate"`
	ValidityEndDate        string            `json:"ValidityEndDate"`
	BusinessPartnerProduct *string           `json:"BusinessPartnerProduct"`
	CreationDate           *string           `json:"CreationDate"`
	LastChangeDate         *string           `json:"LastChangeDate"`
	IsMarkedForDeletion    *bool             `json:"IsMarkedForDeletion"`
	BPPlant                []BPPlant         `json:"BPPlant"`
	ProductDescByBP        []ProductDescByBP `json:"ProductDescByBP"`
	Allergen               []Allergen        `json:"Allergen"`
	Calories               []Calories        `json:"Calories"`
	NutritionalInfo        []NutritionalInfo `json:"NutritionalInfo"`
}

type ProductDescription struct {
	Product             string            `json:"Product"`
	Language            string            `json:"Language"`
	ProductDescription  *string           `json:"ProductDescription"`
	CreationDate        *string           `json:"CreationDate"`
	LastChangeDate      *string           `json:"LastChangeDate"`
	IsMarkedForDeletion *bool             `json:"IsMarkedForDeletion"`
	ProductDescByBP     []ProductDescByBP `json:"ProductDescByBP"`
}

type ProductDescByBP struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Language            string  `json:"Language"`
	ProductDescription  *string `json:"ProductDescription"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type BPPlant struct {
	Product                                   string            `json:"Product"`
	BusinessPartner                           int               `json:"BusinessPartner"`
	Plant                                     string            `json:"Plant"`
	MRPType                                   *string           `json:"MRPType"`
	MRPController                             *string           `json:"MRPController"`
	ReorderThresholdQuantityInBaseUnit        *float32          `json:"ReorderThresholdQuantityInBaseUnit"`
	PlanningTimeFenceInDays                   *int              `json:"PlanningTimeFenceInDays"`
	MRPPlanningCalendar                       *string           `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32          `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *float32          `json:"SafetyDuration"`
	SafetyDurationUnit                        *string           `json:"SafetyDurationUnit"`
	MaximumStockQuantityInBaseUnit            *float32          `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32          `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32          `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryQuantityInBaseUnit        *float32          `json:"StandardDeliveryQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32          `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32          `json:"MaximumDeliveryQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32          `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32          `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool             `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDuration                  *float32          `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit              *string           `json:"StandardDeliveryDurationUnit"`
	IsBatchManagementRequired                 *bool             `json:"IsBatchManagementRequired"`
	BatchManagementPolicy                     *string           `json:"BatchManagementPolicy"`
	ProfitCenter                              *string           `json:"ProfitCenter"`
	CreationDate                              *string           `json:"CreationDate"`
	LastChangeDate                            *string           `json:"LastChangeDate"`
	IsMarkedForDeletion                       *bool             `json:"IsMarkedForDeletion"`
	MRPArea                                   []MRPArea         `json:"MRPArea"`
	Production                                []Production      `json:"Production"`
	Quality                                   []Quality         `json:"Quality"`
	Accounting                                []Accounting      `json:"Accounting"`
	StorageLocation                           []StorageLocation `json:"StorageLocation"`
	BPPlantDoc                                []BPPlantDoc      `json:"BPPlantDoc"`
}

type MRPArea struct {
	Product                                   string   `json:"Product"`
	BusinessPartner                           int      `json:"BusinessPartner"`
	Plant                                     string   `json:"Plant"`
	MRPArea                                   string   `json:"MRPArea"`
	MRPType                                   *string  `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	StorageLocationForMRP                     *string  `json:"StorageLocationForMRP"`
	ReorderThresholdQuantityInBaseUnit        *float32 `json:"ReorderThresholdQuantityInBaseUnit"`
	PlanningTimeFenceInDays                   *int     `json:"PlanningTimeFenceInDays"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *float32 `json:"SafetyDuration"`
	SafetyDurationUnit                        *string  `json:"SafetyDurationUnit"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinumumDeliveryQuantityInBaseUnit         *float32 `json:"MinumumDeliveryQuantityInBaseUnit"`
	MinumumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinumumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryQuantityInBaseUnit        *float32 `json:"StandardDeliveryQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDuration                  *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit              *string  `json:"StandardDeliveryDurationUnit"`
	CreationDate                              *string  `json:"CreationDate"`
	LastChangeDate                            *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}

type Production struct {
	Product                                           string   `json:"Product"`
	BusinessPartner                                   int      `json:"BusinessPartner"`
	Plant                                             string   `json:"Plant"`
	ProductionStorageLocation                         *string   `json:"ProductionStorageLocation"`
	ProductionDuration                         		  *float32  `json:"ProductionDuration"`
	ProductionDurationUnit                     		  *string   `json:"ProductionDurationUnit"`
	ProductionQuantityUnit                     		  *string   `json:"ProductionQuantityUnit"`
	MinimumProductionQuantityInBaseUnit               *float32  `json:"MinimumProductionQuantityInBaseUnit"`
	MinimumProductionLotSizeQuantityInBaseUnit        *float32  `json:"MinimumProductionLotSizeQuantityInBaseUnit"`
	StandardProductionQuantityInBaseUnit              *float32  `json:"StandardProductionQuantityInBaseUnit"`
	StandardProductionLotSizeQuantityInBaseUnit       *float32  `json:"StandardProductionLotSizeQuantityInBaseUnit"`
	MaximumProductionQuantityInBaseUnit               *float32  `json:"MaximumProductionQuantityInBaseUnit"`
	MaximumProductionLotSizeQuantityInBaseUnit        *float32  `json:"MaximumProductionLotSizeQuantityInBaseUnit"`
	ProductionLotSizeRoundingQuantityInBaseUnit       *float32 `json:"ProductionLotSizeRoundingQuantityInBaseUnit"`
	MinimumProductionQuantityInProductionUnit         *float32  `json:"MinimumProductionQuantityInProductionUnit"`
	MinimumProductionLotSizeQuantityInProductionUnit  *float32  `json:"MinimumProductionLotSizeQuantityInProductionUnit"`
	StandardProductionQuantityInProductionUnit        *float32  `json:"StandardProductionQuantityInProductionUnit"`
	StandardProductionLotSizeQuantityInProductionUnit *float32  `json:"StandardProductionLotSizeQuantityInProductionUnit"`
	MaximumProductionQuantityInProductionUnit         *float32  `json:"MaximumProductionQuantityInProductionUnit"`
	MaximumProductionLotSizeQuantityInProductionUnit  *float32  `json:"MaximumProductionLotSizeQuantityInProductionUnit"`
	ProductionLotSizeRoundingQuantityInProductionUnit *float32 `json:"ProductionLotSizeRoundingQuantityInProductionUnit"`
	ProductionLotSizeIsFixed                          *bool    `json:"ProductionLotSizeIsFixed"`
	ProductIsBatchManagedInProductionPlant            *bool    `json:"ProductIsBatchManagedInProductionPlant"`
	ProductIsMarkedForBackflush                       *bool    `json:"ProductIsMarkedForBackflush"`
	ProductionSchedulingProfile                       *string  `json:"ProductionSchedulingProfile"`
	CreationDate                                      string   `json:"CreationDate"`
	LastChangeDate                                    string   `json:"LastChangeDate"`
	IsMarkedForDeletion                               *bool    `json:"IsMarkedForDeletion"`
}

type Quality struct {
	Product                       string  `json:"Product"`
	BusinessPartner               int     `json:"BusinessPartner"`
	Plant                         string  `json:"Plant"`
	QualityMgmtCtrlKey            *string `json:"QualityMgmtCtrlKey"`
	ProductSpecification          *string `json:"ProductSpecification"`
	MaximumStoragePeriodInDays    *int    `json:"MaximumStoragePeriodInDays"`
	RecrrgInspIntervalTimeInDays  *int    `json:"RecrrgInspIntervalTimeInDays"`
	ProductQualityCertificateType *string `json:"ProductQualityCertificateType"`
	CreationDate                  *string `json:"CreationDate"`
	LastChangeDate                *string `json:"LastChangeDate"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}

type StorageBin struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Plant               string  `json:"Plant"`
	StorageLocation     string  `json:"StorageLocation"`
	StorageBin          string  `json:"StorageBin"`
	BlockStatus         *bool   `json:"BlockStatus"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type GeneralDoc struct {
	Product                  string  `json:"Product"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            *string `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type Accounting struct {
	Product             string   `json:"Product"`
	BusinessPartner     int      `json:"BusinessPartner"`
	Plant               string   `json:"Plant"`
	ValuationClass      *string  `json:"ValuationClass"`
	CostingPolicy       *string  `json:"CostingPolicy"`
	PriceUnitQty        *int     `json:"PriceUnitQty"`
	StandardPrice       *float32 `json:"StandardPrice"`
	MovingAveragePrice  *float32 `json:"MovingAveragePrice"`
	CreationDate        *string  `json:"CreationDate"`
	LastChangeDate      *string  `json:"LastChangeDate"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}

type Tax struct {
	Product                  string  `json:"Product"`
	Country                  string  `json:"Country"`
	ProductTaxCategory       string  `json:"ProductTaxCategory"`
	ProductTaxClassification *string `json:"ProductTaxClassification"`
	CreationDate             *string `json:"CreationDate"`
	LastChangeDate           *string `json:"LastChangeDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

type Allergen struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Allergen            string  `json:"Allergen"`
	AllergenIsContained *bool   `json:"AllergenIsContained"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type Calories struct {
	Product             string   `json:"Product"`
	BusinessPartner     int      `json:"BusinessPartner"`
	Calories            *float32 `json:"Calories"`
	CaloryUnit          *string  `json:"CaloryUnit"`
	CaloryUnitQuantity  *int     `json:"CaloryUnitQuantity"`
	CreationDate        *string  `json:"CreationDate"`
	LastChangeDate      *string  `json:"LastChangeDate"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}

type NutritionalInfo struct {
	Product             string   `json:"Product"`
	BusinessPartner     int      `json:"BusinessPartner"`
	Nutrient            string   `json:"Nutrient"`
	NutrientContent     *float32 `json:"NutrientContent"`
	NutrientContentUnit *string  `json:"NutrientContentUnit"`
	CreationDate        *string  `json:"CreationDate"`
	LastChangeDate      *string  `json:"LastChangeDate"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}

type BPPlantDoc struct {
	Product                  string  `json:"Product"`
	BusinessPartner          int     `json:"BusinessPartner"`
	Plant                    string  `json:"Plant"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            *string `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type StorageLocation struct {
	Product             string       `json:"Product"`
	BusinessPartner     int          `json:"BusinessPartner"`
	Plant               string       `json:"Plant"`
	StorageLocation     string       `json:"StorageLocation"`
	BlockStatus         *bool        `json:"BlockStatus"`
	CreationDate        *string      `json:"CreationDate"`
	LastChangeDate      *string      `json:"LastChangeDate"`
	IsMarkedForDeletion *bool        `json:"IsMarkedForDeletion"`
	StorageBin          []StorageBin `json:"StorageBin"`
}

func CreateProductMasterRequestGeneral(
	requestPram *apiInputReader.Request,
	input General,
) ProductMasterReq {
	req := ProductMasterReq{
		General: General{
			Product: input.Product,
		},
		Accepter: []string{
			"General",
		},
	}
	return req
}

func CreateProductMasterRequestGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductMaster,
) ProductMasterReq {
	isMarkedForDeletion := false

	req := ProductMasterReq{
		General: General{
			//IsMarkedForDeletion: requestPram.IsMarkedForDeletion,
			//IsMarkedForDeletion: &isMarkedForDeletion,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		Accepter: []string{
			"Generals",
		},
	}
	return req
}

func CreateProductMasterProductDescriptions(
	requestPram *apiInputReader.Request,
	generals []General,
) ProductMasterReq {
	req := ProductMasterReq{
		Generals: generals,
		Accepter: []string{
			"ProductDescriptions",
		},
	}
	return req
}

func CreateProductMasterBusinessPartners(
	requestPram *apiInputReader.Request,
	general General,
) ProductMasterReq {
	req := ProductMasterReq{
		General: general,
		Accepter: []string{
			"BusinessPartners",
		},
	}
	return req
}

func CreateProductMasterRequestProductDescsByBP(
	requestPram *apiInputReader.Request,
	generals []General,
) ProductMasterReq {
	req := ProductMasterReq{
		Generals: generals,
		Accepter: []string{
			"ProductDescsByBP",
		},
	}
	return req
}

func CreateProductMasterByBPPlant(
	requestPram *apiInputReader.Request,
	general General,
) ProductMasterReq {
	req := ProductMasterReq{
		General: general,
		Accepter: []string{
			"BPPlant",
		},
	}
	return req
}

func CreateProductMasterByBPPlants(
	requestPram *apiInputReader.Request,
	general General,
) ProductMasterReq {
	req := ProductMasterReq{
		General: general,
		Accepter: []string{
			"BPPlants",
		},
	}
	return req
}

func CreateProductMasterProduction(
	requestPram *apiInputReader.Request,
	general General,
) ProductMasterReq {
	req := ProductMasterReq{
		General: general,
		Accepter: []string{
			"Production",
		},
	}
	return req
}

func ProductMasterReadsGeneral(
	requestPram *apiInputReader.Request,
	input General,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCT_MASTER_SRV"
	aPIType := "reads"

	request := CreateProductMasterRequestGeneral(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}

func ProductMasterReadsGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductMaster,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCT_MASTER_SRV"
	aPIType := "reads"

	request := CreateProductMasterRequestGenerals(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}

func ProductMasterReadsProductDescriptions(
	requestPram *apiInputReader.Request,
	generals []General,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCT_MASTER_SRV"
	aPIType := "reads"

	request := CreateProductMasterProductDescriptions(
		requestPram,
		generals,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}

func ProductMasterReadsBusinessPartners(
	requestPram *apiInputReader.Request,
	general General,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCT_MASTER_SRV"
	aPIType := "reads"

	request := CreateProductMasterBusinessPartners(
		requestPram,
		general,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}

func ProductMasterReadsProductDescsByBP(
	requestPram *apiInputReader.Request,
	generals []General,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCT_MASTER_SRV"
	aPIType := "reads"

	request := CreateProductMasterRequestProductDescsByBP(
		requestPram,
		generals,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}

func ProductMasterReadsByBPPlant(
	requestPram *apiInputReader.Request,
	input General,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCT_MASTER_SRV"
	aPIType := "reads"

	request := CreateProductMasterByBPPlant(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}

func ProductMasterReadsByBPPlants(
	requestPram *apiInputReader.Request,
	input General,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCT_MASTER_SRV"
	aPIType := "reads"

	request := CreateProductMasterByBPPlants(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}

func ProductMasterReadsProduction(
	requestPram *apiInputReader.Request,
	input General,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCT_MASTER_SRV"
	aPIType := "reads"

	request := CreateProductMasterProduction(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}
