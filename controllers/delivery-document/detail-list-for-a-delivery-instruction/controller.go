package controllersDeliveryDocumentDetailListForADeliveryInstruction

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsDeliveryDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/delivery-document/delivery-document"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProject "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/project"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesDeliveryDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/delivery-document"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProject "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/project"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"io/ioutil"
	"strconv"
	"strings"
)

type DeliveryDocumentDetailListForADeliveryInstructionController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *DeliveryDocumentDetailListForADeliveryInstructionController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	deliveryDocument, _ := controller.GetInt("deliveryDocument")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "delivery-document"
	redisKeyCategory2 := "detail-list-for-a-delivery-instruction"

	deliveryDocumentHeader := apiInputReader.DeliveryDocument{}

	deliveryDocumentHeader = apiInputReader.DeliveryDocument{
		DeliveryDocumentHeader: &apiInputReader.DeliveryDocumentHeader{
			DeliveryDocument: deliveryDocument,
		},
		DeliveryDocumentItems: &apiInputReader.DeliveryDocumentItems{
			DeliveryDocument: deliveryDocument,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(deliveryDocument),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.DeliveryDocument

		err := json.Unmarshal(cacheData, &responseData)

		if err != nil {
			services.HandleError(
				&controller.Controller,
				err,
				nil,
			)
		}

		services.Respond(
			&controller.Controller,
			&responseData,
		)
	}

	if cacheData != nil {
		go func() {
			controller.request(deliveryDocumentHeader)
		}()
	} else {
		controller.request(deliveryDocumentHeader)
	}
}

