package apiModuleRuntimesRequestsMillSheetPdf

import (
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type MillSheetPdfReq struct {
	Header                []Header                `json:"OrdersItemSingleUnitMillSheetHeader"`
	HeaderInspectionLot   []HeaderInspectionLot   `json:"OrdersItemSingleUnitMillSheetHeaderInspectionLot"`
	SpecDetails           []SpecDetails           `json:"OrdersItemSingleUnitMillSheetSpecDetails"`
	ComponentCompositions []ComponentCompositions `json:"OrdersItemSingleUnitMillSheetComponentCompositions"`
	Inspections           []Inspections           `json:"OrdersItemSingleUnitMillSheetInspections"`
	Accepter              []string                `json:"accepter"`
}

type Header struct {
	OrderID                            int     `json:"OrderID"`
	OrderItem                          int     `json:"OrderItem"`
	BuyerName                          string  `json:"BuyerName"`
	SellerName                         string  `json:"SellerName"`
	Product                            string  `json:"Product"`
	SizeOrDimensionText                string  `json:"SizeOrDimensionText"`
	OrderQuantityInBaseUnit            float32 `json:"OrderQuantityInBaseUnit"`
	ProductSpecification               string  `json:"ProductSpecification"`
	MarkingOfMaterial                  string  `json:"MarkingOfMaterial"`
	ProductionVersion                  *int    `json:"ProductionVersion"`
	ProductionVersionItem              *int    `json:"ProductionVersionItem"`
	ProductionOrder                    *int    `json:"ProductionOrder"`
	ProductionOrderItem                *int    `json:"ProductionOrderItem"`
	Contract                           *int    `json:"Contract"`
	ContractItem                       *int    `json:"ContractItem"`
	Project                            *int    `json:"Project"`
	WBSElement                         *int    `json:"WBSElement"`
	InspectionLot                      int     `json:"InspectionLot"`
	InspectionPlantBusinessPartnerName string  `json:"InspectionPlantBusinessPartnerName"`
	HeatNumber                         string  `json:"HeatNumber"`
	DrawingNo                          string  `json:"DrawingNo"`
	PurchaseOrderNo                    string  `json:"PurchaseOrderNo"`
	Remarks                            string  `json:"Remarks"`
	ChiefOfInspectionSection           string  `json:"ChiefOfInspectionSection"`
}

type HeaderInspectionLot struct {
	OrderID                 int    `json:"OrderID"`
	OrderItem               int    `json:"OrderItem"`
	InspectionLot           int    `json:"InspectionLot"`
	InspectionLotDate       string `json:"InspectionLotDate"`
	InspectionSpecification string `json:"InspectionSpecification"`
}

type SpecDetails struct {
	OrderID          int      `json:"OrderID"`
	OrderItem        int      `json:"OrderItem"`
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
	OrderID                                    int      `json:"OrderID"`
	OrderItem                                  int      `json:"OrderItem"`
	InspectionLot                              int      `json:"InspectionLot"`
	ComponentCompositionType                   string   `json:"ComponentCompositionType"`
	ComponentCompositionUpperLimitInPercent    *float32 `json:"ComponentCompositionUpperLimitInPercent"`
	ComponentCompositionLowerLimitInPercent    *float32 `json:"ComponentCompositionLowerLimitInPercent"`
	ComponentCompositionStandardValueInPercent *float32 `json:"ComponentCompositionStandardValueInPercent"`
	HeatNumber                                 *string  `json:"HeatNumber"`
}

type Inspections struct {
	OrderID                                  int      `json:"OrderID"`
	OrderItem                                int      `json:"OrderItem"`
	InspectionLot                            int      `json:"InspectionLot"`
	Inspection                               int      `json:"Inspection"`
	InspectionType                           string   `json:"InspectionType"`
	InspectionTypeCertificateValueInText     *string  `json:"InspectionTypeCertificateValueInText"`
	InspectionTypeCertificateValueInQuantity *float32 `json:"InspectionTypeCertificateValueInQuantity"`
	InspectionTypeValueUnit                  *string  `json:"InspectionTypeValueUnit"`
	InspectionTypeTextInJA                   string   `json:"InspectionTypeTextInJA"`
}

func FunctionMillSheetPdfGenerates(
	input apiOutputFormatter.Orders,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_FUNCTION_MILL_SHEET_PDF_SRV"
	aPIType := "generates"

	var request MillSheetPdfReq

	if accepter == "MillSheet" {
		request = MillSheetPdfReq{
			Header:                make([]Header, len(input.OrdersItemSingleUnitMillSheetHeader)),
			HeaderInspectionLot:   make([]HeaderInspectionLot, len(input.OrdersItemSingleUnitMillSheetHeaderInspectionLot)),
			SpecDetails:           make([]SpecDetails, len(input.OrdersItemSingleUnitMillSheetSpecDetails)),
			ComponentCompositions: make([]ComponentCompositions, len(input.OrdersItemSingleUnitMillSheetComponentCompositions)),
			Inspections:           make([]Inspections, len(input.OrdersItemSingleUnitMillSheetInspections)),
			Accepter: []string{
				"MillSheet",
			},
		}

		for i, millSheetHeader := range input.OrdersItemSingleUnitMillSheetHeader {
			// 各要素をループして詰め替え
			request.Header[i] = Header{
				OrderID:                 millSheetHeader.OrderID,
				OrderItem:               millSheetHeader.OrderItem,
				BuyerName:               millSheetHeader.BuyerName,
				SellerName:              millSheetHeader.SellerName,
				Product:                 millSheetHeader.Product,
				SizeOrDimensionText:     millSheetHeader.SizeOrDimensionText,
				OrderQuantityInBaseUnit: millSheetHeader.OrderQuantityInBaseUnit,
				ProductSpecification:    millSheetHeader.ProductSpecification,
				MarkingOfMaterial:       millSheetHeader.MarkingOfMaterial,
				ProductionVersion:       millSheetHeader.ProductionVersion,
				ProductionVersionItem:   millSheetHeader.ProductionVersionItem,
				ProductionOrder:         millSheetHeader.ProductionOrder,
				ProductionOrderItem:     millSheetHeader.ProductionOrderItem,
				Contract:                millSheetHeader.Contract,
				ContractItem:            millSheetHeader.ContractItem,
				Project:                 millSheetHeader.Project,
				WBSElement:              millSheetHeader.WBSElement,
				InspectionLot:           millSheetHeader.InspectionLot,
				//InspectionPlantBusinessPartnerName: millSheetHeader.InspectionPlantBusinessPartnerName,
				//HeatNumber:                     millSheetHeader.HeatNumber,
				//DrawingNo:                      millSheetHeader.DrawingNo,
				//PurchaseOrderNo:                millSheetHeader.PurchaseOrderNo,
				//Remarks:                        millSheetHeader.Remarks,
				//ChiefOfInspectionSection:       millSheetHeader.ChiefOfInspectionSection,
			}
		}

		for i, millSheetHeaderInspectionLot := range input.OrdersItemSingleUnitMillSheetHeaderInspectionLot {
			request.HeaderInspectionLot[i] = HeaderInspectionLot{
				OrderID:                 millSheetHeaderInspectionLot.OrderID,
				OrderItem:               millSheetHeaderInspectionLot.OrderItem,
				InspectionLot:           millSheetHeaderInspectionLot.InspectionLot,
				InspectionLotDate:       millSheetHeaderInspectionLot.InspectionLotDate,
				InspectionSpecification: *millSheetHeaderInspectionLot.InspectionSpecification,
			}
		}

		for i, millSheetSpecDetails := range input.OrdersItemSingleUnitMillSheetSpecDetails {
			request.SpecDetails[i] = SpecDetails{
				OrderID:         millSheetSpecDetails.OrderID,
				OrderItem:       millSheetSpecDetails.OrderItem,
				InspectionLot:   millSheetSpecDetails.InspectionLot,
				SpecType:        millSheetSpecDetails.SpecType,
				UpperLimitValue: millSheetSpecDetails.UpperLimitValue,
				LowerLimitValue: millSheetSpecDetails.LowerLimitValue,
				StandardValue:   millSheetSpecDetails.StandardValue,
				SpecTypeUnit:    millSheetSpecDetails.SpecTypeUnit,
				//HeatNumber:       millSheetSpecDetails.HeatNumber,
				//SpecTypeTextInJA: millSheetSpecDetails.SpecTypeTextInJA,
			}
		}

		for i, millSheetComponentCompositions := range input.OrdersItemSingleUnitMillSheetComponentCompositions {
			request.ComponentCompositions[i] = ComponentCompositions{
				OrderID:                                    millSheetComponentCompositions.OrderID,
				OrderItem:                                  millSheetComponentCompositions.OrderItem,
				InspectionLot:                              millSheetComponentCompositions.InspectionLot,
				ComponentCompositionType:                   millSheetComponentCompositions.ComponentCompositionType,
				ComponentCompositionUpperLimitInPercent:    millSheetComponentCompositions.ComponentCompositionUpperLimitInPercent,
				ComponentCompositionLowerLimitInPercent:    millSheetComponentCompositions.ComponentCompositionLowerLimitInPercent,
				ComponentCompositionStandardValueInPercent: millSheetComponentCompositions.ComponentCompositionStandardValueInPercent,
				//HeatNumber:                                 millSheetComponentCompositions.HeatNumber,
			}
		}

		for i, millSheetInspections := range input.OrdersItemSingleUnitMillSheetInspections {
			request.Inspections[i] = Inspections{
				OrderID:                                  millSheetInspections.OrderID,
				OrderItem:                                millSheetInspections.OrderItem,
				InspectionLot:                            millSheetInspections.InspectionLot,
				Inspection:                               millSheetInspections.Inspection,
				InspectionType:                           millSheetInspections.InspectionType,
				InspectionTypeCertificateValueInText:     millSheetInspections.InspectionTypeCertificateValueInText,
				InspectionTypeCertificateValueInQuantity: millSheetInspections.InspectionTypeCertificateValueInQuantity,
				InspectionTypeValueUnit:                  millSheetInspections.InspectionTypeValueUnit,
				//InspectionTypeTextInJA:                   millSheetInspections.InspectionTypeTextInJA,
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
	)

	return responseBody
}
