package apiModuleRuntimesRequestsInspectionLot

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type InspectionLotReq struct {
	Header   Header   `json:"InspectionLot"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	InspectionLot                  int                    `json:"InspectionLot"`
	InspectionLotDate              string                 `json:"InspectionLotDate"`
	InspectionPlan                 int                    `json:"InspectionPlan"`
	InspectionPlantBusinessPartner int                    `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                string                 `json:"InspectionPlant"`
	Product                        string                 `json:"Product"`
	ProductSpecification           *string                `json:"ProductSpecification"`
	InspectionSpecification        *string                `json:"InspectionSpecification"`
	ProductionOrder                *int                   `json:"ProductionOrder"`
	ProductionOrderItem            *int                   `json:"ProductionOrderItem"`
	InspectionLotHeaderText        *string                `json:"InspectionLotHeaderText"`
	ExternalReferenceDocument      *string                `json:"ExternalReferenceDocument"`
	CertificateAuthorityChain      *string                `json:"CertificateAuthorityChain"`
	UsageControlChain              *string                `json:"UsageControlChain"`
	CreationDate                   string                 `json:"CreationDate"`
	CreationTime                   string                 `json:"CreationTime"`
	LastChangeDate                 string                 `json:"LastChangeDate"`
	LastChangeTime                 string                 `json:"LastChangeTime"`
	IsReleased                     *bool                  `json:"IsReleased"`
	IsPartiallyConfirmed           *bool                  `json:"IsPartiallyConfirmed"`
	IsConfirmed                    *bool                  `json:"IsConfirmed"`
	IsLocked                       *bool                  `json:"IsLocked"`
	IsCancelled                    *bool                  `json:"IsCancelled"`
	IsMarkedForDeletion            *bool                  `json:"IsMarkedForDeletion"`
	SpecGeneral                    []SpecGeneral          `json:"SpecGeneral"`
	SpecDetail                     []SpecDetail           `json:"SpecDetail"`
	ComponentComposition           []ComponentComposition `json:"ComponentComposition"`
	Inspection                     []Inspection           `json:"Inspection"`
	Operation                      []Operation            `json:"Operation"`
	Partner                        []Partner              `json:"Partner"`
}