func (
	controller *DeliveryDocumentDetailListForADeliveryInstructionController,
) createDeliveryDocumentRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.DeliveryDocument,
) *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes {
	responseJsonData := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes{}
	responseBody := apiModuleRuntimesRequestsDeliveryDocument.DeliveryDocumentReads(
		requestPram,
		input,
		&controller.Controller,
		"Header",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("DeliveryDocumentReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentDetailListForADeliveryInstructionController,
) createDeliveryDocumentRequestItems(
	requestPram *apiInputReader.Request,
	input apiInputReader.DeliveryDocument,
) *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes {
	responseJsonData := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes{}
	responseBody := apiModuleRuntimesRequestsDeliveryDocument.DeliveryDocumentReads(
		requestPram,
		input,
		&controller.Controller,
		"Items",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("DeliveryDocumentReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentDetailListForADeliveryInstructionController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	deliveryDocumentRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*deliveryDocumentRes.Message.Header))

	for _, v := range *deliveryDocumentRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.DeliverToParty,
		})
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.DeliverFromParty,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsGeneralsByBusinessPartners(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("BusinessPartnerGeneralReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentDetailListForADeliveryInstructionController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	deliveryDocumentRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*deliveryDocumentRes.Message.Item))

	for _, v := range *deliveryDocumentRes.Message.Item {
		input = append(input, apiModuleRuntimesRequestsPlant.General{
			BusinessPartner: v.DeliverToParty,
			Plant:           v.DeliverToPlant,
		})
		input = append(input, apiModuleRuntimesRequestsPlant.General{
			BusinessPartner: v.DeliverFromParty,
			Plant:           v.DeliverFromPlant,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesPlant.PlantRes{}
	responseBody := apiModuleRuntimesRequestsPlant.PlantReadsGeneralsByPlants(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("PlantReadsGeneralsByPlants Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentDetailListForADeliveryInstructionController,
) CreateProjectRequestWBSElement(
	requestPram *apiInputReader.Request,
	deliveryDocumentRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
) *apiModuleRuntimesResponsesProject.ProjectRes {
	var input apiModuleRuntimesRequestsProject.WBSElement
	//input := make([]apiModuleRuntimesRequestsProject.Project, len(*deliveryDocumentRes.Message.Item))

	for _, v := range *deliveryDocumentRes.Message.Header {
		input = apiModuleRuntimesRequestsProject.WBSElement{
			Project:    *v.Project,
			WBSElement: *v.WBSElement,
		}
	}

	responseJsonData := apiModuleRuntimesResponsesProject.ProjectRes{}
	responseBody := apiModuleRuntimesRequestsProject.ProjectReadsWBSElement(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("ProjectReadsWBSElement Unmarshal error")
	}

	return &responseJsonData
}

func functionDeliveryInstructionPdfGenerates(
	input apiOutputFormatter.DeliveryDocument,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_FUNCTION_DELIVERY_INSTRUCTION_PDF_SRV"
	aPIType := "generates"

	if accepter == "DeliveryInstruction" {
		input.Accepter = []string{
			"DeliveryInstruction",
		}
	}

	marshaledRequest, err := json.Marshal(input)
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

func (
	controller *DeliveryDocumentDetailListForADeliveryInstructionController,
) request(
	input apiInputReader.DeliveryDocument,
) {
	defer services.Recover(controller.CustomLogger)

	headerRes := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes{}
	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	headerRes = *controller.createDeliveryDocumentRequestHeader(
		controller.UserInfo,
		input,
	)

	businessPartnerRes = *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		&headerRes,
	)

	itemRes := controller.createDeliveryDocumentRequestItems(
		controller.UserInfo,
		input,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		itemRes,
	)

	wBSElementRes := controller.CreateProjectRequestWBSElement(
		controller.UserInfo,
		&headerRes,
	)

	controller.fin(
		&headerRes,
		itemRes,
		&businessPartnerRes,
		plantRes,
		wBSElementRes,
	)
}

func (
	controller *DeliveryDocumentDetailListForADeliveryInstructionController,
) fin(
	headerRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
	itemRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	wBSElementRes *apiModuleRuntimesResponsesProject.ProjectRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	wBSElementMapper := services.WBSElementMapper(
		wBSElementRes.Message.WBSElement,
	)

	data := apiOutputFormatter.DeliveryDocument{}

	for _, v := range *headerRes.Message.Header {
		var wBSElementDescription *string

		wBSElementMapperDescription := wBSElementMapper[*v.WBSElement].WBSElementDescription
		if &wBSElementMapperDescription != nil {
			wBSElementDescription = &wBSElementMapperDescription
		}

		data.DeliveryDocumentHeaderWithItem = append(data.DeliveryDocumentHeaderWithItem,
			apiOutputFormatter.DeliveryDocumentHeaderWithItem{
				DeliveryDocument:        v.DeliveryDocument,
				DeliveryDocumentDate:    v.DeliveryDocumentDate,
				DeliverToParty:          v.DeliverToParty,
				DeliverToPartyName:      businessPartnerMapper[v.DeliverToParty].BusinessPartnerName,
				DeliverToPlant:          v.DeliverToPlant,
				DeliverToPlantName:      plantMapper[strconv.Itoa(v.DeliverToParty)].PlantName,
				DeliverFromParty:        v.DeliverFromParty,
				DeliverFromPartyName:    businessPartnerMapper[v.DeliverFromParty].BusinessPartnerName,
				DeliverFromPlant:        v.DeliverFromPlant,
				DeliverFromPlantName:    plantMapper[strconv.Itoa(v.DeliverFromParty)].PlantName,
				IsExportImport:          v.IsExportImport,
				OrderID:                 v.OrderID,
				OrderItem:               v.OrderItem,
				Contract:                v.Contract,
				ContractItem:            v.ContractItem,
				Project:                 v.Project,
				WBSElement:              v.WBSElement,
				WBSElementDescription:   wBSElementDescription,
				ProductionOrder:         v.ProductionOrder,
				ProductionOrderItem:     v.ProductionOrderItem,
				PlannedGoodsIssueDate:   v.PlannedGoodsIssueDate,
				PlannedGoodsIssueTime:   v.PlannedGoodsIssueTime,
				PlannedGoodsReceiptDate: v.PlannedGoodsReceiptDate,
				PlannedGoodsReceiptTime: v.PlannedGoodsReceiptTime,
				HeaderGrossWeight:       *v.HeaderGrossWeight,
				HeaderNetWeight:         *v.HeaderNetWeight,
				HeaderWeightUnit:        *v.HeaderWeightUnit,
				Incoterms:               v.Incoterms,
			},
		)
	}

	for _, v := range *itemRes.Message.Item {
		data.DeliveryDocumentItem = append(data.DeliveryDocumentItem,
			apiOutputFormatter.DeliveryDocumentItem{
				DeliveryDocument:                     v.DeliveryDocument,
				DeliveryDocumentItem:                 v.DeliveryDocumentItem,
				DeliveryDocumentItemCategory:         v.DeliveryDocumentItemCategory,
				Product:                              v.Product,
				ProductSpecification:                 v.ProductSpecification,
				SizeOrDimensionText:                  v.SizeOrDimensionText,
				DeliveryDocumentItemText:             v.DeliveryDocumentItemText,
				DeliveryDocumentItemItemTextByBuyer:  v.DeliveryDocumentItemTextByBuyer,
				DeliveryDocumentItemItemTextBySeller: v.DeliveryDocumentItemTextBySeller,
				PlannedGoodsIssueQuantity:            v.PlannedGoodsIssueQuantity,
				PlannedGoodsIssueQtyInBaseUnit:       v.PlannedGoodsIssueQtyInBaseUnit,
				BaseUnit:                             v.BaseUnit,
				DeliveryUnit:                         v.DeliveryUnit,
				PlannedGoodsIssueDate:                v.PlannedGoodsIssueDate,
				PlannedGoodsIssueTime:                v.PlannedGoodsIssueTime,
				PlannedGoodsReceiptDate:              v.PlannedGoodsReceiptDate,
				PlannedGoodsReceiptTime:              v.PlannedGoodsReceiptTime,
				ItemGrossWeight:                      v.ItemGrossWeight,
				ItemNetWeight:                        v.ItemNetWeight,
				ItemWeightUnit:                       v.ItemWeightUnit,
				ProductNetWeight:                     v.ProductNetWeight,
				Project:                              v.Project,
				WBSElement:                           v.WBSElement,
				OrderID:                              v.OrderID,
				OrderItem:                            v.OrderItem,
			},
		)
	}

	// ここから generates に rabbitmq で送信
	// accepter 対応
	responseJsonData := apiOutputFormatter.DeliveryDocument{}
	responseBody := functionDeliveryInstructionPdfGenerates(
		data,
		&controller.Controller,
		"DeliveryInstruction",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("apiModuleRuntimesRequestsDeliveryInstructionPdf.FunctionDeliveryInstructionPdfGenerates Unmarshal error")
	}

	data.MountPath = responseJsonData.MountPath

	//err = controller.RedisCache.SetCache(
	//	controller.RedisKey,
	//	data,
	//)
	//if err != nil {
	//	services.HandleError(
	//		&controller.Controller,
	//		err,
	//		nil,
	//	)
	//}

	controller.Data["json"] = &data
	controller.ServeJSON()
}
