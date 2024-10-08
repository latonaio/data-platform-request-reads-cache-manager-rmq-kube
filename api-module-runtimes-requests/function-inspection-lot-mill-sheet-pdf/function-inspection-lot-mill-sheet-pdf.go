package apiModuleRuntimesRequestsInspectionLotMillSheetPdf

import (
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type InspectionLotMillSheetPdfReq struct {
	Header                []Header                `json:"Header"`
	SpecDetails           []SpecDetails           `json:"SpecDetail"`
	ComponentCompositions []ComponentCompositions `json:"ComponentComposition"`
	Inspections           []Inspections           `json:"Inspection"`
	Accepter              []string                `json:"accepter"`
}

type Header struct {
	InspectionLot                      int    `json:"InspectionLot"`
	InspectionPlantBusinessPartnerName string `json:"InspectionPlantBusinessPartnerName"`
	InspectionLotDate                  string `json:"InspectionLotDate"`
	InspectionSpecification            string `json:"InspectionSpecification"`
	Product                            string `json:"Product"`
	ProductionOrder                    *int   `json:"ProductionOrder"`
	ProductionOrderItem                *int   `json:"ProductionOrderItem"`
}

type SpecDetails struct {
	InspectionLot    int      `json:"InspectionLot"`
	SpecType         string   `json:"SpecType"`
	UpperLimitValue  *float32 `json:"UpperLimitValue"`
	LowerLimitValue  *float32 `json:"LowerLimitValue"`
	StandardValue    *float32 `json:"StandardValue"`
	SpecTypeUnit     *string  `json:"SpecTypeUnit"`
	HeatNumber       *string  `json:"HeatNumber"`
	SpecTypeTextInJA string   `json:"SpecTypeTextInJA"`
}

type ComponentCompositions struct {
	InspectionLot                              int      `json:"InspectionLot"`
	ComponentCompositionType                   string   `json:"ComponentCompositionType"`
	ComponentCompositionUpperLimitInPercent    *float32 `json:"ComponentCompositionUpperLimitInPercent"`
	ComponentCompositionLowerLimitInPercent    *float32 `json:"ComponentCompositionLowerLimitInPercent"`
	ComponentCompositionStandardValueInPercent *float32 `json:"ComponentCompositionStandardValueInPercent"`
	HeatNumber                                 *string  `json:"HeatNumber"`
}

type Inspections struct {
	InspectionLot                            int      `json:"InspectionLot"`
	Inspection                               int      `json:"Inspection"`
	InspectionType                           string   `json:"InspectionType"`
	InspectionTypeCertificateValueInText     *string  `json:"InspectionTypeCertificateValueInText"`
	InspectionTypeCertificateValueInQuantity *float32 `json:"InspectionTypeCertificateValueInQuantity"`
	InspectionTypeValueUnit                  *string  `json:"InspectionTypeValueUnit"`
	InspectionTypeTextInJA                   string   `json:"InspectionTypeTextInJA"`
}

func FunctionInspectionLotMillSheetPdfGenerates(
	input apiOutputFormatter.InspectionLot,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_FUNCTION_INSPECTION_LOT_MILL_SHEET_PDF_SRV"
	aPIType := "generates"

	var request InspectionLotMillSheetPdfReq

	if accepter == "InspectionLotMillSheet" {
		request = InspectionLotMillSheetPdfReq{
			Header:                make([]Header, len(input.InspectionLotHeader)),
			SpecDetails:           make([]SpecDetails, len(input.InspectionLotSpecDetail)),
			ComponentCompositions: make([]ComponentCompositions, len(input.InspectionLotComponentComposition)),
			Inspections:           make([]Inspections, len(input.InspectionLotInspection)),
			Accepter: []string{
				"InspectionLotMillSheet",
			},
		}

		for i, header := range input.InspectionLotHeader {
			request.Header[i] = Header{
				InspectionLot:                      header.InspectionLot,
				InspectionPlantBusinessPartnerName: header.InspectionPlantBusinessPartnerName,
				InspectionLotDate:                  header.InspectionLotDate,
				InspectionSpecification:            *header.InspectionSpecification,
				Product:                            header.Product,
				ProductionOrder:                    header.ProductionOrder,
				ProductionOrderItem:                header.ProductionOrderItem,
			}
		}

		for i, specDetails := range input.InspectionLotSpecDetail {
			request.SpecDetails[i] = SpecDetails{
				InspectionLot:   specDetails.InspectionLot,
				SpecType:        specDetails.SpecType,
				UpperLimitValue: &specDetails.UpperLimitValue,
				LowerLimitValue: &specDetails.LowerLimitValue,
				StandardValue:   &specDetails.StandardValue,
				SpecTypeUnit:    &specDetails.SpecTypeUnit,
				//HeatNumber:       specDetails.HeatNumber,
				//SpecTypeTextInJA: specDetails.SpecTypeTextInJA,
			}
		}

		for i, componentCompositions := range input.InspectionLotComponentComposition {
			request.ComponentCompositions[i] = ComponentCompositions{
				InspectionLot:                              componentCompositions.InspectionLot,
				ComponentCompositionType:                   componentCompositions.ComponentCompositionType,
				ComponentCompositionUpperLimitInPercent:    &componentCompositions.ComponentCompositionUpperLimitInPercent,
				ComponentCompositionLowerLimitInPercent:    &componentCompositions.ComponentCompositionLowerLimitInPercent,
				ComponentCompositionStandardValueInPercent: &componentCompositions.ComponentCompositionStandardValueInPercent,
				//HeatNumber:                                 componentCompositions.HeatNumber,
			}
		}

		for i, inspections := range input.InspectionLotInspection {
			request.Inspections[i] = Inspections{
				InspectionLot:                            inspections.InspectionLot,
				Inspection:                               inspections.Inspection,
				InspectionType:                           inspections.InspectionType,
				InspectionTypeCertificateValueInText:     inspections.InspectionTypeCertificateValueInText,
				InspectionTypeCertificateValueInQuantity: inspections.InspectionTypeCertificateValueInQuantity,
				InspectionTypeValueUnit:                  inspections.InspectionTypeValueUnit,
				//InspectionTypeTextInJA:                 inspections.InspectionTypeTextInJA,
			}
		}
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
		nil,
	)

	return responseBody
}