type SpecGeneral struct {
	InspectionLot       int    `json:"InspectionLot"`
	HeatNumber          string `json:"HeatNumber"`
	CreationDate        string `json:"CreationDate"`
	CreationTime        string `json:"CreationTime"`
	LastChangeDate      string `json:"LastChangeDate"`
	LastChangeTime      string `json:"LastChangeTime"`
	IsReleased          *bool  `json:"IsReleased"`
	IsLocked            *bool  `json:"IsLocked"`
	IsCancelled         *bool  `json:"IsCancelled"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type SpecDetail struct {
	InspectionLot       int     `json:"InspectionLot"`
	SpecType            string  `json:"SpecType"`
	UpperLimitValue     float32 `json:"UpperLimitValue"`
	LowerLimitValue     float32 `json:"LowerLimitValue"`
	StandardValue       float32 `json:"StandardValue"`
	SpecTypeUnit        string  `json:"SpecTypeUnit"`
	Formula             *string `json:"Formula"`
	CreationDate        string  `json:"CreationDate"`
	CreationTime        string  `json:"CreationTime"`
	LastChangeDate      string  `json:"LastChangeDate"`
	LastChangeTime      string  `json:"LastChangeTime"`
	IsReleased          *bool   `json:"IsReleased"`
	IsLocked            *bool   `json:"IsLocked"`
	IsCancelled         *bool   `json:"IsCancelled"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type ComponentComposition struct {
	InspectionLot                              int     `json:"InspectionLot"`
	ComponentCompositionType                   string  `json:"ComponentCompositionType"`
	ComponentCompositionUpperLimitInPercent    float32 `json:"ComponentCompositionUpperLimitInPercent"`
	ComponentCompositionLowerLimitInPercent    float32 `json:"ComponentCompositionLowerLimitInPercent"`
	ComponentCompositionStandardValueInPercent float32 `json:"ComponentCompositionStandardValueInPercent"`
	CreationDate                               string  `json:"CreationDate"`
	CreationTime                               string  `json:"CreationTime"`
	LastChangeDate                             string  `json:"LastChangeDate"`
	LastChangeTime                             string  `json:"LastChangeTime"`
	IsReleased                                 *bool   `json:"IsReleased"`
	IsLocked                                   *bool   `json:"IsLocked"`
	IsCancelled                                *bool   `json:"IsCancelled"`
	IsMarkedForDeletion                        *bool   `json:"IsMarkedForDeletion"`
}

type Inspection struct {
	InspectionLot                            int      `json:"InspectionLot"`
	Inspection                               int      `json:"Inspection"`
	InspectionDate                           string   `json:"InspectionDate"`
	InspectionType                           string   `json:"InspectionType"`
	InspectionTypeValueUnit                  *string  `json:"InspectionTypeValueUnit"`
	InspectionTypePlannedValue               *float32 `json:"InspectionTypePlannedValue"`
	InspectionTypeCertificateType            *string  `json:"InspectionTypeCertificateType"`
	InspectionTypeCertificateValueInText     *string  `json:"InspectionTypeCertificateValueInText"`
	InspectionTypeCertificateValueInQuantity *float32 `json:"InspectionTypeCertificateValueInQuantity"`
	InspectionLotInspectionText              *string  `json:"InspectionLotInspectionText"`
	CreationDate                             string   `json:"CreationDate"`
	CreationTime                             string   `json:"CreationTime"`
	LastChangeDate                           string   `json:"LastChangeDate"`
	LastChangeTime                           string   `json:"LastChangeTime"`
	IsReleased                               *bool    `json:"IsReleased"`
	IsLocked                                 *bool    `json:"IsLocked"`
	IsCancelled                              *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}

type Operation struct {
	InspectionLot                                   int      `json:"InspectionLot"`
	Operations                                      int      `json:"Operations"`
	OperationsItem                                  int      `json:"OperationsItem"`
	OperationID                                     int      `json:"OperationID"`
	Inspection                                      int      `json:"Inspection"`
	OperationType                                   string   `json:"OperationType"`
	SupplyChainRelationshipID                       int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID               int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID          int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID        int      `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                         string   `json:"Product"`
	Buyer                                           int      `json:"Buyer"`
	Seller                                          int      `json:"Seller"`
	DeliverToParty                                  int      `json:"DeliverToParty"`
	DeliverToPlant                                  string   `json:"DeliverToPlant"`
	DeliverFromParty                                int      `json:"DeliverFromParty"`
	DeliverFromPlant                                string   `json:"DeliverFromPlant"`
	InspectionPlantBusinessPartner                  int      `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                                 string   `json:"InspectionPlant"`
	Sequence                                        int      `json:"Sequence"`
	SequenceText                                    *string  `json:"SequenceText"`
	OperationText                                   string   `json:"OperationText"`
	OperationStatus                                 *string  `json:"OperationStatus"`
	ResponsiblePlannerGroup                         *string  `json:"ResponsiblePlannerGroup"`
	OperationUnit                                   string   `json:"OperationUnit"`
	StandardLotSizeQuantity                         *float32 `json:"StandardLotSizeQuantity"`
	MinimumLotSizeQuantity                          *float32 `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity                          *float32 `json:"MaximumLotSizeQuantity"`
	PlainLongText                                   *string  `json:"PlainLongText"`
	WorkCenter                                      *int     `json:"WorkCenter"`
	CapacityCategoryCode                            *string  `json:"CapacityCategoryCode"`
	OperationCostingRelevancyType                   *string  `json:"OperationCostingRelevancyType"`
	OperationSetupType                              *string  `json:"OperationSetupType"`
	OperationSetupGroupCategory                     *string  `json:"OperationSetupGroupCategory"`
	OperationSetupGroup                             *string  `json:"OperationSetupGroup"`
	OperationReferenceQuantity                      *float32 `json:"OperationReferenceQuantity"`
	MaximumWaitDuration                             *float32 `json:"MaximumWaitDuration"`
	StandardWaitDuration                            *float32 `json:"StandardWaitDuration"`
	MinimumWaitDuration                             *float32 `json:"MinimumWaitDuration"`
	WaitDurationUnit                                *string  `json:"WaitDurationUnit"`
	MaximumQueueDuration                            *float32 `json:"MaximumQueueDuration"`
	StandardQueueDuration                           *float32 `json:"StandardQueueDuration"`
	MinimumQueueDuration                            *float32 `json:"MinimumQueueDuration"`
	QueueDurationUnit                               *string  `json:"QueueDurationUnit"`
	MaximumMoveDuration                             *float32 `json:"MaximumMoveDuration"`
	StandardMoveDuration                            *float32 `json:"StandardMoveDuration"`
	MinimumMoveDuration                             *float32 `json:"MinimumMoveDuration"`
	MoveDurationUnit                                *string  `json:"MoveDurationUnit"`
	StandardDeliveryDuration                        *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit                    *string  `json:"StandardDeliveryDurationUnit"`
	StandardOperationScrapPercent                   *float32 `json:"StandardOperationScrapPercent"`
	PlannedOperationStandardValue                   *float32 `json:"PlannedOperationStandardValue"`
	PlannedOperationLowerValue                      *float32 `json:"PlannedOperationLowerValue"`
	PlannedOperationUpperValue                      *float32 `json:"PlannedOperationUpperValue"`
	PlannedOperationValueUnit                       *string  `json:"PlannedOperationValueUnit"`
	CostElement                                     *string  `json:"CostElement"`
	OperationErlstSchedldExecStrtDte                *string  `json:"OperationErlstSchedldExecStrtDte"`
	OperationErlstSchedldExecStrtTme                *string  `json:"OperationErlstSchedldExecStrtTme"`
	OperationErlstSchedldExecEndDte                 *string  `json:"OperationErlstSchedldExecEndDte"`
	OperationErlstSchedldExecEndTme                 *string  `json:"OperationErlstSchedldExecEndTme"`
	OperationActualExecutionStartDate               *string  `json:"OperationActualExecutionStartDate"`
	OperationActualExecutionStartTime               *string  `json:"OperationActualExecutionStartTime"`
	OperationActualExecutionEndDate                 *string  `json:"OperationActualExecutionEndDate"`
	OperationActualExecutionEndTime                 *string  `json:"OperationActualExecutionEndTime"`
	OperationConfirmedYieldQuantityInBaseUnit       *float32 `json:"OperationConfirmedYieldQuantityInBaseUnit"`
	OperationConfirmedYieldQuantityInProductionUnit *float32 `json:"OperationConfirmedYieldQuantityInProductionUnit"`
	OperationConfirmedYieldQuantityInOperationUnit  *float32 `json:"OperationConfirmedYieldQuantityInOperationUnit"`
	ValidityStartDate                               *string  `json:"ValidityStartDate"`
	ValidityEndDate                                 *string  `json:"ValidityEndDate"`
	CreationDate                                    string   `json:"CreationDate"`
	CreationTime                                    string   `json:"CreationTime"`
	LastChangeDate                                  string   `json:"LastChangeDate"`
	LastChangeTime                                  string   `json:"LastChangeTime"`
	IsReleased                                      *bool    `json:"IsReleased"`
	IsPartiallyConfirmed                            *bool    `json:"IsPartiallyConfirmed"`
	IsConfirmed                                     *bool    `json:"IsConfirmed"`
	IsLocked                                        *bool    `json:"IsLocked"`
	IsCancelled                                     *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion"`
}

type Partner struct {
	InspectionLot           int     `json:"InspectionLot"`
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

func CreateInspectionLotRequestHeader(
	requestPram *apiInputReader.Request,
	inspectionLotHeader *apiInputReader.InspectionLotHeader,
) InspectionLotReq {
	req := InspectionLotReq{
		Header: Header{
			InspectionLot:       inspectionLotHeader.InspectionLot,
			IsReleased:          inspectionLotHeader.IsReleased,
			IsMarkedForDeletion: inspectionLotHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreateInspectionLotRequestHeaders(
	requestPram *apiInputReader.Request,
	inspectionLotHeaders *apiInputReader.InspectionLotHeader,
) InspectionLotReq {
	req := InspectionLotReq{
		Header: Header{
			IsReleased:          inspectionLotHeaders.IsReleased,
			IsMarkedForDeletion: inspectionLotHeaders.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Headers",
		},
	}
	return req
}

func CreateInspectionLotRequestSpecDetail(
	requestPram *apiInputReader.Request,
	inspectionLotSpecDetail *apiInputReader.InspectionLotSpecDetail,
) InspectionLotReq {
	req := InspectionLotReq{
		Header: Header{
			InspectionLot: inspectionLotSpecDetail.InspectionLot,
			SpecDetail: []SpecDetail{
				{
					InspectionLot:       inspectionLotSpecDetail.InspectionLot,
					SpecType:            inspectionLotSpecDetail.SpecType,
					IsReleased:          inspectionLotSpecDetail.IsReleased,
					IsMarkedForDeletion: inspectionLotSpecDetail.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"SpecDetail",
		},
	}
	return req
}

func CreateInspectionLotRequestSpecDetails(
	requestPram *apiInputReader.Request,
	inspectionLotSpecDetails *apiInputReader.InspectionLotSpecDetail,
) InspectionLotReq {
	req := InspectionLotReq{
		Header: Header{
			InspectionLot: inspectionLotSpecDetails.InspectionLot,
			SpecDetail: []SpecDetail{
				{
					InspectionLot:       inspectionLotSpecDetails.InspectionLot,
					IsReleased:          inspectionLotSpecDetails.IsReleased,
					IsMarkedForDeletion: inspectionLotSpecDetails.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"SpecDetails",
		},
	}
	return req
}

func CreateInspectionLotRequestComponentComposition(
	requestPram *apiInputReader.Request,
	inspectionLotComponentComposition *apiInputReader.InspectionLotComponentComposition,
) InspectionLotReq {
	req := InspectionLotReq{
		Header: Header{
			InspectionLot: inspectionLotComponentComposition.InspectionLot,
			ComponentComposition: []ComponentComposition{
				{
					InspectionLot:            inspectionLotComponentComposition.InspectionLot,
					ComponentCompositionType: inspectionLotComponentComposition.ComponentCompositionType,
					IsReleased:               inspectionLotComponentComposition.IsReleased,
					IsMarkedForDeletion:      inspectionLotComponentComposition.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"ComponentComposition",
		},
	}
	return req
}

func CreateInspectionLotRequestComponentCompositions(
	requestPram *apiInputReader.Request,
	inspectionLotComponentCompositions *apiInputReader.InspectionLotComponentComposition,
) InspectionLotReq {
	req := InspectionLotReq{
		Header: Header{
			InspectionLot: inspectionLotComponentCompositions.InspectionLot,
			ComponentComposition: []ComponentComposition{
				{
					InspectionLot:       inspectionLotComponentCompositions.InspectionLot,
					IsReleased:          inspectionLotComponentCompositions.IsReleased,
					IsMarkedForDeletion: inspectionLotComponentCompositions.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"ComponentCompositions",
		},
	}
	return req
}

func CreateInspectionLotRequestInspection(
	requestPram *apiInputReader.Request,
	inspectionLotInspection *apiInputReader.InspectionLotInspection,
) InspectionLotReq {
	req := InspectionLotReq{
		Header: Header{
			InspectionLot: inspectionLotInspection.InspectionLot,
			Inspection: []Inspection{
				{
					InspectionLot:       inspectionLotInspection.InspectionLot,
					Inspection:          inspectionLotInspection.Inspection,
					IsReleased:          inspectionLotInspection.IsReleased,
					IsMarkedForDeletion: inspectionLotInspection.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Inspection",
		},
	}
	return req
}

func CreateInspectionLotRequestInspections(
	requestPram *apiInputReader.Request,
	inspectionLotInspections *apiInputReader.InspectionLotInspection,
) InspectionLotReq {
	req := InspectionLotReq{
		Header: Header{
			InspectionLot: inspectionLotInspections.InspectionLot,
			Inspection: []Inspection{
				{
					InspectionLot:       inspectionLotInspections.InspectionLot,
					IsReleased:          inspectionLotInspections.IsReleased,
					IsMarkedForDeletion: inspectionLotInspections.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Inspections",
		},
	}
	return req
}

func CreateInspectionLotRequestPartner(
	requestPram *apiInputReader.Request,
	inspectionLotPartner *apiInputReader.InspectionLotPartner,
) InspectionLotReq {
	req := InspectionLotReq{
		Header: Header{
			InspectionLot: inspectionLotPartner.InspectionLot,
			Partner: []Partner{
				{
					PartnerFunction: inspectionLotPartner.PartnerFunction,
					BusinessPartner: inspectionLotPartner.BusinessPartner,
				},
			},
			//IsReleased:          inspectionLotHeader.IsReleased,
			//IsMarkedForDeletion: inspectionLotHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Partner",
		},
	}
	return req
}

func CreateInspectionLotRequestPartners(
	requestPram *apiInputReader.Request,
	inspectionLotPartner *apiInputReader.InspectionLotPartner,
) InspectionLotReq {
	req := InspectionLotReq{
		Header: Header{
			//IsReleased:          inspectionLotHeader.IsReleased,
			//IsMarkedForDeletion: inspectionLotHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Partners",
		},
	}
	return req
}

func InspectionLotReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.InspectionLot,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_INSPECTION_LOT_SRV"
	aPIType := "reads"

	var request InspectionLotReq

	if accepter == "Header" {
		request = CreateInspectionLotRequestHeader(
			requestPram,
			&apiInputReader.InspectionLotHeader{
				InspectionLot:       input.InspectionLotHeader.InspectionLot,
				IsReleased:          input.InspectionLotHeader.IsReleased,
				IsMarkedForDeletion: input.InspectionLotHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Headers" {
		request = CreateInspectionLotRequestHeaders(
			requestPram,
			&apiInputReader.InspectionLotHeader{
				IsReleased:          input.InspectionLotHeader.IsReleased,
				IsMarkedForDeletion: input.InspectionLotHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "SpecDetail" {
		request = CreateInspectionLotRequestSpecDetail(
			requestPram,
			&apiInputReader.InspectionLotSpecDetail{
				InspectionLot:       input.InspectionLotSpecDetail.InspectionLot,
				SpecType:            input.InspectionLotSpecDetail.SpecType,
				IsReleased:          input.InspectionLotSpecDetail.IsReleased,
				IsMarkedForDeletion: input.InspectionLotSpecDetail.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "SpecDetails" {
		request = CreateInspectionLotRequestSpecDetails(
			requestPram,
			&apiInputReader.InspectionLotSpecDetail{
				InspectionLot:       input.InspectionLotSpecDetails.InspectionLot,
				IsReleased:          input.InspectionLotSpecDetails.IsReleased,
				IsMarkedForDeletion: input.InspectionLotSpecDetails.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "ComponentComposition" {
		request = CreateInspectionLotRequestComponentComposition(
			requestPram,
			&apiInputReader.InspectionLotComponentComposition{
				InspectionLot:            input.InspectionLotComponentComposition.InspectionLot,
				ComponentCompositionType: input.InspectionLotComponentComposition.ComponentCompositionType,
				IsReleased:               input.InspectionLotComponentComposition.IsReleased,
				IsMarkedForDeletion:      input.InspectionLotComponentComposition.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "ComponentCompositions" {
		request = CreateInspectionLotRequestComponentCompositions(
			requestPram,
			&apiInputReader.InspectionLotComponentComposition{
				InspectionLot:       input.InspectionLotComponentCompositions.InspectionLot,
				IsReleased:          input.InspectionLotComponentCompositions.IsReleased,
				IsMarkedForDeletion: input.InspectionLotComponentCompositions.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Inspection" {
		request = CreateInspectionLotRequestInspection(
			requestPram,
			&apiInputReader.InspectionLotInspection{
				InspectionLot:       input.InspectionLotInspection.InspectionLot,
				Inspection:          input.InspectionLotInspection.Inspection,
				IsReleased:          input.InspectionLotInspection.IsReleased,
				IsMarkedForDeletion: input.InspectionLotInspection.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Inspections" {
		request = CreateInspectionLotRequestInspections(
			requestPram,
			&apiInputReader.InspectionLotInspection{
				InspectionLot:       input.InspectionLotInspections.InspectionLot,
				IsReleased:          input.InspectionLotInspections.IsReleased,
				IsMarkedForDeletion: input.InspectionLotInspections.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Partner" {
		request = CreateInspectionLotRequestPartner(
			requestPram,
			&apiInputReader.InspectionLotPartner{
				InspectionLot:   input.InspectionLotPartner.InspectionLot,
				PartnerFunction: input.InspectionLotPartner.PartnerFunction,
				BusinessPartner: input.InspectionLotPartner.BusinessPartner,
			},
		)
	}

	if accepter == "Partners" {
		request = CreateInspectionLotRequestPartners(
			requestPram,
			&apiInputReader.InspectionLotPartner{},
		)
	}

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
